use futures::StreamExt;

use crate::nats_types::PublishFixRequest;

#[allow(dead_code)]
#[derive(Debug)]
enum PublishError{
	Get(rbx_asset::cookie::GetError),
	Json(serde_json::Error),
	Upload(rbx_asset::cookie::UploadError),
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
			subscriber:nats.subscribe("publish_fix").await?,
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
			Ok(())=>println!("[PublishFix] Published, hooray!"),
			Err(e)=>println!("[PublishFix] There was an error, oopsie! {e}"),
		}
	}
	async fn publish(&self,message:async_nats::Message)->Result<(),PublishError>{
		println!("publish_fix {:?}",message);
		// decode json
		let publish_info:PublishFixRequest=serde_json::from_slice(&message.payload).map_err(PublishError::Json)?;

		// download the map model version
		let model_data=self.roblox_cookie.get_asset(rbx_asset::cookie::GetAssetRequest{
			asset_id:publish_info.ModelID,
			version:Some(publish_info.ModelVersion),
		}).await.map_err(PublishError::Get)?;

		// upload the map to the strafesnet group
		let _upload_response=self.roblox_cookie.upload(rbx_asset::cookie::UploadRequest{
			assetid:publish_info.TargetAssetID,
			groupId:Some(crate::GROUP_STRAFESNET),
			name:None,
			description:None,
			ispublic:None,
			allowComments:None,
		},model_data).await.map_err(PublishError::Upload)?;

		// that's it, the database entry does not need to be changed.

		// mark submission as published
		self.api.action_submission_publish(
			api::SubmissionID(publish_info.SubmissionID)
		).await.map_err(PublishError::ApiActionSubmissionPublish)?;

		Ok(())
	}
}
