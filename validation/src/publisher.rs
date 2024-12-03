use futures::StreamExt;

enum PublishError{
}

pub struct Publisher{
	nats:async_nats::Client,
	subscriber:async_nats::Subscriber,
}
impl Publisher{
	pub async fn new(nats:async_nats::Client)->Result<Self,async_nats::SubscribeError>{
		Ok(Self{
			subscriber:nats.subscribe("publish").await?,
			nats,
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
