package service

import (
	"ordering-system-backend/repository"

	"github.com/gin-gonic/gin"
)

type Menus interface {
	GetMenus(c *gin.Context)
}

type Services struct {
	Menus Menus
}

type Deps struct {
	Repos *repository.Repositories
}

func NewServices(deps Deps) *Services {
	menusService := NewMenusService(deps.Repos.Menus)

	return &Services{
		Menus: menusService,
	}
}
