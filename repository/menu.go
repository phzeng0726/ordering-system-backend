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

// func scanMenusRow(rows *sql.Rows) ([]domain.Menu, error) {
// 	var menus []domain.Menu
// 	var err error

// 	for rows.Next() {
// 		var menu domain.Menu
// 		var createAtStr string // 創建一個字串來暫存日期時間字串

// 		err = rows.Scan(
// 			&menu.Id,
// 			&menu.StoreId,
// 			&menu.Title,
// 			&menu.Description,
// 			&menu.IsHide,
// 			&createAtStr, // 接收日期時間字串
// 		)
// 		if err != nil {
// 			return menus, err
// 		}

// 		menu.CreateAt, err = utils.DateTimeConverter(createAtStr)
// 		if err != nil {
// 			return menus, err
// 		}
// 		menus = append(menus, menu)
// 	}

// 	return menus, nil
// }

// func scanMenuByIdRow(rows *sql.Rows) (domain.Menu, error) {
// 	var menu domain.Menu
// 	var menuItems []domain.MenuItem
// 	var err error

// 	for rows.Next() {
// 		var menuItem domain.MenuItem
// 		var menuCategory domain.MenuCategory

// 		err = rows.Scan(
// 			&menu.Id,
// 			&menu.StoreId,
// 			&menu.Title,
// 			&menu.Description,
// 			&menu.IsHide,
// 			&menuItem.Id,
// 			&menuItem.Title,
// 			&menuItem.Description,
// 			&menuItem.Quantity,
// 			&menuItem.Price,
// 			&menuCategory.Id,
// 			&menuCategory.Title,
// 		)
// 		if err != nil {
// 			return menu, err
// 		}
// 		menuItem.MenuCategory = menuCategory
// 		menuItems = append(menuItems, menuItem)
// 	}

// 	menu.MenuItems = menuItems
// 	return menu, nil
// }

// func insertMenu(tx *sql.Tx, m domain.Menu) (int64, error) {
// 	sql := "INSERT INTO menu (store_id, title, `description`, is_hide, create_at)" +
// 		"VALUE (?, ?, ?, ?, ?)"
// 	res, err := tx.Exec(sql, m.StoreId, m.Title, m.Description, m.IsHide, m.CreateAt)
// 	if err != nil {
// 		return 0, err
// 	}

//		return res.LastInsertId()
//	}

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

// func insertMenuItem(tx *sql.Tx, mi domain.MenuItem) (int64, error) {
// 	fmt.Println(mi)
// 	sql := "INSERT INTO menu_item (menu_category_id, title, `description`, quantity, price)" +
// 		"VALUE (?, ?, ?, ?, ?)"
// 	res, err := tx.Exec(sql, mi.MenuCategory.Id, mi.Title, mi.Description, mi.Quantity, mi.Price)
// 	if err != nil {
// 		return 0, err
// 	}

// 	return res.LastInsertId()
// }

// func insertMenuItemMapping(tx *sql.Tx, menuId, menuItemId int64) error {
// 	sql := "INSERT INTO menu_item_mapping (menu_id, menu_item_id)" +
// 		"VALUE (?, ?)"
// 	_, err := tx.Exec(sql, menuId, menuItemId)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

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

// func deleteMenu(tx *sql.Tx, menuId int) error {
// 	sql := "DELETE FROM menu_item_mapping WHERE menu_id = ?"
// 	_, err := tx.Exec(sql, menuId)
// 	if err != nil {
// 		return err
// 	}

// 	sql = "DELETE FROM menu WHERE id = ?"

// 	_, err = tx.Exec(sql, menuId)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
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

func (r *MenusRepo) Create(m domain.Menu) error {
	// // 開始 SQL Transaction
	// tx, err := r.db.Begin()
	// if err != nil {
	// 	return err
	// }
	// defer tx.Rollback()

	// menuId, err := insertMenu(tx, m)
	// if err != nil {
	// 	return err
	// }

	// for _, mi := range m.MenuItems {
	// 	menuItemId, err := insertMenuItem(tx, mi)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	err = insertMenuItemMapping(tx, menuId, menuItemId)
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	// // 提交
	// if err := tx.Commit(); err != nil {
	// 	return err
	// }
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

	// sql := "UPDATE menu" +
	// 	" SET title=?, `description`=?, is_hide=?" +
	// 	" WHERE id=?"
	// _, err := tx.Exec(sql, m.Title, m.Description, m.IsHide, m.Id)
	// if err != nil {
	// 	return err
	// }

	return res.Error
}

// func getMappingMenuItemId(r *MenusRepo, menuId int) ([]int, error) {
// 	var menuItemIds []string

// 	res := r.db.Where("menu_id = ?", menuId).Find(&domain.MenuItemMapping{}).Pluck("menu_item_id", &menuItemIds)
// 	if res.Error != nil {
// 		return nil, res.Error
// 	}
// 	// sql := "SELECT menu_item_id" +
// 	// 	" FROM menu_item_mapping" +
// 	// 	" WHERE menu_id = ?"

// 	// rows, err := r.db.Query(sql, menuId)
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// 	return menuItemIds, err
// 	// }

// 	// defer rows.Close()

// 	// for rows.Next() {
// 	// 	var menuItemId int

// 	// 	err = rows.Scan(
// 	// 		&menuItemId,
// 	// 	)
// 	// 	if err != nil {
// 	// 		return menuItemIds, err
// 	// 	}
// 	// 	menuItemIds = append(menuItemIds, menuItemId)
// 	// }

// 	// if err != nil {
// 	// 	return menuItemIds, err
// 	// }

// 	return menuItemIds, nil
// }

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
	// sql := "DELETE FROM menu_item_mapping WHERE menu_id = ?"
	// _, err = tx.Exec(sql, menuId)
	// if err != nil {
	// 	return err
	// }

	// for _, mId := range menuItemIds {
	// 	sql = "DELETE FROM menu_item WHERE id = ?"
	// 	_, err = tx.Exec(sql, mId)
	// 	if err != nil {
	// 		return err
	// 	}
	// }

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

// func deleteMenu(tx *sql.Tx, menuId int) error {
// 	sql := "DELETE FROM menu_item_mapping WHERE menu_id = ?"
// 	_, err := tx.Exec(sql, menuId)
// 	if err != nil {
// 		return err
// 	}

// 	sql = "DELETE FROM menu WHERE id = ?"

// 	_, err = tx.Exec(sql, menuId)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

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
	// // 開始 SQL Transaction
	// tx, err := r.db.Begin()
	// if err != nil {
	// 	return err
	// }
	// defer tx.Rollback()
	// // TODO 要先確定菜單所有人是storeId
	// err = deleteMenu(tx, menuId)
	// if err != nil {
	// 	return err
	// }

	// // 提交
	// if err := tx.Commit(); err != nil {
	// 	return err
	// }
	return nil
}

func (r *MenusRepo) GetAll(storeId string) ([]domain.Menu, error) {
	// var menus []domain.Menu
	// sql := "SELECT *" +
	// 	" FROM menu" +
	// 	" WHERE store_id = ?"
	// rows, err := r.db.Query(sql, storeId)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return menus, err
	// }
	// defer rows.Close()

	// menus, err = scanMenusRow(rows)

	// if err != nil {
	// 	return menus, err
	// }
	var menus []domain.Menu
	result := r.db.Where("store_id = ?", storeId).Find(&menus)
	if result.Error != nil {
		return nil, result.Error
	}
	return menus, nil
}

// func scanMenuByIdRow(rows *sql.Rows) (domain.Menu, error) {
// 	var menu domain.Menu
// 	var menuItems []domain.MenuItem
// 	var err error

// 	for rows.Next() {
// 		var menuItem domain.MenuItem
// 		var menuCategory domain.MenuCategory

// 		err = rows.Scan(
// 			&menu.Id,
// 			&menu.StoreId,
// 			&menu.Title,
// 			&menu.Description,
// 			&menu.IsHide,
// 			&menuItem.Id,
// 			&menuItem.Title,
// 			&menuItem.Description,
// 			&menuItem.Quantity,
// 			&menuItem.Price,
// 			&menuCategory.Id,
// 			&menuCategory.Title,
// 		)
// 		if err != nil {
// 			return menu, err
// 		}
// 		menuItem.MenuCategory = menuCategory
// 		menuItems = append(menuItems, menuItem)
// 	}

//		menu.MenuItems = menuItems
//		return menu, nil
//	}
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

// rows, err := r.db.Query(sql, storeId, menuId)
// if err != nil {
// 	fmt.Println(err)
// 	return menu, err
// }

// defer rows.Close()

// menu, err = scanMenuByIdRow(rows)

// if err != nil {
// 	return menu, err
// }
// var menu domain.Menu
// func scanMenuByIdRow(rows *sql.Rows) (domain.Menu, error) {
// 	var menu domain.Menu
// 	var menuItems []domain.MenuItem
// 	var err error

// 	for rows.Next() {
// 		var menuItem domain.MenuItem
// 		var menuCategory domain.MenuCategory

// 		err = rows.Scan(
// 			&menu.Id,
// 			&menu.StoreId,
// 			&menu.Title,
// 			&menu.Description,
// 			&menu.IsHide,
// 			&menuItem.Id,
// 			&menuItem.Title,
// 			&menuItem.Description,
// 			&menuItem.Quantity,
// 			&menuItem.Price,
// 			&menuCategory.Id,
// 			&menuCategory.Title,
// 		)
// 		if err != nil {
// 			return menu, err
// 		}
// 		menuItem.MenuCategory = menuCategory
// 		menuItems = append(menuItems, menuItem)
// 	}

// 	menu.MenuItems = menuItems
// 	return menu, nil
// }

// func (r *MenusRepo) GetById(storeId string, menuId int) (domain.Menu, error) {
// 	var menu domain.Menu

// 	sql := "SELECT m.id menu_id, m.store_id, m.title, m.`description`, m.is_hide, mi.id , mi.title, mi.`description`, mi.quantity, mi.price, mc.id, mc.title" +
// 		" FROM menu m" +
// 		" JOIN menu_item_mapping mim ON m.id = mim.menu_id" +
// 		" JOIN menu_item mi ON mi.id = mim.menu_item_id" +
// 		" JOIN menu_category mc ON mi.menu_category_id = mc.id" +
// 		" WHERE m.store_id = ?" +
// 		" AND m.id = ?"

// 	rows, err := r.db.Query(sql, storeId, menuId)
// 	if err != nil {
// 		fmt.Println(err)
// 		return menu, err
// 	}

// 	defer rows.Close()

// 	menu, err = scanMenuByIdRow(rows)

// 	if err != nil {
// 		return menu, err
// 	}

// 	return menu, nil
// }
