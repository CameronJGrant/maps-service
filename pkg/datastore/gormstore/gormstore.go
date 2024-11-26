package gormstore

import (
	"git.itzana.me/strafesnet/maps-service/pkg/datastore"
	"gorm.io/gorm"
)

type Gormstore struct {
	db *gorm.DB
}

func (g Gormstore) Submissions() datastore.Submissions {
	return &Submissions{db: g.db}
}
