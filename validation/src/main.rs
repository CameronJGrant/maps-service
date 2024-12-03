mod publisher;
mod validator;

#[allow(dead_code)]
#[derive(Debug)]
enum StartupError{
	Connect(async_nats::ConnectError),
	Subscribe(async_nats::SubscribeError),
}

#[tokio::main]
async fn main()->Result<(),StartupError>{
	let nasty=async_nats::connect("nats").await.map_err(StartupError::Connect)?;
	let (publisher,validator)=tokio::try_join!(
		publisher::Publisher::new(nasty.clone()),
		validator::Validator::new(nasty)
	).map_err(StartupError::Subscribe)?;
	// publisher thread
	tokio::spawn(publisher.run());
	// run validator on the main thread indefinitely
	validator.run().await;
	Ok(())
}
