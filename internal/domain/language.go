package domain

// gorm: db的column name
// json: 最後吐出去的json key name
// autoIncrement;
// not null;
// primaryKey;
type Language struct {
	Id        string `gorm:"column:id;not null;primaryKey;" json:"id"`
	Code      string `gorm:"column:code;" json:"code"`
	IsDefault *bool  `gorm:"column:is_default;" json:"isDefault"`
	IsEnable  *bool  `gorm:"column:is_enable;" json:"isEnable"`
}

func (Language) TableName() string {
	return "language"
}
