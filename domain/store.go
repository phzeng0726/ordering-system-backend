package domain

// gorm: db的column name
// json: 最後吐出去的json key name
// autoIncrement;
// not null;
// primaryKey;
type Store struct {
	Id          string `gorm:"column:id;not null;primaryKey;" json:"id"`
	Name        string `gorm:"column:name;" json:"name"`
	Description string `gorm:"column:description;" json:"description"`
	Email       string `gorm:"column:email;" json:"email"`
	Phone       string `gorm:"column:phone;" json:"phone"`
	IsOpen      bool   `gorm:"column:is_open;" json:"isOpen"`
}
