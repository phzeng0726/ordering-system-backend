package repository

import (
	"context"
	"errors"
	"ordering-system-backend/internal/domain"
	"strings"

	"gorm.io/gorm"
)

type StoreMenusRepo struct {
	db *gorm.DB
	rt *RepoTools
}

func NewStoreMenusRepo(db *gorm.DB, rt *RepoTools) *StoreMenusRepo {
	return &StoreMenusRepo{
		db: db,
		rt: rt,
	}
}

func (r *StoreMenusRepo) checkUserStoreMenuExist(tx *gorm.DB, userId string, storeMenuMapping domain.StoreMenuMapping) error {
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

func (r *StoreMenusRepo) CreateMenuReference(ctx context.Context, userId string, storeMenuMapping domain.StoreMenuMapping) error {
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

func (r *StoreMenusRepo) UpdateMenuReference(ctx context.Context, userId string, storeMenuMapping domain.StoreMenuMapping) error {
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

func (r *StoreMenusRepo) DeleteMenuReference(ctx context.Context, userId string, storeId string) error {
	db := r.db.WithContext(ctx)

	if err := db.Transaction(func(tx *gorm.DB) error {
		// TODO 是否有需要留著userId
		// 刪除Reference
		if err := tx.Where("store_id = ?", storeId).Delete(&domain.StoreMenuMapping{}).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (r *StoreMenusRepo) GetMenuByStoreId(ctx context.Context, userId string, storeId string, languageId int) ([]domain.MenuItemMapping, error) {
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
		return menuItemMappings, errors.New("menu not found")
	}

	return menuItemMappings, nil
}
