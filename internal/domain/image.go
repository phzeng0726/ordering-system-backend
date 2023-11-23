package domain

type Image struct {
	Id        int    `gorm:"column:id;not null;primaryKey;autoIncrement;" json:"id"`
	BytesData []byte `gorm:"column:bytes_data;" json:"imageBytes"`
}
