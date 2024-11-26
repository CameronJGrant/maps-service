package datastore

import (
	"context"
	"errors"
	"time"

	"git.itzana.me/strafesnet/data-service/internal/model"
)

var (
	ErrNotExist = errors.New("resource does not exist")
)

type Datastore interface {
	Times() Times
	Users() Users
	Bots() Bots
	Maps() Maps
	Events() Events
	Servers() Servers
	Transactions() Transactions
	Ranks() Ranks
}

type Times interface {
	Get(ctx context.Context, id int64) (model.Time, error)
	Create(ctx context.Context, time model.Time) (model.Time, error)
	Update(ctx context.Context, id int64, values OptionalMap) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, filters OptionalMap, blacklisted bool, page model.Page, sort uint32) (int64, []model.Time, error)
	Rank(ctx context.Context, id int64) (int64, error)
	DistinctStylePairs(ctx context.Context) ([]model.Time, error)
}

type Users interface {
	Get(ctx context.Context, id int64) (model.User, error)
	GetList(ctx context.Context, id []int64) ([]model.User, error)
	Create(ctx context.Context, user model.User) (model.User, error)
	Update(ctx context.Context, id int64, values OptionalMap) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, filters OptionalMap, page model.Page) ([]model.User, error)
}

type Bots interface {
	Get(ctx context.Context, id int64) (model.Bot, error)
	GetList(ctx context.Context, id []int64) ([]model.Bot, error)
	Create(ctx context.Context, bot model.Bot) (model.Bot, error)
	Update(ctx context.Context, id int64, values OptionalMap) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, filters OptionalMap, page model.Page) ([]model.Bot, error)
}

type Maps interface {
	Get(ctx context.Context, id int64) (model.Map, error)
	GetList(ctx context.Context, id []int64) ([]model.Map, error)
	Create(ctx context.Context, time model.Map) (model.Map, error)
	Update(ctx context.Context, id int64, values OptionalMap) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, filters OptionalMap, page model.Page) ([]model.Map, error)
}

type Events interface {
	Latest(ctx context.Context, date int64, page model.Page) ([]model.Event, error)
	Create(ctx context.Context, event model.Event) (model.Event, error)
	Clean(ctx context.Context) error
}

type Servers interface {
	Get(ctx context.Context, id string) (model.Server, error)
	Create(ctx context.Context, server model.Server) (model.Server, error)
	Update(ctx context.Context, id string, values OptionalMap) error
	Delete(ctx context.Context, id string) error
	DeleteByLastUpdated(ctx context.Context, date time.Time) error
	List(ctx context.Context, filters OptionalMap, page model.Page) ([]model.Server, error)
}

type Transactions interface {
	Balance(ctx context.Context, user int64) (int64, error)
	Get(ctx context.Context, id string) (model.Transaction, error)
	Create(ctx context.Context, transaction model.Transaction) (model.Transaction, error)
	Update(ctx context.Context, id string, values OptionalMap) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, filters OptionalMap, page model.Page) ([]model.Transaction, error)
}

type Ranks interface {
	Delete(ctx context.Context, id int64) error
	Get(ctx context.Context, user int64, style, game, mode int32, state []int32) (model.Rank, error)
	List(ctx context.Context, style, game, mode int32, sort int64, state []int32, page model.Page) ([]model.Rank, error)
	UpdateRankCalc(ctx context.Context) error
	UpdateAll(ctx context.Context, style, game, mode int32) error
	UpdateUsers(ctx context.Context, style, game, mode int32, users []int) error
}
