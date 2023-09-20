package domain

import "time"

// TODO 修改成真正的User
// autoIncrement 會在insert之後自行塞回model，所以可以直接用Id取lastInsertId
type User struct {
	Id         int    `gorm:"column:id;not null;primaryKey;autoIncrement;" json:"id"`
	FirstName  string `gorm:"column:first_name;" json:"firstName"`
	LastName   string `gorm:"column:last_name;" json:"lastName"`
	Phone      string `gorm:"column:phone;" json:"phone"`
	LanguageId int    `gorm:"column:language_id;" json:"languageId"`
}

type UserAccount struct {
	Id        string    `gorm:"column:id;not null;primaryKey;" json:"id"`
	UidCode   string    `gorm:"column:uid_code;" json:"userCode"`
	Email     string    `gorm:"column:email;" json:"email"`
	UserType  int       `gorm:"column:user_type;" json:"userType"`
	CreatedAt time.Time `gorm:"column:created_at;" json:"createAt"`
}
