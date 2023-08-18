package models

import (
	"time"
)

type Menu struct {
	Id          int        `json:"id"`
	StoreId     string     `json:"store_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	IsHide      bool       `json:"is_hide"`
	MenuItems   []MenuItem `json:"menu_items"`
	CreateAt    time.Time  `json:"create_at"`
}

type MenuItem struct {
	Id           int          `json:"id"`
	MenuCategory MenuCategory `json:"menu_category"`
	Title        string       `json:"title"`
	Description  string       `json:"description"`
	Quantity     int          `json:"quantity"`
	Price        int          `json:"price"`
}

type MenuCategory struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}
