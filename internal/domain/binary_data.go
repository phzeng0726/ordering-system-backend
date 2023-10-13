package domain

type BinaryData struct {
	Id      int    `gorm:"column:id;not null;primaryKey;autoIncrement;" json:"id"`
	BinData []byte `gorm:"column:bin_data;" json:"binData"`
}

func (BinaryData) TableName() string {
	return "binary_data"
}
