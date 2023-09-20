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
	CreatedAt   time.Time  `gorm:"column:created_at;" json:"createdAt"`
	MenuItems   []MenuItem `gorm:"-" json:"menuItems,omitempty"`
}

type MenuItemMapping struct {
	Id         int      `gorm:"column:id;not null;primaryKey;autoIncrement;" json:"id"`
	MenuId     int      `gorm:"column:menu_id;" json:"menuId"`
	MenuItemId int      `gorm:"column:menu_item_id;" json:"menuItemId"`
	Menu       Menu     `gorm:"foreignKey:MenuId;references:id;" json:"menu"`
	MenuItem   MenuItem `gorm:"foreignKey:MenuItemId;references:id;" json:"menuItem"`
}

type MenuItem struct {
	Id             int          `gorm:"column:id;not null;primaryKey;autoIncrement;" json:"id"`
	Title          string       `gorm:"column:title;" json:"title"`
	Description    string       `gorm:"column:description;" json:"description"`
	Quantity       int          `gorm:"column:quantity;" json:"quantity"`
	Price          int          `gorm:"column:price;" json:"price"`
	MenuCategoryId int          `gorm:"column:menu_category_id;" json:"-"` // 外鍵欄位名稱
	MenuCategory   MenuCategory `gorm:"foreignKey:MenuCategoryId;references:id;" json:"menuCategory"`
}

type MenuCategory struct {
	Id    int    `gorm:"column:id;not null;primaryKey;" json:"id"`
	Title string `gorm:"column:title;" json:"title"`
}

func (MenuItemMapping) TableName() string {
	return "menu_item_mapping"
}
