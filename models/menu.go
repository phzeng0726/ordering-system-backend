package models

import "time"

type Menu struct {
	Id          int       `json:"id"`
	StoreId     string    `json:"store_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	IsHide      bool      `json:"is_hide"`
	CreateAt    time.Time `json:"create_at"`
}

// type MenuItem struct {
// 	Id             int    `json:"id"`
// 	StoreId        int    `json:"store_id"`
// 	MenuCategoryId int    `json:"menu_category_id"`
// 	Name           string `json:"name"`
// 	Description    string `json:"description"`
// 	Price          int    `json:"price"`
// }
