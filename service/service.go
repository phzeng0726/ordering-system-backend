package service

import (
	"ordering-system-backend/repository"

	"github.com/gin-gonic/gin"
)

type Stores interface {
	GetStores(c *gin.Context)
	GetStoreById(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}
type Menus interface {
	GetMenus(c *gin.Context)
	GetMenuById(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
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
