package service

import (
	"context"
	"ordering-system-backend/internal/repository"

	"github.com/gin-gonic/gin"
)

type CreateOTPInput struct {
	Token    string
	Email    string
	Password string
}

type OTP interface {
	CreateTesting(ctx context.Context, input CreateOTPInput) error
	Create(c *gin.Context)
	Verify(c *gin.Context)
}

type Users interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	GetByEmail(c *gin.Context)
	GetById(c *gin.Context)
	ResetPassword(c *gin.Context)
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
