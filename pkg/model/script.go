package model

import "time"

type Script struct {
	ID             int64  `gorm:"primaryKey"`
	Hash           uint64
	Source         string
	SubmissionID   int64 // which submission did this script first appear in
	CreatedAt time.Time
	UpdatedAt time.Time
}
