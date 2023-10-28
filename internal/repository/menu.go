package repository

import (
	"context"
	"errors"
	"ordering-system-backend/internal/domain"

	"gorm.io/gorm"
)

type MenusRepo struct {
	db *gorm.DB
	rt *RepoTools
}

func NewMenusRepo(db *gorm.DB, rt *RepoTools) *MenusRepo {
	return &MenusRepo{
		db: db,
		rt: rt,
	}
}

func (r *MenusRepo) updateMenu(tx *gorm.DB, menu domain.Menu) error {
	updatedMenu := map[string]interface{}{
		"title":       menu.Title,
		"description": menu.Description,
	}

	if err := tx.Model(&domain.Menu{}).Where("id = ?", menu.Id).Updates(updatedMenu).Error; err != nil {
		return err
	}

	return nil
}

func (r *MenusRepo) deleteMenuItems(tx *gorm.DB, menuId string) error {
	var menuItemIds []int
	if err := tx.Model(&domain.MenuItemMapping{}).Where("menu_id = ?", menuId).Pluck("menu_item_id", &menuItemIds).Error; err != nil {
		return err
	}

	if err := tx.Where("menu_id = ?", menuId).Delete(&domain.MenuItemMapping{}).Error; err != nil {
		return err
	}

	if err := tx.Where("id IN (?)", menuItemIds).Delete(&domain.MenuItem{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *MenusRepo) createMenuItems(tx *gorm.DB, menu domain.Menu) error {
	for _, mi := range menu.MenuItems {
		if err := tx.Create(&mi).Error; err != nil {
			return err
		}
		mapping := domain.MenuItemMapping{
			MenuId:     menu.Id,
			MenuItemId: mi.Id,
		}
		if err := tx.Create(&mapping).Error; err != nil {
			return err
		}
	}

	return nil
}

func (r *MenusRepo) Create(ctx context.Context, menu domain.Menu) error {
	db := r.db.WithContext(ctx)

	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := r.rt.CheckUserExist(tx, menu.UserId); err != nil {
			return err
		}

		if err := tx.Create(&menu).Error; err != nil {
			return err
		}

		if err := r.createMenuItems(tx, menu); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (r *MenusRepo) Update(ctx context.Context, menu domain.Menu) error {
	db := r.db.WithContext(ctx)

	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := r.updateMenu(tx, menu); err != nil {
			return err
		}

		if err := r.deleteMenuItems(tx, menu.Id); err != nil {
			return err
		}

		if err := r.createMenuItems(tx, menu); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (r *MenusRepo) Delete(ctx context.Context, userId string, menuId string) error {
	db := r.db.WithContext(ctx)

	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := r.deleteMenuItems(tx, menuId); err != nil {
			return err
		}

		if err := tx.Where("menu_id = ?", menuId).Delete(&domain.StoreMenuMapping{}).Error; err != nil {
			return err
		}

		if err := tx.Where("id = ?", menuId).Delete(&domain.Menu{}).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
func (r *MenusRepo) TempGetAllByUserId(ctx context.Context, userId string, languageId int) ([]domain.Menu, error) {
	var menus []domain.Menu
	db := r.db.WithContext(ctx)

	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := r.rt.CheckUserExist(tx, userId); err != nil {
			return err
		}

		if err := tx.Where("user_id = ?", userId).
			Find(&menus).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return menus, err
	}

	return menus, nil
}

func (r *MenusRepo) GetAllByUserId(ctx context.Context, userId string, languageId int) ([]domain.MenuItemMapping, error) {
	var menuItemMappings []domain.MenuItemMapping
	db := r.db.WithContext(ctx)

	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := r.rt.CheckUserExist(tx, userId); err != nil {
			return err
		}

		if err := tx.Preload("Menu").
			Preload("MenuItem.Image").
			Preload("MenuItem.Category").
			Preload("MenuItem.Category.CategoryLanguage", "language_id = ?", languageId).
			Where("menu_id IN (SELECT id FROM menus WHERE user_id = ?)", userId).
			Find(&menuItemMappings).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return menuItemMappings, err
	}

	return menuItemMappings, nil
}

func (r *MenusRepo) GetById(ctx context.Context, userId string, menuId string, languageId int) ([]domain.MenuItemMapping, error) {
	var menuItemMappings []domain.MenuItemMapping
	db := r.db.WithContext(ctx)

	if err := db.Preload("Menu").
		Preload("MenuItem.Image").
		Preload("MenuItem.Category").
		Preload("MenuItem.Category.CategoryLanguage", "language_id = ?", languageId).
		Where("menu_id = ?", menuId).Find(&menuItemMappings).Error; err != nil {
		return menuItemMappings, err
	}

	if len(menuItemMappings) == 0 {
		return menuItemMappings, errors.New("menu not found")
	}

	return menuItemMappings, nil
}
