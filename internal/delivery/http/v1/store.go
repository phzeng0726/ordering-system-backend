package v1

import "github.com/gin-gonic/gin"

func (h *Handler) initUserStoreRoutes(api *gin.RouterGroup) {
	stores := api.Group("/stores")
	{
		stores.POST("", h.services.Stores.Create)
		stores.PATCH("/:store_id", h.services.Stores.Update)
		stores.DELETE("/:store_id", h.services.Stores.Delete)
		stores.GET("", h.services.Stores.GetAllByUserId)
		stores.GET("/:store_id", h.services.Stores.GetByStoreId)
	}
}

// 不帶有userId，目前用不到
func (h *Handler) initStoreRoutes(api *gin.RouterGroup) {
	stores := api.Group("/stores")
	{
		stores.GET("", h.services.Stores.GetAll)
	}
}
