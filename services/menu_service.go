package services

import (
	"database/sql"
	"fmt"
	"ordering-system-backend/db"
	"ordering-system-backend/models"
	"ordering-system-backend/utils"
)

func GetMenus(storeId string) ([]models.Menu, error) {
	var menus []models.Menu
	var createAtStr string // 創建一個字串來暫存日期時間字串

	sql := "SELECT *" +
		" FROM menu" +
		" WHERE store_id = ?"
	rows, err := db.DB.Query(sql, storeId)

	if err != nil {
		fmt.Println(err)
		return menus, err
	}
	defer rows.Close()

	for rows.Next() {
		var menu models.Menu
		err := rows.Scan(
			&menu.Id,
			&menu.StoreId,
			&menu.Title,
			&menu.Description,
			&menu.IsHide,
			&createAtStr, // 接收日期時間字串
		)
		if err != nil {
			fmt.Println(err)
			return menus, err
		}

		menu.CreateAt, err = utils.DateTimeConverter(createAtStr)
		if err != nil {
			fmt.Println(err)
			return menus, err
		}
		menus = append(menus, menu)
	}

	return menus, err
}

func insertMenu(tx *sql.Tx, m models.Menu) (int64, error) {
	res, err := tx.Exec(
		"INSERT INTO menu (store_id, title, `description`, is_hide, create_at)"+
			"VALUE (?, ?, ?, ?, ?)", m.StoreId, m.Title, m.Description, m.IsHide, m.CreateAt,
	)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func insertMenuItem(tx *sql.Tx, mi models.MenuItem) (int64, error) {
	res, err := tx.Exec(
		"INSERT INTO menu_item (menu_category_id, title, `description`, quantity, price)"+
			"VALUE (?, ?, ?, ?, ?)", mi.MenuCategoryId, mi.Title, mi.Description, mi.Quantity, mi.Price,
	)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func insertMenuItemMapping(tx *sql.Tx, menuId, menuItemId int64) error {
	_, err := tx.Exec(
		"INSERT INTO menu_item_mapping (menu_id, menu_item_id)"+
			"VALUE (?, ?)", menuId, menuItemId,
	)
	return err
}

func CreateMenus(m models.Menu) error {
	// 開始 SQL Transaction
	tx, err := db.DB.Begin()
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer tx.Rollback()

	menuId, err := insertMenu(tx, m)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(menuId)
	for _, mi := range m.MenuItems {
		menuItemId, err := insertMenuItem(tx, mi)
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Println(menuItemId)

		err = insertMenuItemMapping(tx, menuId, menuItemId)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	// 提交事務
	if err := tx.Commit(); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
