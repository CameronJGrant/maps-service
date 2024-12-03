mod nats_types;
mod publisher;
mod validator;

#[allow(dead_code)]
#[derive(Debug)]
enum StartupError{
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

	// nats
	let nasty=async_nats::connect("nats").await.map_err(StartupError::Connect)?;

	// connect to nats
	let (publisher,validator)=tokio::try_join!(
		publisher::Publisher::new(nasty.clone(),cookie_context.clone()),
		validator::Validator::new(nasty,cookie_context)
	).map_err(StartupError::Subscribe)?;

	// publisher thread
	tokio::spawn(publisher.run());

	// run validator on the main thread indefinitely
	validator.run().await;

	Ok(())
}
