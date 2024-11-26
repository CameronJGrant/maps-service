package model

type User struct {
	ID       int64
	Username string `gorm:"not null"`
	StateID  int32  `gorm:"not null;default:0"`
}
