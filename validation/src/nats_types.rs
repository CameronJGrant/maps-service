// These represent the information needed in the nats message
// to perform the operation, not necessarily the over-the-wire format

// Requests are sent from maps-service to validator
// Validation invokes the REST api to update the submissions

#[allow(nonstandard_style)]
#[derive(serde::Deserialize)]
pub struct ValidateRequest{
	// submission_id is passed back in the response message
	pub SubmissionID:i64,
	pub ModelID:u64,
	pub ModelVersion:u64,
	pub ValidatedModelID:Option<u64>,
}

// Create a new map
#[allow(nonstandard_style)]
#[derive(serde::Deserialize)]
pub struct PublishNewRequest{
	pub SubmissionID:i64,
	pub ModelID:u64,
	pub ModelVersion:u64,
	pub Creator:String,
	pub DisplayName:String,
	pub GameID:u32,
	//games:HashSet<GameID>,
}

#[allow(nonstandard_style)]
#[derive(serde::Deserialize)]
pub struct PublishFixRequest{
	pub SubmissionID:i64,
	pub ModelID:u64,
	pub ModelVersion:u64,
	pub TargetAssetID:u64,
}
