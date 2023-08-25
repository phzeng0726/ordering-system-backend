package repository

import (
	"errors"
	"ordering-system-backend/domain"

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

func insertMenu(db *gorm.DB, m domain.Menu) (int, error) {
	res := db.Create(&m)
	if res.Error != nil {
		return 0, res.Error
	}
	return m.Id, nil
}

func insertMenuItem(db *gorm.DB, mi domain.MenuItem) (int, error) {
	res := db.Create(&mi)
	if res.Error != nil {
		return 0, res.Error
	}
	return mi.Id, nil
}

func insertMenuItemMapping(db *gorm.DB, menuId int, menuItemId int) error {
	mapping := domain.MenuItemMapping{
		MenuId:     menuId,
		MenuItemId: menuItemId,
	}
	res := db.Create(&mapping)
	return res.Error
}

func (r *MenusRepo) Create(m domain.Menu) error {
	tx := r.db.Begin() // 開始事務
	if tx.Error != nil {
		return tx.Error
	}
	defer tx.Rollback()

	menuId, err := insertMenu(tx, m)
	if err != nil {
		return err
	}

	for _, mi := range m.MenuItems {
		menuItemId, err := insertMenuItem(tx, mi)
		if err != nil {
			return err
		}

		err = insertMenuItemMapping(tx, menuId, menuItemId)
		if err != nil {
			return err
		}
	}

	// 提交
	if err := tx.Commit(); err != nil {
		return err.Error
	}
	return nil
}

func updateMenus(db *gorm.DB, m domain.Menu) error {
	res := db.Model(&domain.Menu{}).Where("id = ?", m.Id).
		Updates(map[string]interface{}{
			"title":       m.Title,
			"description": m.Description,
			"is_hide":     m.IsHide,
		})
	return res.Error
}

func deleteMenuItemId(db *gorm.DB, r *MenusRepo, menuId int) error {
	var menuItemIds []string

	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer tx.Rollback()

	res := db.Where("menu_id = ?", menuId).Delete(&domain.MenuItemMapping{})
	if res.Error != nil {
		return res.Error
	}

	res = r.db.Model(&domain.MenuItemMapping{}).Select("menu_item_id").Where("menu_id = ?", menuId).Find(&menuItemIds)
	if res.Error != nil {
		return res.Error
	}
	res = db.Delete(&domain.MenuItem{}, menuItemIds)
	if res.Error != nil {
		return res.Error
	}

	// 提交
	if err := tx.Commit(); err != nil {
		return err.Error
	}

	return nil
}

func (r *MenusRepo) Update(m domain.Menu) error {
	// 開始 SQL Transaction
	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer tx.Rollback()

	err := updateMenus(tx, m)
	if err != nil {
		return err
	}

	err = deleteMenuItemId(tx, r, m.Id)
	if err != nil {
		return err
	}

	for _, mi := range m.MenuItems {
		menuItemId, err := insertMenuItem(tx, mi)
		if err != nil {
			return err
		}

		err = insertMenuItemMapping(tx, m.Id, menuItemId)
		if err != nil {
			return err
		}
	}

	// 提交
	if err := tx.Commit(); err != nil {
		return err.Error
	}

	return nil
}

// TODO 要先確定菜單所有人是storeId
func (r *MenusRepo) Delete(storeId string, menuId int) error {
	// 開始 GORM 事務
	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer tx.Rollback()

	// 刪除 menu_item_mapping 記錄
	if err := tx.Where("menu_id = ?", menuId).Delete(&domain.MenuItemMapping{}).Error; err != nil {
		return err
	}

	// 刪除 menu 記錄
	if err := tx.Where("id = ?", menuId).Delete(&domain.Menu{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 提交
	if err := tx.Commit(); err != nil {
		return err.Error
	}
	return nil
}

func (r *MenusRepo) GetAll(storeId string) ([]domain.Menu, error) {
	var menus []domain.Menu
	result := r.db.Where("store_id = ?", storeId).Find(&menus)
	if result.Error != nil {
		return nil, result.Error
	}
	return menus, nil
}

func (r *MenusRepo) GetById(storeId string, menuId int) (domain.Menu, error) {
	var menu domain.Menu
	var menuItemMappings []domain.MenuItemMapping
	if err := r.db.Preload("Menu", "store_id = ?", storeId).Preload("MenuItem.MenuCategory").Where("menu_id = ?", menuId).Find(&menuItemMappings).Error; err != nil {
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
