package model

import "time"

type Status int32

const(
	StatusPublished			Status=8
	StatusRejected			Status=7

	StatusPublishing		Status=6
	StatusValidated			Status=5
	StatusValidating		Status=4
	StatusAccepted			Status=3

	StatusChangesRequested	Status=2
	StatusSubmitted			Status=1
	StatusUnderConstruction	Status=0
)

type Submission struct {
	ID             int64  `gorm:"primaryKey"`
	DisplayName    string
	Creator        string
	GameID         int32
	CreatedAt time.Time
	UpdatedAt time.Time
	Submitter      int64 // UserID
	AssetID        int64
	AssetVersion   int64
	Completed      bool // Has this version of the map been completed at least once on maptest
	TargetAssetID  int64 // where to upload map fix.  if the TargetAssetID is 0, it's a new map.
	StatusID       Status
}
