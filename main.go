package main

import (
	"ordering-system-backend/db"
	"ordering-system-backend/services"
)

func main() {
	db.Connect()

	services.GetStores()

}

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
