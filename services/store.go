package services

import (
	"fmt"
	"log"
	"ordering-system-backend/db"
	"ordering-system-backend/models"
)

func GetStores() []models.Store {
	// 執行 SELECT 查詢
	rows, err := db.DB.Query("SELECT * FROM store")
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

	return stores
}
