use futures::StreamExt;

#[derive(Debug)]
enum PublishError{
}
impl std::fmt::Display for PublishError{
	fn fmt(&self,f:&mut std::fmt::Formatter<'_>)->std::fmt::Result{
		write!(f,"{self:?}")
	}
}
impl std::error::Error for PublishError{}

pub struct Publisher{
	subscriber:async_nats::Subscriber,
	roblox_cookie:rbx_asset::cookie::CookieContext,
}
impl Publisher{
	pub async fn new(
		nats:async_nats::Client,
		roblox_cookie:rbx_asset::cookie::CookieContext,
	)->Result<Self,async_nats::SubscribeError>{
		Ok(Self{
			subscriber:nats.subscribe("publish_new").await?,
			roblox_cookie,
		})
	}
	pub async fn run(mut self){
		while let Some(message)=self.subscriber.next().await{
			self.publish(message).await
		}
	}
	async fn publish(&self,message:async_nats::Message){
		println!("publish {:?}",message);
	}
}
