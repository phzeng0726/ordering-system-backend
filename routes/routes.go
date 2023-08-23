package routes

import (
	"ordering-system-backend/service"

	"github.com/gin-gonic/gin"
)

type RoutesSetup struct {
	Router  *gin.Engine
	Service *service.Services
}

func NewRoutesSetup(router *gin.Engine, service *service.Services) *RoutesSetup {
	return &RoutesSetup{
		Router:  router,
		Service: service,
	}
}

func (rs *RoutesSetup) setupStoreRoutes() {
	rs.Router.POST("/stores", rs.Service.Stores.Create)
	rs.Router.PATCH("/stores", rs.Service.Stores.Update)
	rs.Router.DELETE("/stores/:store_id", rs.Service.Stores.Delete)
	rs.Router.GET("/stores", rs.Service.Stores.GetAll)
	rs.Router.GET("/stores/:store_id", rs.Service.Stores.GetById)
}

func (rs *RoutesSetup) setupMenuRoutes() {
	rs.Router.POST("/stores/:store_id/menus", rs.Service.Menus.Create)
	rs.Router.PATCH("/stores/:store_id/menus", rs.Service.Menus.Update)
	rs.Router.GET("/stores/:store_id/menus", rs.Service.Menus.GetAll)
	rs.Router.GET("/stores/:store_id/menus/:menu_id", rs.Service.Menus.GetById)
}

func SetUpRoutes(router *gin.Engine, s *service.Services) {
	routesSetup := NewRoutesSetup(router, s)
	routesSetup.setupMenuRoutes()
	routesSetup.setupStoreRoutes()
}
