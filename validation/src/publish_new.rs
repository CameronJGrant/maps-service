use futures::StreamExt;

use crate::nats_types::PublishNewRequest;

#[allow(dead_code)]
#[derive(Debug)]
enum PublishError{
	Get(rbx_asset::cookie::GetError),
	Json(serde_json::Error),
	Create(rbx_asset::cookie::CreateError),
	SystemTime(std::time::SystemTimeError),
	MapCreate(tonic::Status),
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
	maps_grpc:crate::MapsServiceClient,
}
impl Publisher{
	pub async fn new(
		nats:async_nats::Client,
		roblox_cookie:rbx_asset::cookie::CookieContext,
		api:api::Context,
		maps_grpc:crate::MapsServiceClient,
	)->Result<Self,async_nats::SubscribeError>{
		Ok(Self{
			subscriber:nats.subscribe("publish_new").await?,
			roblox_cookie,
			api,
			maps_grpc,
		})
	}
	pub async fn run(mut self){
		while let Some(message)=self.subscriber.next().await{
			self.publish_supress_error(message).await
		}
	}
	async fn publish_supress_error(&self,message:async_nats::Message){
		match self.publish(message).await{
			Ok(())=>println!("[PublishNew] Published, hooray!"),
			Err(e)=>println!("[PublishNew] There was an error, oopsie! {e}"),
		}
	}
	async fn publish(&self,message:async_nats::Message)->Result<(),PublishError>{
		println!("publish_new {:?}",message);
		// decode json
		let publish_info:PublishNewRequest=serde_json::from_slice(&message.payload).map_err(PublishError::Json)?;

		// download the map model version
		let model_data=self.roblox_cookie.get_asset(rbx_asset::cookie::GetAssetRequest{
			asset_id:publish_info.ModelID,
			version:Some(publish_info.ModelVersion),
		}).await.map_err(PublishError::Get)?;

		// upload the map to the strafesnet group
		let upload_response=self.roblox_cookie.create(rbx_asset::cookie::CreateRequest{
			name:publish_info.DisplayName.clone(),
			description:"".to_owned(),
			ispublic:false,
			allowComments:false,
			groupId:Some(crate::GROUP_STRAFESNET),
		},model_data).await.map_err(PublishError::Create)?;

		// create the map entry in the game database, including release date
		self.maps_grpc.clone().create(rust_grpc::maps::MapRequest{
			id:upload_response.AssetId as i64,
			display_name:Some(publish_info.DisplayName),
			creator:Some(publish_info.Creator),
			game_id:Some(publish_info.GameID as i32),
			// TODO: scheduling algorithm
			date:Some(
				// Publish one week from now
				(
					std::time::SystemTime::now()
					+std::time::Duration::from_secs(7*24*60*60)
				)
				.duration_since(std::time::SystemTime::UNIX_EPOCH)
				.map_err(PublishError::SystemTime)?
				.as_secs() as i64
			),
		}).await.map_err(PublishError::MapCreate)?;

		// mark submission as published
		self.api.action_submission_publish(
			api::SubmissionID(publish_info.SubmissionID)
		).await.map_err(PublishError::ApiActionSubmissionPublish)?;

		Ok(())
	}
}
