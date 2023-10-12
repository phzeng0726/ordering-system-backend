package domain

type Seat struct {
	Id          int    `gorm:"column:id;not null;primaryKey;" json:"id"`
	StoreId     string `gorm:"column:store_id;" json:"-"`
	Description string `gorm:"column:description;" json:"description"`
}

func (Seat) TableName() string {
	return "store_seats"
}
