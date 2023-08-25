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

// func updateMenus(tx *sql.Tx, m domain.Menu) error {
// 	sql := "UPDATE menu" +
// 		" SET title=?, `description`=?, is_hide=?" +
// 		" WHERE id=?"
// 	_, err := tx.Exec(sql, m.Title, m.Description, m.IsHide, m.Id)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func getMappingMenuItemId(r *MenusRepo, menuId int) ([]int, error) {
// 	var menuItemIds []int

// 	sql := "SELECT menu_item_id" +
// 		" FROM menu_item_mapping" +
// 		" WHERE menu_id = ?"

// 	rows, err := r.db.Query(sql, menuId)
// 	if err != nil {
// 		fmt.Println(err)
// 		return menuItemIds, err
// 	}

// 	defer rows.Close()

// 	for rows.Next() {
// 		var menuItemId int

// 		err = rows.Scan(
// 			&menuItemId,
// 		)
// 		if err != nil {
// 			return menuItemIds, err
// 		}
// 		menuItemIds = append(menuItemIds, menuItemId)
// 	}

// 	if err != nil {
// 		return menuItemIds, err
// 	}

// 	return menuItemIds, nil
// }

// func deleteMenuItemId(tx *sql.Tx, r *MenusRepo, menuId int) error {
// 	menuItemIds, err := getMappingMenuItemId(r, menuId)
// 	if err != nil {
// 		return err
// 	}

// 	sql := "DELETE FROM menu_item_mapping WHERE menu_id = ?"
// 	_, err = tx.Exec(sql, menuId)
// 	if err != nil {
// 		return err
// 	}

// 	for _, mId := range menuItemIds {
// 		sql = "DELETE FROM menu_item WHERE id = ?"
// 		_, err = tx.Exec(sql, mId)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

func (r *MenusRepo) Update(m domain.Menu) error {
	tx := r.db.Begin() // 開始事務
	if tx.Error != nil {
		return tx.Error
	}
	defer tx.Rollback()

	menuId, err := insertMenu(tx, m)
	// err = updateMenus(tx, m)
	// if err != nil {
	// 	return err
	// }

	// err = deleteMenuItemId(tx, r, m.Id)
	// if err != nil {
	// 	return err
	// }

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
