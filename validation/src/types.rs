
// annoying mile-long types
pub type MapsServiceClient=rust_grpc::maps::maps_service_client::MapsServiceClient<tonic::transport::channel::Channel>;
pub type MessageResult=Result<async_nats::jetstream::Message,async_nats::jetstream::consumer::pull::MessagesError>;

#[derive(Debug)]
pub enum NatsStartupError{
	GetStream(async_nats::jetstream::context::GetStreamError),
	ConsumerCreateStrict(async_nats::jetstream::stream::ConsumerCreateStrictError),
	Stream(async_nats::jetstream::consumer::StreamError),
}
impl std::fmt::Display for NatsStartupError{
	fn fmt(&self,f:&mut std::fmt::Formatter<'_>)->std::fmt::Result{
		write!(f,"{self:?}")
	}
}
impl std::error::Error for NatsStartupError{}
