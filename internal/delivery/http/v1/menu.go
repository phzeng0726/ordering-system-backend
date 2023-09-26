package v1

import "github.com/gin-gonic/gin"

func (h *Handler) initUserMenusRoutes(api *gin.RouterGroup) {
	menus := api.Group("/menus")
	{
		menus.POST("", h.services.Menus.Create)
		menus.PATCH("/:menu_id", h.services.Menus.Update)
		menus.DELETE("/:menu_id", h.services.Menus.Delete)
		menus.GET("", h.services.Menus.GetAll)
		menus.GET("/:menu_id", h.services.Menus.GetById)
	}
}
