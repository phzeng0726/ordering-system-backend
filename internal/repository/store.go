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
	rt *RepoTools
}

func NewStoresRepo(db *gorm.DB, rt *RepoTools) *StoresRepo {
	return &StoresRepo{
		db: db,
		rt: rt,
	}
}

func (r *StoresRepo) Create(ctx context.Context, store domain.Store) error {
	db := r.db.WithContext(ctx)

	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := r.rt.CheckUserExist(tx, store.UserId); err != nil {
			return err
		}

		// 確認該User存在，才可新增Store
		if err := tx.Create(&store).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (r *StoresRepo) Update(ctx context.Context, store domain.Store) error {
	db := r.db.WithContext(ctx)
	fmt.Println(store)

	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("user_id = ? AND id = ?", store.UserId, store.Id).First(&domain.Store{}).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// userId避免print在log上
				return fmt.Errorf("no store found with id %s for this user id", store.Id)
			}
			return err
		}

		if err := tx.Where("store_id = ?", store.Id).Delete(&domain.StoreOpeningHour{}).Error; err != nil {
			return err
		}

		for i := range store.StoreOpeningHours {
			store.StoreOpeningHours[i].StoreId = store.Id
		}

		if err := tx.Create(&store.StoreOpeningHours).Error; err != nil {
			return err
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
	var store domain.Store
	db := r.db.WithContext(ctx)

	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("user_id = ? AND id = ?", userId, storeId).First(&store).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// userId避免print在log上
				return fmt.Errorf("no store found with id %s for this user id", storeId)
			}
			return err
		}

		if err := tx.Where("store_id = ?", storeId).Delete(&domain.Seat{}).Error; err != nil {
			return err
		}

		if err := tx.Where("store_id = ?", storeId).Delete(&domain.StoreOpeningHour{}).Error; err != nil {
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
