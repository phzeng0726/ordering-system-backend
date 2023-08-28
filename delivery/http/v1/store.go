package v1

import "github.com/gin-gonic/gin"

func (h *Handler) initStoreRoutes(api *gin.RouterGroup) {
	stores := api.Group("/stores")
	{
		stores.POST("", h.services.Stores.Create)
		stores.PATCH("/:store_id", h.services.Stores.Update)
		stores.DELETE("/:store_id", h.services.Stores.Delete)
		stores.GET("", h.services.Stores.GetAll)
		stores.GET("/:store_id", h.services.Stores.GetById)
	}
}
