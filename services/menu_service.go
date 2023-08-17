package services

import (
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
