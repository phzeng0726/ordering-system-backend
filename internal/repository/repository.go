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
	Create(ctx context.Context, category domain.Category, categoryLanguage domain.CategoryLanguage, categoryUserMapping domain.CategoryUserMapping) error
	Update(ctx context.Context, userId string, category domain.Category, categoryLanguage domain.CategoryLanguage) error
	Delete(ctx context.Context, userId string, categoryId int) error
	GetAllByUserId(ctx context.Context, userId string, languageId int) ([]domain.CategoryUserMapping, error)
}

type Seats interface {
	Create(ctx context.Context, seat domain.Seat) error
	Update(ctx context.Context, seat domain.Seat) error
	Delete(ctx context.Context, storeId string, seatId int) error
	GetAllByStoreId(ctx context.Context, storeId string) ([]domain.Seat, error)
	GetById(ctx context.Context, storeId string, seatId int) (domain.Seat, error)
	GetSeatBySeatId(ctx context.Context, seatId int) (domain.Seat, error)
}

type Menus interface {
	Create(ctx context.Context, menu domain.Menu) error
	Update(ctx context.Context, menu domain.Menu) error
	Delete(ctx context.Context, userId string, menuId string) error
	TempGetAllByUserId(ctx context.Context, userId string, languageId int) ([]domain.Menu, error)
	GetAllByUserId(ctx context.Context, userId string, languageId int) ([]domain.MenuItemMapping, error)
	GetById(ctx context.Context, userId string, menuId string, languageId int) ([]domain.MenuItemMapping, error)
}

type StoreMenus interface {
	CreateMenuReference(ctx context.Context, userId string, storeMenuMapping domain.StoreMenuMapping) error
	UpdateMenuReference(ctx context.Context, userId string, storeMenuMapping domain.StoreMenuMapping) error
	DeleteMenuReference(ctx context.Context, userId string, storeId string) error
	GetMenuByStoreId(ctx context.Context, userId string, storeId string, languageId int, userType int) (domain.Menu, error)
}

type OrderTickets interface {
	Create(ctx context.Context, ticket domain.OrderTicket) error
	Update(ctx context.Context, storeId string, ticket domain.OrderTicket) error
	Delete(ctx context.Context, storeId string, ticketId int) error
	GetAllByStoreId(ctx context.Context, storeId string) ([]domain.OrderTicket, error)
	GetAllByUserId(ctx context.Context, userId string) ([]domain.OrderTicket, error)
}

type FCMTokens interface {
	Create(ctx context.Context, token domain.FCMToken) error
	Delete(ctx context.Context, token domain.FCMToken) error
	GetByUserId(ctx context.Context, userId string) (string, error)
	GetAllBySeatId(ctx context.Context, seatId int) ([]string, error)     // For client user，用來通知store user
	GetAllByTicketId(ctx context.Context, ticketId int) ([]string, error) // For store user，用來通知client user

}

type Images interface {
	Create(tx *gorm.DB, imageBytes []byte) error
	GetById(tx *gorm.DB, imageId int) (domain.Image, error)
}

type Repositories struct {
	Users        Users
	OTP          OTP
	Stores       Stores
	Seats        Seats
	Categories   Categories
	Menus        Menus
	StoreMenus   StoreMenus
	OrderTickets OrderTickets
	FCMTokens    FCMTokens
	Images       Images
}

type RepoTools struct{}

func (*RepoTools) CheckStoreSeatExist(tx *gorm.DB, storeId string, seatId int) error {
	if err := tx.Where("id = ?", seatId).First(&domain.Seat{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("seat id not found")
		}
		return err
	}

	return nil
}

func (*RepoTools) GetStoreInfo(tx *gorm.DB, storeId string, store *domain.Store) error {
	// 沒有傳入指針時，代表外部不需要使用到
	if store == nil {
		store = &domain.Store{}
	}

	if err := tx.Where("id = ?", storeId).First(&store).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// userId避免print在log上
			return fmt.Errorf("no store found with id %s", storeId)
		}
		return err
	}

	if err := tx.Where("store_id = ?", storeId).Find(&store.StoreOpeningHours).Error; err != nil {
		return err
	}

	fmt.Println(store)

	return nil
}

func (*RepoTools) CheckUserStoreExist(tx *gorm.DB, userId string, storeId string, store *domain.Store) error {
	// 沒有傳入指針時，代表外部不需要使用到
	if store == nil {
		store = &domain.Store{}
	}

	if err := tx.Where("user_id = ? AND id = ?", userId, storeId).First(&store).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// userId避免print在log上
			return fmt.Errorf("no store found with id %s for this user id", storeId)
		}
		return err
	}

	return nil
}

func (*RepoTools) CheckUserMenuExist(tx *gorm.DB, userId string, menuId string, menu *domain.Menu) error {
	// 沒有傳入指針時，代表外部不需要使用到
	if menu == nil {
		menu = &domain.Menu{}
	}

	if err := tx.Where("user_id = ? AND id = ?", userId, menuId).First(&menu).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// userId避免print在log上
			return fmt.Errorf("no menu found with id %s for this user id", menuId)
		}
		return err
	}

	return nil
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		OTP:          NewOTPRepo(db),
		Users:        NewUsersRepo(db),
		Stores:       NewStoresRepo(db),
		Seats:        NewSeatsRepo(db),
		Categories:   NewCategoriesRepo(db),
		Menus:        NewMenusRepo(db),
		StoreMenus:   NewStoreMenusRepo(db),
		OrderTickets: NewOrderTicketsRepo(db),
		FCMTokens:    NewFCMTokensRepo(db),
		Images:       NewImagesRepo(db),
	}
}
