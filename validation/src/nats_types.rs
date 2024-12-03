// These represent the information needed in the nats message
// to perform the operation, not necessarily the over-the-wire format

// Requests are sent from maps-service to validator
// Responses are sent from validator to maps-service

#[derive(serde::Deserialize)]
pub struct ValidateRequest{
	// submission_id is passed back in the response message
	pub submission_id:u64,
	pub model_id:u64,
	pub model_version:u64,
}

pub struct ValidateResponse{
	pub submission_id:u64,
	// model id will be changed if scripts were replaced
	pub model_id:u64,
	pub model_version:u64,
}

// an invalidate is sent instead if validation fails
pub struct InvalidateResponse{
	pub submission_id:u64,
}

// Create a new map
pub struct PublishNewRequest{
	pub submission_id:u64,
	pub model_id:u64,
	pub model_version:u64,
	pub creator:String,
	pub display_name:String,
	//games:HashSet<GameID>,
}

pub struct PublishFixRequest{
	pub submission_id:u64,
	pub model_id:u64,
	pub model_version:u64,
	pub target_asset_id:u64,
}

pub struct PublishResponse{
	pub submission_id:u64,
	pub success:bool,
}
