package model

import "time"

type Map struct {
	ID          int64
	DisplayName string
	Creator     string
	GameID      int32
	Date        time.Time
}
