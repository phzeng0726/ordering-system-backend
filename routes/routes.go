package routes

import (
	"ordering-system-backend/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	router  *gin.Engine
	service *service.Services
}

func NewHandler(router *gin.Engine, service *service.Services) *Handler {
	return &Handler{
		router:  router,
		service: service,
	}
}

func (h *Handler) initStoreRoutes(api *gin.RouterGroup) {
	stores := api.Group("/stores")
	{
		stores.POST("", h.service.Stores.Create)
		stores.PATCH("/:store_id", h.service.Stores.Update)
		stores.DELETE("/:store_id", h.service.Stores.Delete)
		stores.GET("", h.service.Stores.GetAll)
		stores.GET("/:store_id", h.service.Stores.GetById)
	}
}

func (h *Handler) initMenuRoutes(api *gin.RouterGroup) {
	menus := api.Group("/stores/:store_id/menus")
	{
		menus.POST("", h.service.Menus.Create)
		menus.PATCH("/:menu_id", h.service.Menus.Update)
		menus.DELETE("/:menu_id", h.service.Menus.Delete)
		menus.GET("", h.service.Menus.GetAll)
		menus.GET("/:menu_id", h.service.Menus.GetById)
	}
}

func SetUpRoutes(h *Handler) {
	api := h.router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			h.initMenuRoutes(v1)
			h.initStoreRoutes(v1)
		}
	}
}
