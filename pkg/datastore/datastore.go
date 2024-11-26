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
}
