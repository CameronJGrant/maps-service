package gormstore

import (
	"context"
	"git.itzana.me/strafesnet/maps-service/internal/datastore"
	"git.itzana.me/strafesnet/maps-service/internal/model"
	"gorm.io/gorm"
)

type Maps struct {
	db *gorm.DB
}

func (m Maps) Get(ctx context.Context, id int64) (model.Map, error) {
	var smap model.Map
	if err := m.db.WithContext(ctx).First(&smap, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return smap, datastore.ErrNotExist
		}
		return smap, err
	}

	return smap, nil
}

func (m Maps) GetList(ctx context.Context, id []int64) ([]model.Map, error) {
	var mapList []model.Map
	if err := m.db.WithContext(ctx).Find(&mapList, "id IN ?", id).Error; err != nil {
		return mapList, err
	}

	return mapList, nil
}

func (m Maps) Create(ctx context.Context, smap model.Map) (model.Map, error) {
	if err := m.db.WithContext(ctx).Create(&smap).Error; err != nil {
		return smap, err
	}

	return smap, nil
}

func (m Maps) Update(ctx context.Context, id int64, values datastore.OptionalMap) error {
	if err := m.db.WithContext(ctx).Model(&model.Map{}).Where("id = ?", id).Updates(values.Map()).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return datastore.ErrNotExist
		}
		return err
	}

	return nil
}

func (m Maps) Delete(ctx context.Context, id int64) error {
	if err := m.db.WithContext(ctx).Delete(&model.Map{}, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return datastore.ErrNotExist
		}
		return err
	}

	return nil
}

func (m Maps) List(ctx context.Context, filters datastore.OptionalMap, page model.Page) ([]model.Map, error) {
	var maps []model.Map
	if err := m.db.WithContext(ctx).Where(filters.Map()).Offset(int((page.Number - 1) * page.Size)).Limit(int(page.Size)).Find(&maps).Error; err != nil {
		return nil, err
	}

	return maps, nil
}
