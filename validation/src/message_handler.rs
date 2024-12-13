#[allow(dead_code)]
#[derive(Debug)]
pub enum HandleMessageError{
	Messages(async_nats::jetstream::consumer::pull::MessagesError),
	DoubleAck(async_nats::Error),
	UnknownSubject(String),
	PublishNew(crate::publish_new::PublishError),
	PublishFix(crate::publish_fix::PublishError),
	Validation(crate::validator::ValidateError),
}
impl std::fmt::Display for HandleMessageError{
	fn fmt(&self,f:&mut std::fmt::Formatter<'_>)->std::fmt::Result{
		write!(f,"{self:?}")
	}
}
impl std::error::Error for HandleMessageError{}

pub type MessageResult=Result<async_nats::jetstream::Message,async_nats::jetstream::consumer::pull::MessagesError>;

pub struct MessageHandler{
	publish_new:crate::publish_new::Publisher,
	publish_fix:crate::publish_fix::Publisher,
	validator:crate::validator::Validator,
}

impl MessageHandler{
	pub fn new(
		cookie_context:rbx_asset::cookie::CookieContext,
		api:api::Context,
		maps_grpc:crate::MapsServiceClient,
	)->Self{
		Self{
			publish_new:crate::publish_new::Publisher::new(cookie_context.clone(),api.clone(),maps_grpc),
			publish_fix:crate::publish_fix::Publisher::new(cookie_context.clone(),api.clone()),
			validator:crate::validator::Validator::new(cookie_context,api),
		}
	}
	pub async fn handle_message_result(&self,message_result:MessageResult)->Result<(),HandleMessageError>{
		let message=message_result.map_err(HandleMessageError::Messages)?;
		message.double_ack().await.map_err(HandleMessageError::DoubleAck)?;
		match message.subject.as_str(){
			"maptest.submissions.publish.new"=>self.publish_new.publish(message).await.map_err(HandleMessageError::PublishNew),
			"maptest.submissions.publish.fix"=>self.publish_fix.publish(message).await.map_err(HandleMessageError::PublishFix),
			"maptest.submissions.validate"=>self.validator.validate(message).await.map_err(HandleMessageError::Validation),
			other=>Err(HandleMessageError::UnknownSubject(other.to_owned()))
		}
	}
}
