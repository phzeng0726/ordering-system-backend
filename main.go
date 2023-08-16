package main

import (
	"database/sql"
	"fmt"
	"log"
	"ordering-system-backend/models"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	appConfig := AppConfig{
		UserName: "root",
		Password: "0000",
		DBName:   "ordering_db",
	}

	// 設定資料庫連線字串
	dsn := fmt.Sprintf("%s:%s@/%s", appConfig.UserName, appConfig.Password, appConfig.DBName)

	// 連線資料庫
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 測試連線
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MySQL database!")

	// 執行 SELECT 查詢
	rows, err := db.Query("SELECT * FROM store")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var stores []models.Store
	for rows.Next() {
		var store models.Store
		err := rows.Scan(
			&store.Id,
			&store.Name,
			&store.Description,
			&store.Email,
			&store.Phone,
			&store.IsOpen,
		)
		if err != nil {
			log.Fatal(err)
		}
		stores = append(stores, store)
	}
	fmt.Println(stores)

	// router := gin.Default()

	// menuItem := models.MenuItem{
	// 	Id:             1,
	// 	StoreId:        1,
	// 	MenuCategoryId: 1,
	// 	Name:           "",
	// 	Description:    "",
	// 	Price:          0,
	// }

	// jsonData, err := json.MarshalIndent(menuItem, "", "  ")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(jsonData)
}
