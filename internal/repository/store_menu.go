package repository

import (
	"context"
	"errors"
	"fmt"
	"ordering-system-backend/internal/domain"
	"strings"

	"gorm.io/gorm"
)

// 用來把store跟menu進行連接
type StoreMenusRepo struct {
	db *gorm.DB
}

func NewStoreMenusRepo(db *gorm.DB) *StoreMenusRepo {
	return &StoreMenusRepo{
		db: db,
	}
}

func (r *StoreMenusRepo) checkReferencePermission(tx *gorm.DB, userId string, storeMenuMapping domain.StoreMenuMapping) error {
	var store domain.Store
	var menu domain.Menu

	// 確認該User擁有此StoreId
	if err := tx.Where("user_id = ? AND id = ?", userId, storeMenuMapping.StoreId).First(&store).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// userId避免print在log上
			return fmt.Errorf("no store found with id [%s] for this user id", storeMenuMapping.StoreId)
		}
		return err
	}

	// 確認該User擁有此MenuId
	if err := tx.Where("user_id = ? AND id = ?", userId, storeMenuMapping.MenuId).First(&menu).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// userId避免print在log上
			return fmt.Errorf("no menu found with id [%s] for this user id", storeMenuMapping.MenuId)
		}
		return err
	}

	return nil
}

func (r *StoreMenusRepo) CreateMenuReference(ctx context.Context, userId string, storeMenuMapping domain.StoreMenuMapping) error {
	db := r.db.WithContext(ctx)

	if err := db.Transaction(func(tx *gorm.DB) error {
		// 確認權限
		if err := r.checkReferencePermission(tx, userId, storeMenuMapping); err != nil {
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
		// 確認權限
		if err := r.checkReferencePermission(tx, userId, storeMenuMapping); err != nil {
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

func (r *StoreMenusRepo) DeleteMenuReference(ctx context.Context, userId string, storeMenuMapping domain.StoreMenuMapping) error {
	db := r.db.WithContext(ctx)

	if err := db.Transaction(func(tx *gorm.DB) error {
		// 確認權限
		if err := r.checkReferencePermission(tx, userId, storeMenuMapping); err != nil {
			return err
		}

		// 刪除Reference
		if err := tx.Where("store_id = ? AND menu_id = ?", storeMenuMapping.StoreId, storeMenuMapping.MenuId).Delete(&domain.StoreMenuMapping{}).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
