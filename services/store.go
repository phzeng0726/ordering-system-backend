package services

import (
	"errors"
	"fmt"
	"log"
	"ordering-system-backend/db"
	"ordering-system-backend/models"
)

func GetStores() ([]models.Store, error) {
	sql := "SELECT * FROM store"
	rows, err := db.DB.Query(sql)
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

	return stores, err
}

func GetStoreById(storeId string) (models.Store, error) {
	sql := "SELECT *" +
		" FROM store" +
		" WHERE id = ?" +
		" LIMIT 1"
	rows, err := db.DB.Query(sql, storeId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var store models.Store

	found := false // 用於標記是否找到符合條件的資料行
	for rows.Next() {
		found = true // 找到資料行
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
	}

	if !found {
		err = errors.New("store not found")
	}

	return store, err
}
