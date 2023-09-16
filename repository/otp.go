package repository

import (
	"gorm.io/gorm"
)

type OTPRepo struct {
	db *gorm.DB
}

func NewOTPRepo(db *gorm.DB) *OTPRepo {
	return &OTPRepo{
		db: db,
	}
}

func (r *OTPRepo) Create(token string, email string) error {
	return nil
}

func (r *OTPRepo) Verify(token string, password string) error {
	return nil
}
