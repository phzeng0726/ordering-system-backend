package repository

import (
	"context"
	"errors"
	"fmt"
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
	GetByUid(ctx context.Context, uid string, userType int) (string, error)
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

type Categories interface {
	GetAllByUserId(ctx context.Context, userId string, languageId int) ([]domain.CategoryUserMapping, error)
}

type Seats interface {
	GetById(ctx context.Context, storeId string, seatId int) (domain.Seat, error)
}

type Menus interface {
	Create(ctx context.Context, menu domain.Menu) error
	Update(ctx context.Context, menu domain.Menu) error
	Delete(ctx context.Context, userId string, menuId string) error
	GetAllByUserId(ctx context.Context, userId string, languageId int) ([]domain.MenuItemMapping, error)
	GetById(ctx context.Context, userId string, menuId string, languageId int) ([]domain.MenuItemMapping, error)
}

type Repositories struct {
	Users      Users
	OTP        OTP
	Stores     Stores
	Seats      Seats
	Categories Categories
	Menus      Menus
}

type RepoTools struct{}

func (*RepoTools) CheckUserExist(tx *gorm.DB, userId string) error {
	if err := tx.Where("id = ?", userId).First(&domain.User{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user id not found")
		}
		return err
	}

	return nil
}

func (*RepoTools) CheckStoreExist(tx *gorm.DB, userId string, storeId string) (domain.Store, error) {
	var store domain.Store

	if err := tx.Where("user_id = ? AND id = ?", userId, storeId).First(&store).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// userId避免print在log上
			return store, fmt.Errorf("no store found with id %s for this user id", storeId)
		}
		return store, err
	}

	return store, nil
}

func (*RepoTools) GetUserAccount(tx *gorm.DB, userId string) (domain.UserAccount, error) {
	var userAccount domain.UserAccount

	if err := tx.Where("id = ?", userId).First(&userAccount).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return userAccount, errors.New("user id not found")
		}
		return userAccount, err
	}

	return userAccount, nil
}

func NewRepositories(db *gorm.DB, rt *RepoTools) *Repositories {
	return &Repositories{
		OTP:        NewOTPRepo(db),
		Users:      NewUsersRepo(db, rt),
		Stores:     NewStoresRepo(db, rt),
		Seats:      NewSeatsRepo(db, rt),
		Categories: NewCategoriesRepo(db, rt),
		Menus:      NewMenusRepo(db, rt),
	}
}
