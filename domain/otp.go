package domain

import "time"

// OTPCreate: Token & Email
// OTPVerify: Token & Password
type OTPRequest struct {
	Token    string `json:"token" binding:"required"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type OTP struct {
	Token     string    `gorm:"column:token;not null;primaryKey;" json:"token"`
	Password  string    `gorm:"column:password;" json:"password"`
	CreatedAt time.Time `gorm:"column:created_at;" json:"createdAt"`
	Email     string    `gorm:"column:email;" json:"email"`
}

func (OTP) TableName() string {
	return "otp"
}
