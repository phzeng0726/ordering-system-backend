package v1

import (
	"net/http"
	"ordering-system-backend/internal/domain"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initUserStoreRoutes(api *gin.RouterGroup) {
	stores := api.Group("/stores")
	{
		stores.POST("", h.createStore)
		stores.PATCH("/:store_id", h.updateStore)
		stores.DELETE("/:store_id", h.deleteStore)
		stores.GET("", h.getAllStoresByUserId)
		stores.GET("/:store_id", h.getStoreByStoreId)
	}
}

// 不帶有userId，目前用不到
func (h *Handler) initStoreRoutes(api *gin.RouterGroup) {
	stores := api.Group("/stores")
	{
		stores.GET("", h.getAllStores)
	}
}

func (h *Handler) createStore(c *gin.Context) {
	var newStore domain.Store
	userId := c.Param("user_id")

	if err := c.BindJSON(&newStore); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	newStore.UserId = userId

	storeId, err := h.services.Stores.Create(c.Request.Context(), newStore)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	newStore.Id = storeId
	c.IndentedJSON(http.StatusOK, newStore)
}

func (h *Handler) updateStore(c *gin.Context) {
	var newStore domain.Store
	userId := c.Param("user_id")
	storeId := c.Param("store_id")

	if err := c.BindJSON(&newStore); err != nil {
		return
	}

	newStore.Id = storeId
	newStore.UserId = userId

	if err := h.services.Stores.Update(c.Request.Context(), newStore); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, newStore)
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

func (h *Handler) getAllStoresByUserId(c *gin.Context) {
	userId := c.Param("user_id")

	stores, err := h.services.Stores.GetAllByUserId(c.Request.Context(), userId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, stores)
}

func (h *Handler) getStoreByStoreId(c *gin.Context) {
	userId := c.Param("user_id")
	storeId := c.Param("store_id")

	store, err := h.services.Stores.GetByStoreId(c.Request.Context(), userId, storeId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, store)
}

func (h *Handler) getAllStores(c *gin.Context) {
	stores, err := h.services.Stores.GetAll(c.Request.Context())
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, stores)
}
