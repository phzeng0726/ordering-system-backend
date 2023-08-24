package domain

import (
	"time"
)

// autoIncrement 會在insert之後自行塞回model，所以可以直接用Id取lastInsertId
type Menu struct {
	Id          int        `gorm:"column:id;not null;primaryKey;autoIncrement;" json:"id"`
	StoreId     string     `gorm:"column:store_id;" json:"storeId"`
	Title       string     `gorm:"column:title;" json:"title"`
	Description string     `gorm:"column:description;" json:"description"`
	IsHide      bool       `gorm:"column:is_hide;" json:"isHide"`
	CreateAt    time.Time  `gorm:"column:create_at;" json:"createAt"`
	MenuItems   []MenuItem `gorm:"-" json:"menuItems"`
}

type MenuItem struct {
	Id           int          `gorm:"column:id;not null;primaryKey;autoIncrement;" json:"id"`
	Title        string       `gorm:"column:title;" json:"title"`
	Description  string       `gorm:"column:description;" json:"description"`
	Quantity     int          `gorm:"column:quantity;" json:"quantity"`
	Price        int          `gorm:"column:price;" json:"price"`
	MenuCategory MenuCategory `gorm:"-" json:"menuCategory"`
}

type MenuCategory struct {
	Id    int    `gorm:"column:id;not null;primaryKey;" json:"id"`
	Title string `gorm:"column:title;" json:"title"`
}

type MenuItemMapping struct {
	Id         int `gorm:"column:id;not null;primaryKey;autoIncrement;" json:"id"`
	MenuId     int `gorm:"column:menu_id;" json:"menuId"`
	MenuItemId int `gorm:"column:menu_item_id;" json:"menuItemId"`
}
