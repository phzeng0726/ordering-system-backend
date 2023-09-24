package service

import (
	"context"
	"ordering-system-backend/internal/domain"
	"ordering-system-backend/internal/repository"

	"github.com/gin-gonic/gin"
)

// 不同層之間可能需要做資料轉換，所以delivery和service分開
type CreateOTPInput struct {
	Token string
	Email string
}

type VerifyOTPInput struct {
	Token    string
	Password string
}

type OTP interface {
	Create(ctx context.Context, input CreateOTPInput) error
	Verify(ctx context.Context, input VerifyOTPInput) error
}

type CreateUserInput struct {
	Email      string
	Password   string
	UserType   int
	FirstName  string
	LastName   string
	LanguageId int
}

type UpdateUserInput struct {
	FirstName  string
	LastName   string
	LanguageId int
}

type ResetPasswordInput struct {
	UserId   string
	Password string
}

type Users interface {
	Create(ctx context.Context, input CreateUserInput) error
	Update(ctx context.Context, userId string, input UpdateUserInput) error
	Delete(ctx context.Context, userId string) error
	GetByEmail(ctx context.Context, email string, userType int) (string, error)
	GetById(ctx context.Context, userId string) (domain.User, error)
	ResetPassword(ctx context.Context, input ResetPasswordInput) error
}

type Stores interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	GetAllByUserId(c *gin.Context)
	GetByStoreId(c *gin.Context)

	// 不含UserId
	GetAll(c *gin.Context)
}
type Menus interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	GetAll(c *gin.Context)
	GetById(c *gin.Context)
}

type Services struct {
	Users  Users
	OTP    OTP
	Menus  Menus
	Stores Stores
}

type Deps struct {
	Repos *repository.Repositories
}

func NewServices(deps Deps) *Services {
	usersService := NewUsersService(deps.Repos.Users)
	OTPService := NewOTPService(deps.Repos.OTP)
	menusService := NewMenusService(deps.Repos.Menus)
	storesService := NewStoresService(deps.Repos.Stores)

	return &Services{
		Users:  usersService,
		OTP:    OTPService,
		Menus:  menusService,
		Stores: storesService,
	}
}
