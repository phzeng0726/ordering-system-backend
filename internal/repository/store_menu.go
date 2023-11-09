package repository

import (
	"context"
	"errors"
	"fmt"
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
	if err := r.rt.CheckUserStoreExist(tx, userId, storeMenuMapping.StoreId, nil); err != nil {
		return err
	}

	// 確認該User擁有此MenuId
	if err := r.rt.CheckUserMenuExist(tx, userId, storeMenuMapping.MenuId, nil); err != nil {
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

func (r *StoreMenusRepo) GetMenuByStoreId(ctx context.Context, userId string, storeId string, languageId int, userType int) (domain.Menu, error) {
	var menu domain.Menu
	var storeMenuMapping domain.StoreMenuMapping
	var menuItemMappings []domain.MenuItemMapping
	db := r.db.WithContext(ctx)

	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("store_id = ?", storeId).First(&storeMenuMapping).Error; err != nil {
			return err
		}

		// 以menuId下去撈menuItems
		if err := tx.Preload("Menu").
			Preload("MenuItem.Image").
			Preload("MenuItem.Category").
			Preload("MenuItem.Category.CategoryLanguage", "language_id = ?", languageId).
			Where("menu_id = ?", storeMenuMapping.MenuId).Find(&menuItemMappings).Error; err != nil {
			fmt.Print("hi")
			return err

		}

		// 找不到menuItems的話，只回傳menu
		if len(menuItemMappings) == 0 {
			if err := tx.Where("id = ?", storeMenuMapping.MenuId).
				First(&menu).Error; err != nil {
				return err
			}

			menu.MenuItems = []domain.MenuItem{}
		} else {
			// 否則處理menuItems為格式化資料
			menu = menuItemMappings[0].Menu
			for _, mim := range menuItemMappings {
				mim.MenuItem.ImageBytes = mim.MenuItem.Image.BytesData
				mim.MenuItem.Category.Title = mim.MenuItem.Category.CategoryLanguage.Title
				menu.MenuItems = append(menu.MenuItems, mim.MenuItem)
			}

		}

		// 撈取商店資訊，供客戶端使用
		if userType == 1 {
			var store domain.Store
			if err := r.rt.GetStoreInfo(tx, storeId, &store); err != nil {
				return err
			}
			menu.Store = &store
		}

		return nil
	}); err != nil {
		return menu, err
	}

	return menu, nil
}
