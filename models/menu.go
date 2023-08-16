package models

type MenuItem struct {
	Id             int    `json:"id"`
	StoreId        int    `json:"store_id"`
	MenuCategoryId int    `json:"menu_category_id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Price          int    `json:"price"`
}
