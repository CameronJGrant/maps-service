package gormstore

import (
	"context"
	"errors"

	"git.itzana.me/strafesnet/maps-service/pkg/datastore"
	"git.itzana.me/strafesnet/maps-service/pkg/model"
	"gorm.io/gorm"
)

type Scripts struct {
	db *gorm.DB
}

func (env *Scripts) Get(ctx context.Context, id int64) (model.Script, error) {
	var mdl model.Script
	if err := env.db.First(&mdl, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return mdl, datastore.ErrNotExist
		}
	}
	return mdl, nil
}

func (env *Scripts) GetList(ctx context.Context, id []int64) ([]model.Script, error) {
	var mapList []model.Script
	if err := env.db.Find(&mapList, "id IN ?", id).Error; err != nil {
		return mapList, err
	}

	return mapList, nil
}

func (env *Scripts) Create(ctx context.Context, smap model.Script) (model.Script, error) {
	if err := env.db.Create(&smap).Error; err != nil {
		return smap, err
	}

	return smap, nil
}

func (env *Scripts) Update(ctx context.Context, id int64, values datastore.OptionalMap) error {
	if err := env.db.Model(&model.Script{}).Where("id = ?", id).Updates(values.Map()).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return datastore.ErrNotExist
		}
		return err
	}

	return nil
}

// the update can only occur if the status matches one of the provided values.
func (env *Scripts) IfStatusThenUpdate(ctx context.Context, id int64, statuses []model.Status, values datastore.OptionalMap) error {
	if err := env.db.Model(&model.Script{}).Where("id = ?", id).Where("status IN ?",statuses).Updates(values.Map()).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return datastore.ErrNotExist
		}
		return err
	}

	return nil
}

func (env *Scripts) Delete(ctx context.Context, id int64) error {
	if err := env.db.Delete(&model.Script{}, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return datastore.ErrNotExist
		}
		return err
	}

	return nil
}

func (env *Scripts) List(ctx context.Context, filters datastore.OptionalMap, page model.Page) ([]model.Script, error) {
	var maps []model.Script
	if err := env.db.Where(filters.Map()).Offset(int((page.Number - 1) * page.Size)).Limit(int(page.Size)).Find(&maps).Error; err != nil {
		return nil, err
	}

	return maps, nil
}
