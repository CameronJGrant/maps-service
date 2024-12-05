package gormstore

import (
	"context"
	"errors"

	"git.itzana.me/strafesnet/maps-service/pkg/datastore"
	"git.itzana.me/strafesnet/maps-service/pkg/model"
	"gorm.io/gorm"
)

type ScriptPolicy struct {
	db *gorm.DB
}

func (env *ScriptPolicy) Get(ctx context.Context, id int64) (model.ScriptPolicy, error) {
	var mdl model.ScriptPolicy
	if err := env.db.First(&mdl, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return mdl, datastore.ErrNotExist
		}
	}
	return mdl, nil
}

func (env *ScriptPolicy) Create(ctx context.Context, smap model.ScriptPolicy) (model.ScriptPolicy, error) {
	if err := env.db.Create(&smap).Error; err != nil {
		return smap, err
	}

	return smap, nil
}

func (env *ScriptPolicy) Update(ctx context.Context, id int64, values datastore.OptionalMap) error {
	if err := env.db.Model(&model.ScriptPolicy{}).Where("id = ?", id).Updates(values.Map()).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return datastore.ErrNotExist
		}
		return err
	}

	return nil
}

func (env *ScriptPolicy) Delete(ctx context.Context, id int64) error {
	if err := env.db.Delete(&model.ScriptPolicy{}, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return datastore.ErrNotExist
		}
		return err
	}

	return nil
}
