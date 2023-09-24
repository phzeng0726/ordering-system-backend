package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initUserStoreRoutes(api *gin.RouterGroup) {
	stores := api.Group("/stores")
	{
		stores.POST("", h.services.Stores.Create)
		stores.PATCH("/:store_id", h.services.Stores.Update)
		stores.DELETE("/:store_id", h.deleteStore)
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

func (h *Handler) deleteStore(c *gin.Context) {
	userId := c.Param("user_id")
	storeId := c.Param("store_id")

	if err := h.services.Stores.Delete(c.Request.Context(), userId, storeId); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, true)
}
