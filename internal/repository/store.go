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

func (r *StoresRepo) getStoreWithStoreOwnerId(tx *gorm.DB, userId string, storeId string) (domain.Store, error) {
	var store domain.Store

	if err := tx.Where("user_id = ? AND id = ?", userId, storeId).First(&store).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// userId避免print在log上
			return store, fmt.Errorf("no store found with id %s for this user id", storeId)
		}
		return store, err
	}

	return store, nil
}

func (r *StoresRepo) deleteSeats(tx *gorm.DB, storeId string) error {
	return tx.Where("store_id = ?", storeId).Delete(&domain.Seat{}).Error
}

func (r *StoresRepo) deleteStoreOpeningHours(tx *gorm.DB, storeId string) error {
	return tx.Where("store_id = ?", storeId).Delete(&domain.StoreOpeningHour{}).Error
}

func (r *StoresRepo) createStoreOpeningHours(tx *gorm.DB, store domain.Store) error {
	for i := range store.StoreOpeningHours {
		store.StoreOpeningHours[i].StoreId = store.Id
	}

	return tx.Create(&store.StoreOpeningHours).Error
}

func (r *StoresRepo) Create(ctx context.Context, store domain.Store) error {
	db := r.db.WithContext(ctx)

	if err := db.Create(&store).Error; err != nil {
		return err
	}

	return nil
}

func (r *StoresRepo) Update(ctx context.Context, store domain.Store) error {
	db := r.db.WithContext(ctx)

	if err := db.Transaction(func(tx *gorm.DB) error {
		if _, err := r.getStoreWithStoreOwnerId(tx, store.UserId, store.Id); err != nil {
			return err
		}

		if err := r.deleteStoreOpeningHours(tx, store.Id); err != nil {
			return err
		}

		if len(store.StoreOpeningHours) != 0 {
			if err := r.createStoreOpeningHours(tx, store); err != nil {
				return err
			}
		}

		if err := tx.Model(&domain.Store{}).Where("user_id = ? AND id = ?", store.UserId, store.Id).Updates(&store).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (r *StoresRepo) Delete(ctx context.Context, userId string, storeId string) error {
	db := r.db.WithContext(ctx)

	if err := db.Transaction(func(tx *gorm.DB) error {
		if _, err := r.getStoreWithStoreOwnerId(tx, userId, storeId); err != nil {
			return err
		}

		if err := r.deleteSeats(tx, storeId); err != nil {
			return err
		}

		if err := r.deleteStoreOpeningHours(tx, storeId); err != nil {
			return err
		}

		if err := tx.Where("store_id = ?", storeId).Delete(&domain.StoreMenuMapping{}).Error; err != nil {
			return err
		}

		if err := tx.Where("user_id = ? AND id = ?", userId, storeId).Delete(&domain.Store{}).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (r *StoresRepo) GetAllByUserId(ctx context.Context, userId string) ([]domain.Store, error) {
	var stores []domain.Store
	db := r.db.WithContext(ctx)

	// 使用 Preload 一次載入所有的OpeningHours，避免N+1問題
	if err := db.Preload("StoreOpeningHours").Where("user_id = ?", userId).Find(&stores).Error; err != nil {
		return nil, err
	}

	return stores, nil
}

func (r *StoresRepo) GetByStoreId(ctx context.Context, userId string, storeId string) (domain.Store, error) {
	var store domain.Store
	db := r.db.WithContext(ctx)

	if err := db.Transaction(func(tx *gorm.DB) error {
		store, err := r.getStoreWithStoreOwnerId(tx, userId, storeId)
		if err != nil {
			return err
		}

		if err := tx.Where("store_id = ?", storeId).Find(&store.StoreOpeningHours).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return store, err
	}

	return store, nil
}

func (r *StoresRepo) GetAll(ctx context.Context) ([]domain.Store, error) {
	var stores []domain.Store
	db := r.db.WithContext(ctx)

	// 使用 Preload 一次載入所有的OpeningHours，避免N+1問題
	if err := db.Preload("StoreOpeningHours").Find(&stores).Error; err != nil {
		return nil, err
	}

	return stores, nil
}
