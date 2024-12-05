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

func (g Gormstore) Scripts() datastore.Scripts {
	return &Scripts{db: g.db}
}

func (g Gormstore) ScriptPolicy() datastore.ScriptPolicy {
	return &ScriptPolicy{db: g.db}
}
