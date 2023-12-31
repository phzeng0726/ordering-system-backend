package domain

import "time"

type User struct {
	Id          string      `gorm:"column:id;not null;primaryKey;" json:"id"`
	FirstName   string      `gorm:"column:first_name;" json:"firstName"`
	LastName    string      `gorm:"column:last_name;" json:"lastName"`
	LanguageId  int         `gorm:"column:language_id;" json:"languageId"`
	UserAccount UserAccount `gorm:"foreignKey:Id;references:id;" json:"-"` // 為了取email
	Email       string      `gorm:"-" json:"email,omitempty"`
}

type UserAccount struct {
	Id        string    `gorm:"column:id;not null;primaryKey;" json:"id"`
	UidCode   string    `gorm:"column:uid_code;" json:"userCode"`
	Email     string    `gorm:"column:email;" json:"email"`
	UserType  *int      `gorm:"column:user_type;" json:"userType"` // 0: store, 1: customer，設定required如果值是0會有錯，所以要用成*
	CreatedAt time.Time `gorm:"column:created_at;" json:"createdAt"`
}

func (User) TableName() string {
	return "user"
}

func (UserAccount) TableName() string {
	return "user_account"
}
