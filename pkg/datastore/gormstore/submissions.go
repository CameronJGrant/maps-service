package gormstore

import (
	"context"
	"errors"

	"git.itzana.me/strafesnet/maps-service/pkg/datastore"
	"git.itzana.me/strafesnet/maps-service/pkg/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	ErrorStatus = errors.New("Status is not in allowed statuses")
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

func (env *Submissions) GetList(ctx context.Context, id []int64) ([]model.Submission, error) {
	var mapList []model.Submission
	if err := env.db.Find(&mapList, "id IN ?", id).Error; err != nil {
		return mapList, err
	}

	return mapList, nil
}

func (env *Submissions) Create(ctx context.Context, smap model.Submission) (model.Submission, error) {
	if err := env.db.Create(&smap).Error; err != nil {
		return smap, err
	}

	return smap, nil
}

func (env *Submissions) Update(ctx context.Context, id int64, values datastore.OptionalMap) error {
	if err := env.db.Model(&model.Submission{}).Where("id = ?", id).Updates(values.Map()).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return datastore.ErrNotExist
		}
		return err
	}

	return nil
}

// the update can only occur if the status matches one of the provided values.
func (env *Submissions) IfStatusThenUpdate(ctx context.Context, id int64, statuses []model.Status, values datastore.OptionalMap) error {
	if err := env.db.Model(&model.Submission{}).Where("id = ?", id).Where("status_id IN ?",statuses).Updates(values.Map()).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return datastore.ErrNotExist
		}
		return err
	}

	return nil
}

// the update can only occur if the status matches one of the provided values.
// returns the updated value
func (env *Submissions) IfStatusThenUpdateAndGet(ctx context.Context, id int64, statuses []model.Status, values datastore.OptionalMap) (model.Submission, error) {
	var submission model.Submission
	result := env.db.Model(&submission).
		Clauses(clause.Returning{}).
		Where("id = ?", id).
		Where("status_id IN ?",statuses).
		Updates(values.Map())
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return submission, datastore.ErrNotExist
		}
		return submission, result.Error
	}

	if result.RowsAffected == 0 {
		return submission, ErrorStatus
	}

	return submission, nil
}

func (env *Submissions) Delete(ctx context.Context, id int64) error {
	if err := env.db.Delete(&model.Submission{}, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return datastore.ErrNotExist
		}
		return err
	}

	return nil
}

func (env *Submissions) List(ctx context.Context, filters datastore.OptionalMap, page model.Page) ([]model.Submission, error) {
	var maps []model.Submission
	if err := env.db.Where(filters.Map()).Offset(int((page.Number - 1) * page.Size)).Limit(int(page.Size)).Find(&maps).Error; err != nil {
		return nil, err
	}

	return maps, nil
}
