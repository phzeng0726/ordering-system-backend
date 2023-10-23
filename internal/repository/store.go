package repository

import (
	"context"
	"errors"
	"ordering-system-backend/internal/domain"
	"strings"

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

	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := r.rt.CheckStoreExist(tx, store.UserId, store.Id, nil); err != nil {
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
		if err := r.rt.CheckStoreExist(tx, userId, storeId, nil); err != nil {
			return err
		}

		if err := r.deleteSeats(tx, storeId); err != nil {
			return err
		}

		if err := r.deleteStoreOpeningHours(tx, storeId); err != nil {
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
		if err := r.rt.CheckStoreExist(tx, userId, storeId, &store); err != nil {
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

// Store Menu Reference

func (r *StoresRepo) checkUserStoreMenuExist(tx *gorm.DB, userId string, storeMenuMapping domain.StoreMenuMapping) error {
	// 確認該User存在
	if err := r.rt.CheckUserExist(tx, userId); err != nil {
		return err
	}

	// 確認該User擁有此StoreId
	if err := r.rt.CheckStoreExist(tx, userId, storeMenuMapping.StoreId, nil); err != nil {
		return err
	}

	// 確認該User擁有此MenuId
	if err := r.rt.CheckMenuExist(tx, userId, storeMenuMapping.MenuId, nil); err != nil {
		return err
	}

	return nil
}

func (r *StoresRepo) CreateMenuReference(ctx context.Context, userId string, storeMenuMapping domain.StoreMenuMapping) error {
	db := r.db.WithContext(ctx)

	if err := db.Transaction(func(tx *gorm.DB) error {
		// 確認user、store、menu存在
		if err := r.checkUserStoreMenuExist(tx, userId, storeMenuMapping); err != nil {
			return err
		}

		// 新增Reference
		if err := tx.Create(&storeMenuMapping).Error; err != nil {
			if strings.Contains(err.Error(), "store_id_UNIQUE") {
				return errors.New("reference has already existed")
			}

			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (r *StoresRepo) UpdateMenuReference(ctx context.Context, userId string, storeMenuMapping domain.StoreMenuMapping) error {
	db := r.db.WithContext(ctx)

	if err := db.Transaction(func(tx *gorm.DB) error {
		// 確認user、store、menu存在
		if err := r.checkUserStoreMenuExist(tx, userId, storeMenuMapping); err != nil {
			return err
		}

		// 新增Reference
		if err := tx.Model(&domain.StoreMenuMapping{}).Where("store_id = ?", storeMenuMapping.StoreId).Updates(&storeMenuMapping).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (r *StoresRepo) GetMenuByStoreId(ctx context.Context, userId string, storeId string, languageId int) ([]domain.MenuItemMapping, error) {
	var storeMenuMapping domain.StoreMenuMapping
	var menuItemMappings []domain.MenuItemMapping
	db := r.db.WithContext(ctx)

	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("store_id = ?", storeId).First(&storeMenuMapping).Error; err != nil {
			return err
		}

		if err := tx.Preload("Menu").
			Preload("MenuItem.Image").
			Preload("MenuItem.Category").
			Preload("MenuItem.Category.CategoryLanguage", "language_id = ?", languageId).
			Where("menu_id = ?", storeMenuMapping.MenuId).Find(&menuItemMappings).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return menuItemMappings, err
	}

	if len(menuItemMappings) == 0 {
		return menuItemMappings, errors.New("menu with items not found")
	}

	return menuItemMappings, nil
}
