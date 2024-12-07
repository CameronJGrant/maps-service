use futures::StreamExt;

use crate::nats_types::PublishNewRequest;

#[derive(Debug)]
enum PublishError{
	Get(rbx_asset::cookie::GetError),
	Json(serde_json::Error),
	Create(rbx_asset::cookie::CreateError),
	ApiActionSubmissionPublish(api::Error),
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
	api:api::Context,
}
impl Publisher{
	pub async fn new(
		nats:async_nats::Client,
		roblox_cookie:rbx_asset::cookie::CookieContext,
		api:api::Context,
	)->Result<Self,async_nats::SubscribeError>{
		Ok(Self{
			subscriber:nats.subscribe("publish_new").await?,
			roblox_cookie,
			api,
		})
	}
	pub async fn run(mut self){
		while let Some(message)=self.subscriber.next().await{
			self.publish_supress_error(message).await
		}
	}
	async fn publish_supress_error(&self,message:async_nats::Message){
		match self.publish(message).await{
			Ok(())=>println!("Published, hooray!"),
			Err(e)=>println!("[PublishNew] There was an error, oopsie! {e}"),
		}
	}
	async fn publish(&self,message:async_nats::Message)->Result<(),PublishError>{
		println!("publish {:?}",message);
		// decode json
		let publish_info:PublishNewRequest=serde_json::from_slice(&message.payload).map_err(PublishError::Json)?;

		// download the map model version
		let model_data=self.roblox_cookie.get_asset(rbx_asset::cookie::GetAssetRequest{
			asset_id:publish_info.model_id,
			version:Some(publish_info.model_version),
		}).await.map_err(PublishError::Get)?;

		// upload the map to the strafesnet group
		let upload_response=self.roblox_cookie.create(rbx_asset::cookie::CreateRequest{
			name:publish_info.display_name,
			description:"".to_owned(),
			ispublic:false,
			allowComments:false,
			groupId:Some(crate::GROUP_STRAFESNET),
		},model_data).await.map_err(PublishError::Create)?;

		// create the map entry in the game database, including release date
		// game_rpc.maps.create(upload_response.AssetId)

		// mark submission as published
		self.api.action_submission_publish(
			api::SubmissionID(publish_info.submission_id)
		).await.map_err(PublishError::ApiActionSubmissionPublish)?;

		Ok(())
	}
}
