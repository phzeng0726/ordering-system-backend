package domain

import "time"

// gorm: db的column name
// json: 最後吐出去的json key name
// autoIncrement;
// not null;
// primaryKey;
type FCMToken struct {
	Id          int       `gorm:"column:id;not null;primaryKey;autoIncrement;" json:"id"`
	UserId      string    `gorm:"column:user_id;" json:"userId"`
	DeviceToken string    `gorm:"column:token;" json:"token"`
	CreatedAt   time.Time `gorm:"column:created_at;" json:"createdAt"`
}

func (FCMToken) TableName() string {
	return "fcm_tokens"
}
