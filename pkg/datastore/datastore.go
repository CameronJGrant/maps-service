package datastore

import (
	"context"
	"errors"
	"git.itzana.me/strafesnet/maps-service/pkg/model"
)

var (
	ErrNotExist = errors.New("resource does not exist")
	ErroNoRowsAffected = errors.New("query did not affect any rows")
)

type Datastore interface {
	Submissions() Submissions
	Scripts() Scripts
	ScriptPolicy() ScriptPolicy
}

type Submissions interface {
	Get(ctx context.Context, id int64) (model.Submission, error)
	GetList(ctx context.Context, id []int64) ([]model.Submission, error)
	Create(ctx context.Context, smap model.Submission) (model.Submission, error)
	Update(ctx context.Context, id int64, values OptionalMap) error
	IfStatusThenUpdate(ctx context.Context, id int64, statuses []model.Status, values OptionalMap) error
	IfStatusThenUpdateAndGet(ctx context.Context, id int64, statuses []model.Status, values OptionalMap) (model.Submission, error)
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, filters OptionalMap, page model.Page) ([]model.Submission, error)
}

type Scripts interface {
	Get(ctx context.Context, id int64) (model.Script, error)
	Create(ctx context.Context, smap model.Script) (model.Script, error)
	Update(ctx context.Context, id int64, values OptionalMap) error
	Delete(ctx context.Context, id int64) error
}

type ScriptPolicy interface {
	Get(ctx context.Context, id int64) (model.ScriptPolicy, error)
	GetFromHash(ctx context.Context, hash uint64) (model.ScriptPolicy, error)
	Create(ctx context.Context, smap model.ScriptPolicy) (model.ScriptPolicy, error)
	Update(ctx context.Context, id int64, values OptionalMap) error
	Delete(ctx context.Context, id int64) error
}
