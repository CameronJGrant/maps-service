package gormstore

import (
	"context"
	"errors"
	"git.itzana.me/strafesnet/maps-service/pkg/datastore"
	"git.itzana.me/strafesnet/maps-service/pkg/model"
	"gorm.io/gorm"
)

type Submissions struct {
	db *gorm.DB
}

func (env *Submissions) Get(ctx context.Context, id int64) (model.Submission, error) {
	var submission model.Submission
	if err := env.db.First(&submission, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return submission, datastore.ErrNotExist
		}
	}
	return submission, nil
}
