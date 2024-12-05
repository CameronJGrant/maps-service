mod nats_types;
mod publisher;
mod validator;

#[allow(dead_code)]
#[derive(Debug)]
enum StartupError{
	API(api::ReqwestError),
	Connect(async_nats::ConnectError),
	Subscribe(async_nats::SubscribeError),
}
impl std::fmt::Display for StartupError{
	fn fmt(&self,f:&mut std::fmt::Formatter<'_>)->std::fmt::Result{
		write!(f,"{self:?}")
	}
}
impl std::error::Error for StartupError{}

#[tokio::main]
async fn main()->Result<(),StartupError>{
	// talk to roblox
	let cookie_context=rbx_asset::cookie::CookieContext::new(rbx_asset::cookie::Cookie::new("".to_owned()));

	// maps-service api
	let api=api::Context::new("https:://localhost:1234/v1".to_owned()).map_err(StartupError::API)?;

	// nats
	let nasty=async_nats::connect("nats").await.map_err(StartupError::Connect)?;

	// connect to nats
	let (publisher,validator)=tokio::try_join!(
		publisher::Publisher::new(nasty.clone(),cookie_context.clone()),
		// clone nats here because it's dropped within the function scope,
		// meanining the last reference is dropped...
		validator::Validator::new(nasty.clone(),cookie_context,api)
	).map_err(StartupError::Subscribe)?;

	// publisher thread
	tokio::spawn(publisher.run());

	// run validator on the main thread indefinitely
	validator.run().await;

	Ok(())
}
