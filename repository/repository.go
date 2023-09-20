package repository

import (
	"ordering-system-backend/domain"

	"gorm.io/gorm"
)

type OTP interface {
	Create(token string, email string) error
	Verify(token string, password string) error
}

type Users interface {
	Create(userId string, u domain.UserRequest) error
	GetByEmail(email string) (domain.User, error)
}

type Stores interface {
	Create(s domain.Store) error
	Update(s domain.Store) error
	Delete(storeId string) error
	GetAll() ([]domain.Store, error)
	GetById(storeId string) (domain.Store, error)
}

type Menus interface {
	Create(m domain.Menu) error
	Update(m domain.Menu) error
	Delete(storeId string, menuId int) error
	GetAll(storeId string) ([]domain.Menu, error)
	GetById(storeId string, menuId int) (domain.Menu, error)
}

type Repositories struct {
	Users  Users
	OTP    OTP
	Menus  Menus
	Stores Stores
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		OTP:    NewOTPRepo(db),
		Users:  NewUsersRepo(db),
		Menus:  NewMenusRepo(db),
		Stores: NewStoresRepo(db),
	}
}
