package services

import (
	"database/sql"
	"fmt"
	"log"
	"ordering-system-backend/db"
	"ordering-system-backend/models"
	"ordering-system-backend/utils"
)

func GetMenus(storeId string) ([]models.Menu, error) {
	sql := "SELECT *" +
		" FROM menu" +
		" WHERE store_id = ?"
	rows, err := db.DB.Query(sql, storeId)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var menus []models.Menu
	var createAtStr string // 創建一個字串來暫存日期時間字串

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
			log.Fatal(err)
		}

		menu.CreateAt, err = utils.DateTimeConverter(createAtStr)
		if err != nil {
			log.Fatal(err)
		}
		menus = append(menus, menu)
	}
	fmt.Println(menus)

	return menus, err
}

func CreateMenus(m models.Menu) error {

	// 開始 SQL Transaction
	var res sql.Result
	var menuId int64
	var menuItemId int64

	tx, err := db.DB.Begin()

	if err != nil {
		log.Fatal(err)
	}

	res, err = tx.Exec(
		"INSERT INTO menu (store_id, title, `description`, is_hide, create_at)"+
			"VALUE (?, ?, ?, ?, ?)", m.StoreId, m.Title, m.Description, m.IsHide, m.CreateAt,
	)

	if err != nil {
		// 發生錯誤，回滾事務
		tx.Rollback()
		log.Fatal(err)
	}

	menuId, err = res.LastInsertId()

	if err != nil {
		panic(err.Error())
	}

	for _, mi := range m.MenuItems {
		res, err = tx.Exec(
			"INSERT INTO menu_item (store_id, menu_category_id, title, `description`, quantity, price)"+
				"VALUE (?, ?, ?, ?, ?, ?)", m.StoreId, mi.MenuCategoryId, mi.Title, mi.Description, mi.Quantity, mi.Price,
		)

		if err != nil {
			// 發生錯誤，回滾事務
			tx.Rollback()
			log.Fatal(err)
		}

		menuItemId, err = res.LastInsertId()

		if err != nil {
			panic(err.Error())
		}

		_, err = tx.Exec(
			"INSERT INTO menu_item_mapping (menu_id, menu_item_id)"+
				"VALUE (?, ?)", menuId, menuItemId,
		)

		if err != nil {
			// 發生錯誤，回滾事務
			tx.Rollback()
			log.Fatal(err)
		}
	}

	// 提交事務
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}

	return nil
}
