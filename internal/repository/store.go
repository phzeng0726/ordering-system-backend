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

func (r *StoresRepo) Create(ctx context.Context, store domain.Store) error {
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", store.UserId).First(&domain.User{}).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("user id not found")
			}
			return err
		}

		// 確認該User存在，才可新增Store
		if err := tx.Create(&store).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *StoresRepo) Update(ctx context.Context, store domain.Store) error {
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("user_id = ? AND id = ?", store.UserId, store.Id).First(&store).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// userId避免print在log上
				return fmt.Errorf("no store found with id %s for this user id", store.Id)
			}
			return err
		}

		if err := tx.Model(&domain.Store{}).Where("user_id = ? AND id = ?", store.UserId, store.Id).Updates(&store).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *StoresRepo) Delete(ctx context.Context, userId string, storeId string) error {
	var store domain.Store

	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("user_id = ? AND id = ?", userId, storeId).First(&store).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// userId避免print在log上
				return fmt.Errorf("no store found with id %s for this user id", storeId)
			}
			return err
		}

		if err := tx.Where("store_id = ?", storeId).Delete(&domain.StoreOpeningHour{}).Error; err != nil {
			return err
		}

		if err := tx.Where("user_id = ? AND id = ?", userId, storeId).Delete(&domain.Store{}).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *StoresRepo) GetAllByUserId(ctx context.Context, userId string) ([]domain.Store, error) {
	var stores []domain.Store

	// 使用 Preload 一次載入所有的OpeningHours，避免N+1問題
	if err := r.db.WithContext(ctx).Preload("StoreOpeningHours").Where("user_id = ?", userId).Find(&stores).Error; err != nil {
		return nil, err
	}

	return stores, nil
}

func (r *StoresRepo) GetByStoreId(ctx context.Context, userId string, storeId string) (domain.Store, error) {
	var store domain.Store

	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		res := tx.Where("user_id = ? AND id = ?", userId, storeId).First(&store)

		if res.Error != nil {
			if errors.Is(res.Error, gorm.ErrRecordNotFound) {
				return errors.New("store not found")
			}
			return res.Error
		}

		if err := tx.Where("store_id = ?", storeId).Find(&store.StoreOpeningHours).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return store, err
	}

	return store, nil
}

func (r *StoresRepo) GetAll(ctx context.Context) ([]domain.Store, error) {
	var stores []domain.Store

	// 使用 Preload 一次載入所有的OpeningHours，避免N+1問題
	if err := r.db.WithContext(ctx).Preload("StoreOpeningHours").Find(&stores).Error; err != nil {
		return nil, err
	}

	return stores, nil
}
