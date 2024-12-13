use futures::StreamExt;

mod nats_types;
mod validator;
mod publish_new;
mod publish_fix;
mod message_handler;

#[allow(dead_code)]
#[derive(Debug)]
pub enum StartupError{
	API(api::ReqwestError),
	NatsConnect(async_nats::ConnectError),
	NatsGetStream(async_nats::jetstream::context::GetStreamError),
	NatsConsumer(async_nats::jetstream::stream::ConsumerError),
	NatsStream(async_nats::jetstream::consumer::StreamError),
	GRPCConnect(tonic::transport::Error),
}
impl std::fmt::Display for StartupError{
	fn fmt(&self,f:&mut std::fmt::Formatter<'_>)->std::fmt::Result{
		write!(f,"{self:?}")
	}
}
impl std::error::Error for StartupError{}

// annoying mile-long type
pub type MapsServiceClient=rust_grpc::maps::maps_service_client::MapsServiceClient<tonic::transport::channel::Channel>;

pub const GROUP_STRAFESNET:u64=6980477;

pub const PARALLEL_REQUESTS:usize=16;

#[tokio::main]
async fn main()->Result<(),StartupError>{
	// talk to roblox through STRAFESNET_CI2 account
	let cookie=std::env::var("RBXCOOKIE").expect("RBXCOOKIE env required");
	let cookie_context=rbx_asset::cookie::CookieContext::new(rbx_asset::cookie::Cookie::new(cookie));

	// maps-service api
	let api_host=std::env::var("API_HOST").expect("API_HOST env required");
	let api=api::Context::new(api_host).map_err(StartupError::API)?;

	// nats
	let nats_host=std::env::var("NATS_HOST").expect("NATS_HOST env required");
	let nats_fut=async{
		let nasty=async_nats::connect(nats_host).await.map_err(StartupError::NatsConnect)?;
		// use nats jetstream
		async_nats::jetstream::new(nasty)
			.get_stream("maptest").await.map_err(StartupError::NatsGetStream)?
			.get_or_create_consumer("validation",async_nats::jetstream::consumer::pull::Config{
				name:Some("validation".to_owned()),
				durable_name:Some("validation".to_owned()),
				filter_subject:"maptest.submissions.>".to_owned(),
				..Default::default()
			}).await.map_err(StartupError::NatsConsumer)?
			.messages().await.map_err(StartupError::NatsStream)
	};

	// data-service grpc for creating map entries
	let data_host=std::env::var("DATA_HOST").expect("DATA_HOST env required");
	let message_handler_fut=async{
		let maps_grpc=crate::MapsServiceClient::connect(data_host).await.map_err(StartupError::GRPCConnect)?;
		Ok(message_handler::MessageHandler::new(cookie_context,api,maps_grpc))
	};

	// Create a signal listener for SIGTERM
	let mut sig_term=tokio::signal::unix::signal(tokio::signal::unix::SignalKind::terminate()).expect("Failed to create SIGTERM signal listener");

	// run futures
	let (mut messages,message_handler)=tokio::try_join!(nats_fut,message_handler_fut)?;

	// process up to PARALLEL_REQUESTS in parallel
	let main_loop=async move{
		static SEM:tokio::sync::Semaphore=tokio::sync::Semaphore::const_new(PARALLEL_REQUESTS);
		// use memory leak to make static lifetime
		let message_handler=Box::leak(Box::new(message_handler));
		// acquire a permit before attempting to receive a message, exit if either fails
		while let (Ok(permit),Some(message_result))=(SEM.acquire().await,messages.next().await){
			// handle the message on a new thread (mainly to decode the model file)
			tokio::spawn(async{
				match message_handler.handle_message_result(message_result).await{
					Ok(())=>println!("[Validation] Success, hooray!"),
					Err(e)=>println!("[Validation] There was an error, oopsie! {e}"),
				}
				// explicitly call drop to make the move semantics and permit release more obvious
				core::mem::drop(permit);
			});
		}
	};

	// race sigkill and main loop termination and then die
	tokio::select!{
		_=sig_term.recv()=>(),
		_=main_loop=>(),
	};

	Ok(())
}
