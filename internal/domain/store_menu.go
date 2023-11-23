package domain

type StoreMenuMapping struct {
	Id      int    `gorm:"column:id;not null;primaryKey;" json:"-"`
	StoreId string `gorm:"column:store_id;" json:"storeId"`
	MenuId  string `gorm:"column:menu_id;" json:"menuId"`
}

func (StoreMenuMapping) TableName() string {
	return "store_menu_mapping"
}
