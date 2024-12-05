package model

type Script struct {
	ID             int64
	Hash           uint64
	Source         string
	SubmissionID   int64 // which submission did this script first appear in
}
