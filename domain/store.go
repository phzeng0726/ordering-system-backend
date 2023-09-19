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
	Phone       string `gorm:"column:phone;" json:"phone"`
	Address     string `gorm:"column:address;" json:"address"`
	LanguageId  int    `gorm:"column:language_id;" json:"languageId"`
}

// Email       string `gorm:"column:email;" json:"email"`
