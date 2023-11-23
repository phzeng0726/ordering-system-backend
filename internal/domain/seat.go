package domain

type Seat struct {
	Id      int    `gorm:"column:id;not null;primaryKey;" json:"id"`
	StoreId string `gorm:"column:store_id;" json:"-"`
	Title   string `gorm:"column:title;" json:"title"`
	Store   Store  `gorm:"foreignKey:StoreId;references:id" json:"store"`
}

func (Seat) TableName() string {
	return "store_seats"
}
