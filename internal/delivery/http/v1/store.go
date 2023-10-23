package v1

import (
	"net/http"
	"ordering-system-backend/internal/domain"
	"strconv"

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

	storesAndMenus := api.Group("/users/:user_id/stores/:store_id/menus")
	{
		storesAndMenus.GET("", h.getMenuByStoreId)
		storesAndMenus.POST("/:menu_id", h.createStoreMenuReference)
		storesAndMenus.PATCH("/:menu_id", h.updateStoreMenuReference)
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

// Store Menu Reference

func (h *Handler) createStoreMenuReference(c *gin.Context) {
	userId := c.Param("user_id")
	storeId := c.Param("store_id")
	menuId := c.Param("menu_id")

	if err := h.services.Stores.CreateMenuReference(c.Request.Context(), userId, storeId, menuId); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, nil)
}

func (h *Handler) updateStoreMenuReference(c *gin.Context) {
	userId := c.Param("user_id")
	storeId := c.Param("store_id")
	menuId := c.Param("menu_id")

	if err := h.services.Stores.UpdateMenuReference(c.Request.Context(), userId, storeId, menuId); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, nil)
}

func (h *Handler) getMenuByStoreId(c *gin.Context) {
	userId := c.Param("user_id")
	storeId := c.Param("store_id")
	languageIdStr := c.Query("language")
	languageId, err := strconv.Atoi(languageIdStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "language parameter is missing or invalid syntax"})
		return
	}

	menu, err := h.services.Stores.GetMenuByStoreId(c.Request.Context(), userId, storeId, languageId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, menu)
}
