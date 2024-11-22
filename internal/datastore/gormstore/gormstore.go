package gormstore

import (
	"git.itzana.me/strafesnet/data-service/internal/datastore"
	"github.com/eko/gocache/lib/v4/cache"
	"gorm.io/gorm"
)

type Gormstore struct {
	db    *gorm.DB
	cache *cache.Cache[[]byte]
}

func (g Gormstore) Times() datastore.Times {
	return &Times{db: g.db}
}

func (g Gormstore) Users() datastore.Users {
	return &Users{db: g.db}
}

func (g Gormstore) Bots() datastore.Bots {
	return &Bots{db: g.db}
}

func (g Gormstore) Maps() datastore.Maps {
	return &Maps{db: g.db}
}

func (g Gormstore) Events() datastore.Events {
	return &Events{db: g.db}
}

func (g Gormstore) Servers() datastore.Servers {
	return &Servers{db: g.db}
}

func (g Gormstore) Transactions() datastore.Transactions {
	return &Transactions{db: g.db}
}

func (g Gormstore) Ranks() datastore.Ranks {
	return &Ranks{db: g.db, cache: g.cache}
}
