package gormstore

import (
	"context"
	"git.itzana.me/strafesnet/data-service/internal/datastore"
	"git.itzana.me/strafesnet/data-service/internal/model"
	"gorm.io/gorm"
)

type Users struct {
	db *gorm.DB
}

func (u Users) Get(ctx context.Context, id int64) (model.User, error) {
	var user model.User
	if err := u.db.WithContext(ctx).First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return user, datastore.ErrNotExist
		}
		return user, err
	}

	return user, nil
}

func (u Users) GetList(ctx context.Context, id []int64) ([]model.User, error) {
	var user []model.User
	if err := u.db.WithContext(ctx).Find(&user, "id IN ?", id).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (u Users) Create(ctx context.Context, user model.User) (model.User, error) {
	if err := u.db.WithContext(ctx).Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (u Users) Update(ctx context.Context, id int64, values datastore.OptionalMap) error {
	if err := u.db.WithContext(ctx).Model(&model.User{}).Where("id = ?", id).Updates(values.Map()).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return datastore.ErrNotExist
		}
		return err
	}

	return nil
}

func (u Users) Delete(ctx context.Context, id int64) error {
	if err := u.db.WithContext(ctx).Delete(&model.User{}, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return datastore.ErrNotExist
		}
		return err
	}

	return nil
}

func (u Users) List(ctx context.Context, filters datastore.OptionalMap, page model.Page) ([]model.User, error) {
	var users []model.User
	if err := u.db.WithContext(ctx).Where(filters.Map()).Offset(int((page.Number - 1) * page.Size)).Limit(int(page.Size)).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
