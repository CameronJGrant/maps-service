package model

// These represent the information needed in the nats message
// to perform the operation, they are encoded to json

// Requests are sent from maps-service to validator

type ValidateRequest struct{
	// submission_id is passed back in the response message
	SubmissionID     int64
	ModelID          uint64
	ModelVersion     uint64
	ValidatedModelID uint64 // optional value
}

// Create a new map
type PublishNewRequest struct{
	SubmissionID int64
	ModelID      uint64
	ModelVersion uint64
	Creator      string
	DisplayName  string
	GameID       uint32
	//games HashSet<GameID>
}

type PublishFixRequest struct{
	SubmissionID  int64
	ModelID       uint64
	ModelVersion  uint64
	TargetAssetID uint64
}
