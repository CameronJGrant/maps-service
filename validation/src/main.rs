mod nats_types;
mod validator;
mod publish_new;
mod publish_fix;

#[allow(dead_code)]
#[derive(Debug)]
enum StartupError{
	API(api::ReqwestError),
	NatsConnect(async_nats::ConnectError),
	NatsSubscribe(async_nats::SubscribeError),
	GRPCConnect(tonic::transport::Error),
}
impl std::fmt::Display for StartupError{
	fn fmt(&self,f:&mut std::fmt::Formatter<'_>)->std::fmt::Result{
		write!(f,"{self:?}")
	}
}
impl std::error::Error for StartupError{}

pub const GROUP_STRAFESNET:u64=6980477;

// annoying mile-long type
pub type MapsServiceClient=rust_grpc::maps::maps_service_client::MapsServiceClient<tonic::transport::channel::Channel>;

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
	let nasty=async_nats::connect(nats_host).await.map_err(StartupError::NatsConnect)?;

	// data-service grpc for creating map entries
	let data_host=std::env::var("DATA_HOST").expect("DATA_HOST env required");
	let maps_grpc=MapsServiceClient::connect(data_host).await.map_err(StartupError::GRPCConnect)?;

	// connect to nats
	let (publish_new,publish_fix,validator)=tokio::try_join!(
		publish_new::Publisher::new(nasty.clone(),cookie_context.clone(),api.clone(),maps_grpc),
		publish_fix::Publisher::new(nasty.clone(),cookie_context.clone()),
		// clone nats here because it's dropped within the function scope,
		// meanining the last reference is dropped...
		validator::Validator::new(nasty.clone(),cookie_context,api)
	).map_err(StartupError::NatsSubscribe)?;

	// publisher threads
	tokio::spawn(publish_new.run());
	tokio::spawn(publish_fix.run());

	// run validator on the main thread indefinitely
	validator.run().await;

	Ok(())
}
