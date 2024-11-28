package model

import "time"

type Status int32

const(
	Published			Status=8
	Rejected			Status=7

	Publishing			Status=6
	Validated			Status=5
	Validating			Status=4
	Accepted			Status=3

	ChangesRequested	Status=2
	Submitted			Status=1
	UnderConstruction	Status=0
)

type Submission struct {
	ID             int64
	DisplayName    string
	Creator        string
	GameID         int32
	Date           time.Time
	Submitter      int64 // UserID
	AssetID        int64
	AssetVersion   int64
	Completed      bool
	TargetAssetID  int64 // where to upload map fix.  if the TargetAssetID is 0, it's a new map.
	StatusID       Status
}
