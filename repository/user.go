package repository

import (
	"ordering-system-backend/domain"

	"gorm.io/gorm"
)

type UsersRepo struct {
	db *gorm.DB
}

func NewUsersRepo(db *gorm.DB) *UsersRepo {
	return &UsersRepo{
		db: db,
	}
}

func (r *UsersRepo) GetByEmail(email string) (domain.User, error) {
	var menu domain.User
	return menu, nil
}
