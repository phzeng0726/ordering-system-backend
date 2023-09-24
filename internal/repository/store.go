package repository

import (
	"context"
	"errors"
	"fmt"
	"ordering-system-backend/internal/domain"

	"gorm.io/gorm"
)

type StoresRepo struct {
	db *gorm.DB
}

func NewStoresRepo(db *gorm.DB) *StoresRepo {
	return &StoresRepo{
		db: db,
	}
}

func (r *StoresRepo) Create(s domain.Store) error {
	res := r.db.Where("id = ?", s.UserId).First(&domain.User{})
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return errors.New("user id not found")
		}
		return res.Error
	}

	// 確認該User存在，才可新增Store
	if err := r.db.Create(&s).Error; err != nil {
		return err
	}

	return nil
}

func (r *StoresRepo) Update(userId string, s domain.Store) error {
	var store domain.Store
	res := r.db.Where("user_id = ?", userId).Where("id = ?", store.Id).First(&store)

	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			// userId避免print在log上
			return fmt.Errorf("no store found with id %s for this user id", s.Id)
		}
		return res.Error
	}

	res = r.db.Model(&store).Where("user_id = ?", userId).Where("id = ?", store.Id).Updates(&s)

	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (r *StoresRepo) Delete(ctx context.Context, userId string, storeId string) error {
	var store domain.Store

	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		res := tx.Where("id = ?", storeId).Where("user_id = ?", userId).First(&store)
		if res.Error != nil {
			if errors.Is(res.Error, gorm.ErrRecordNotFound) {
				// userId避免print在log上
				return fmt.Errorf("no store found with id %s for this user id", storeId)
			}
			return res.Error
		}

		if err := tx.Where("store_id = ?", storeId).Delete(&domain.StoreOpeningHour{}).Error; err != nil {
			return err
		}

		if err := tx.Where("user_id = ?", userId).Where("id = ?", storeId).Delete(&domain.Store{}).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *StoresRepo) GetAllByUserId(userId string) ([]domain.Store, error) {
	var stores []domain.Store

	// 使用 Preload 一次載入所有的OpeningHours，避免N+1問題
	if err := r.db.Preload("StoreOpeningHours").Where("user_id = ?", userId).Find(&stores).Error; err != nil {
		return nil, err
	}

	return stores, nil
}

func (r *StoresRepo) GetByStoreId(userId string, storeId string) (domain.Store, error) {
	var store domain.Store
	res := r.db.Where("user_id = ?", userId).Where("id = ?", storeId).First(&store)

	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return store, errors.New("store not found")
		}
		return store, res.Error
	}

	if err := r.db.Where("store_id = ?", storeId).Find(&store.StoreOpeningHours).Error; err != nil {
		return store, err
	}

	return store, nil
}

func (r *StoresRepo) GetAll() ([]domain.Store, error) {
	var stores []domain.Store

	// 使用 Preload 一次載入所有的OpeningHours，避免N+1問題
	if err := r.db.Preload("StoreOpeningHours").Find(&stores).Error; err != nil {
		return nil, err
	}

	return stores, nil
}
