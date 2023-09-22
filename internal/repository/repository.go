package repository

import (
	"ordering-system-backend/internal/domain"

	"gorm.io/gorm"
)

type OTP interface {
	Create(token string, email string) error
	Verify(token string, password string) error
}

type Users interface {
	Create(userId string, ur domain.UserRequest) error
	Update(u domain.User) error // 只能更新 User，不能更新 UserAccount
	GetByEmail(email string, userType int) (string, error)
	GetById(userId string) (domain.User, error)
	Delete(userId string) error
	ResetPassword(ur domain.UserRequest) error
}

type Stores interface {
	Create(s domain.Store) error
	Update(userId string, s domain.Store) error
	Delete(userId string, storeId string) error
	GetAllByUserId(userId string) ([]domain.Store, error)
	GetByStoreId(userId string, storeId string) (domain.Store, error)

	// 不含UserId
	GetAll() ([]domain.Store, error)
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
