use futures::StreamExt;
use crate::types::{MessageResult,NatsStartupError};
use crate::nats_types::PublishFixRequest;

#[allow(dead_code)]
#[derive(Debug)]
enum PublishError{
	Messages(async_nats::jetstream::consumer::pull::MessagesError),
	DoubleAck(async_nats::Error),
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
	messages:async_nats::jetstream::consumer::pull::Stream,
	roblox_cookie:rbx_asset::cookie::CookieContext,
	api:api::Context,
}
impl Publisher{
	pub async fn new(
		stream:async_nats::jetstream::stream::Stream,
		roblox_cookie:rbx_asset::cookie::CookieContext,
		api:api::Context,
	)->Result<Self,NatsStartupError>{
		Ok(Self{
			messages:stream.get_or_create_consumer("publish_fix",async_nats::jetstream::consumer::pull::Config{
				name:Some("publish_fix".to_owned()),
				durable_name:Some("publish_fix".to_owned()),
				filter_subject:"maptest.submissions.publish.fix".to_owned(),
				..Default::default()
			}).await.map_err(NatsStartupError::Consumer)?
			.messages().await.map_err(NatsStartupError::Stream)?,
			roblox_cookie,
			api,
		})
	}
	pub async fn run(mut self){
		while let Some(message_result)=self.messages.next().await{
			self.publish_supress_error(message_result).await
		}
	}
	async fn publish_supress_error(&self,message_result:MessageResult){
		match self.publish(message_result).await{
			Ok(())=>println!("[PublishFix] Published, hooray!"),
			Err(e)=>println!("[PublishFix] There was an error, oopsie! {e}"),
		}
	}
	async fn publish(&self,message_result:MessageResult)->Result<(),PublishError>{
		println!("publish_fix {:?}",message_result);
		let message=message_result.map_err(PublishError::Messages)?;
		message.double_ack().await.map_err(PublishError::DoubleAck)?;
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
