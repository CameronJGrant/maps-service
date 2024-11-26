package model

import "time"

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
	StatusID       int32
}
