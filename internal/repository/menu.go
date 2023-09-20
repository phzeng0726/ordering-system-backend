package repository

import (
	"errors"
	"ordering-system-backend/internal/domain"

	"gorm.io/gorm"
)

type MenusRepo struct {
	db *gorm.DB
}

func NewMenusRepo(db *gorm.DB) *MenusRepo {
	return &MenusRepo{
		db: db,
	}
}

func (r *MenusRepo) Create(m domain.Menu) error {
	tx := r.db.Begin()
	defer func() {
		if tx.Error != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	if err := tx.Create(&m).Error; err != nil {
		return err
	}

	for _, mi := range m.MenuItems {
		if err := tx.Create(&mi).Error; err != nil {
			return err
		}
		mapping := domain.MenuItemMapping{
			MenuId:     m.Id,
			MenuItemId: mi.Id,
		}
		if err := tx.Create(&mapping).Error; err != nil {
			return err
		}
	}

	return nil
}

func (r *MenusRepo) Update(m domain.Menu) error {
	tx := r.db.Begin()
	defer func() {
		if tx.Error != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	if err := tx.Model(&domain.Menu{}).Where("id = ?", m.Id).Updates(map[string]interface{}{
		"title":       m.Title,
		"description": m.Description,
		"is_hide":     m.IsHide,
	}).Error; err != nil {
		return err
	}

	if err := tx.Where("menu_id = ?", m.Id).Delete(&domain.MenuItemMapping{}).Error; err != nil {
		return err
	}

	for _, mi := range m.MenuItems {
		if err := tx.Create(&mi).Error; err != nil {
			return err
		}
		mapping := domain.MenuItemMapping{
			MenuId:     m.Id,
			MenuItemId: mi.Id,
		}
		if err := tx.Create(&mapping).Error; err != nil {
			return err
		}
	}

	return nil
}

func (r *MenusRepo) Delete(storeId string, menuId int) error {
	tx := r.db.Begin()
	defer func() {
		if tx.Error != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	if err := tx.Where("menu_id = ?", menuId).Delete(&domain.MenuItemMapping{}).Error; err != nil {
		return err
	}

	if err := tx.Where("id = ?", menuId).Delete(&domain.Menu{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *MenusRepo) GetAll(storeId string) ([]domain.Menu, error) {
	var menus []domain.Menu
	res := r.db.Where("store_id = ?", storeId).Find(&menus)
	if res.Error != nil {
		return nil, res.Error
	}
	return menus, nil
}

func (r *MenusRepo) GetById(storeId string, menuId int) (domain.Menu, error) {
	var menu domain.Menu
	var menuItemMappings []domain.MenuItemMapping
	if err := r.db.Preload("Menu").Preload("MenuItem.MenuCategory").Where("menu_id = ?", menuId).Find(&menuItemMappings).Error; err != nil {
		return menu, err
	}

	if len(menuItemMappings) == 0 {
		err := errors.New("menu with items not found")
		return menu, err
	}

	menu = menuItemMappings[0].Menu
	for _, mim := range menuItemMappings {
		menu.MenuItems = append(menu.MenuItems, mim.MenuItem)
	}

	return menu, nil
}
