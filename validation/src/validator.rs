use futures::StreamExt;
use rbx_dom_weak::WeakDom;

enum Valid{
	Untouched(WeakDom),
	Modified(WeakDom),
}

enum ValidateError{
}


pub struct Validator{
	nats:async_nats::Client,
	subscriber:async_nats::Subscriber,
}
impl Validator{
	pub async fn new(nats:async_nats::Client)->Result<Self,async_nats::SubscribeError>{
		Ok(Self{
			subscriber:nats.subscribe("validate").await?,
			nats,
		})
	}
	pub async fn run(mut self){
		while let Some(message)=self.subscriber.next().await{
			self.validate(message).await
		}
	}
	async fn validate(&self,message:async_nats::Message){
		println!("validate {:?}",message);
		// download map
		// validate map
		// validate(dom)
		// reply with validity
		// Ok(Valid::Untouched(dom))
	}
}
