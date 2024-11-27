package datastore

import (
	"context"
	"errors"
	"git.itzana.me/strafesnet/maps-service/pkg/model"
)

var (
	ErrNotExist = errors.New("resource does not exist")
)

type Datastore interface {
	Submissions() Submissions
}

type Submissions interface {
	Get(ctx context.Context, id int64) (model.Submission, error)
	GetList(ctx context.Context, id []int64) ([]model.Submission, error)
	Create(ctx context.Context, smap model.Submission) (model.Submission, error)
	Update(ctx context.Context, id int64, values OptionalMap) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, filters OptionalMap, page model.Page) ([]model.Submission, error)
}
