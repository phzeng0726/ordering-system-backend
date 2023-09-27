package repository

import (
	"context"
	"errors"
	"ordering-system-backend/internal/domain"

	"gorm.io/gorm"
)

type OTP interface {
	Create(ctx context.Context, token string, email string) error
	Verify(ctx context.Context, token string, password string) error
}

type Users interface {
	Create(ctx context.Context, userAccount domain.UserAccount, user domain.User, password string) error
	ResetPassword(ctx context.Context, userId string, newPassword string) error
	Update(ctx context.Context, user domain.User) error // 只能更新 User，不能更新 UserAccount
	GetByEmail(ctx context.Context, email string, userType int) (string, error)
	GetById(ctx context.Context, userId string) (domain.User, error)
	Delete(ctx context.Context, userId string) error
}

type Stores interface {
	Create(ctx context.Context, store domain.Store) error
	Update(ctx context.Context, store domain.Store) error
	Delete(ctx context.Context, userId string, storeId string) error
	GetAllByUserId(ctx context.Context, userId string) ([]domain.Store, error)
	GetByStoreId(ctx context.Context, userId string, storeId string) (domain.Store, error)

	// 不含UserId
	GetAll(ctx context.Context) ([]domain.Store, error)
}

type Menus interface {
	Create(ctx context.Context, menu domain.Menu) error
	Update(ctx context.Context, menu domain.Menu) error
	Delete(ctx context.Context, userId string, menuId string) error
	GetAllByUserId(ctx context.Context, userId string) ([]domain.Menu, error)
	GetById(ctx context.Context, userId string, menuId string) (domain.Menu, error)
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

func CheckUserExist(tx *gorm.DB, userId string) error {
	if err := tx.Where("id = ?", userId).First(&domain.User{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user id not found")
		}
		return err
	}

	return nil
}
