package v1

import (
	"net/http"
	"ordering-system-backend/internal/domain"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initUserStoresRoutes(api *gin.RouterGroup) {
	stores := api.Group("/users/:user_id/stores")
	{
		stores.POST("", h.createStore)
		stores.PATCH("/:store_id", h.updateStore)
		stores.DELETE("/:store_id", h.deleteStore)
		stores.GET("/:store_id", h.getStoreByStoreId)
		stores.GET("", h.getAllStoresByUserId)
	}
}

func (h *Handler) createStore(c *gin.Context) {
	var inp domain.Store
	userId := c.Param("user_id")

	if err := c.BindJSON(&inp); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	inp.UserId = userId

	storeId, err := h.services.Stores.Create(c.Request.Context(), inp)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	inp.Id = storeId
	c.IndentedJSON(http.StatusOK, inp)
}

func (h *Handler) updateStore(c *gin.Context) {
	var inp domain.Store
	userId := c.Param("user_id")
	storeId := c.Param("store_id")

	if err := c.BindJSON(&inp); err != nil {
		return
	}

	inp.Id = storeId
	inp.UserId = userId

	if err := h.services.Stores.Update(c.Request.Context(), inp); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, inp)
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
