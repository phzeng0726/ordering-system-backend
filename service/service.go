package service

import (
	"ordering-system-backend/repository"

	"github.com/gin-gonic/gin"
)

type Stores interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	GetAll(c *gin.Context)
	GetById(c *gin.Context)
}
type Menus interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	GetAll(c *gin.Context)
	GetById(c *gin.Context)
}

type Services struct {
	Menus  Menus
	Stores Stores
}

type Deps struct {
	Repos *repository.Repositories
}

func NewServices(deps Deps) *Services {
	menusService := NewMenusService(deps.Repos.Menus)
	storesService := NewStoresService(deps.Repos.Stores)

	return &Services{
		Menus:  menusService,
		Stores: storesService,
	}
}
