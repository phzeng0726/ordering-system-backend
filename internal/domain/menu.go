package domain

import (
	"time"
)

// autoIncrement 會在insert之後自行塞回model，所以可以直接用Id取lastInsertId
type Menu struct {
	Id          string     `gorm:"column:id;not null;primaryKey;" json:"id"`
	UserId      string     `gorm:"column:user_id;" json:"userId"`
	Title       string     `gorm:"column:title;" json:"title"`
	Description string     `gorm:"column:description;" json:"description"`
	CreatedAt   time.Time  `gorm:"column:created_at;" json:"createdAt"`
	MenuItems   []MenuItem `gorm:"-" json:"menuItems,omitempty"`
}

type MenuItemMapping struct {
	Id         int      `gorm:"column:id;not null;primaryKey;autoIncrement;" json:"id"`
	MenuId     string   `gorm:"column:menu_id;" json:"menuId"`
	MenuItemId int      `gorm:"column:menu_item_id;" json:"menuItemId"`
	Menu       Menu     `gorm:"foreignKey:MenuId;references:id;" json:"menu"`
	MenuItem   MenuItem `gorm:"foreignKey:MenuItemId;references:id;" json:"menuItem"`
}

type MenuItem struct {
	Id          int      `gorm:"column:id;not null;primaryKey;autoIncrement;" json:"id"`
	Title       string   `gorm:"column:title;" json:"title"`
	Description string   `gorm:"column:description;" json:"description"`
	Quantity    int      `gorm:"column:quantity;" json:"quantity"`
	Price       int      `gorm:"column:price;" json:"price"`
	CategoryId  int      `gorm:"column:category_id;" json:"-"` // 外鍵欄位名稱
	Category    Category `gorm:"foreignKey:CategoryId;references:id;" json:"category"`
	ImageId     int      `gorm:"column:image_id;" json:"-"` // 外鍵欄位名稱
	Image       Image    `gorm:"foreignKey:ImageId;references:id;" json:"-"`
	ImageBytes  []byte   `gorm:"-" json:"imageBytes"`
}

func (MenuItemMapping) TableName() string {
	return "menu_item_mapping"
}
