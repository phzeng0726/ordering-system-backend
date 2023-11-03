package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initStoreMenusRoutes(api *gin.RouterGroup) {
	storesAndMenus := api.Group("/users/:user_id/stores/:store_id/menus")
	{
		storesAndMenus.GET("", h.getMenuByStoreId)
		storesAndMenus.POST("/:menu_id", h.createStoreMenuReference)
		storesAndMenus.PATCH("/:menu_id", h.updateStoreMenuReference)
		storesAndMenus.DELETE("", h.deleteStoreMenuReference)
	}
	storesAndMenusWithoutUser := api.Group("/stores/:store_id/menus")
	{
		storesAndMenusWithoutUser.GET("", h.getMenuByStoreId)
	}
}

func (h *Handler) createStoreMenuReference(c *gin.Context) {
	userId := c.Param("user_id")
	storeId := c.Param("store_id")
	menuId := c.Param("menu_id")

	if err := h.services.StoreMenus.CreateMenuReference(c.Request.Context(), userId, storeId, menuId); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, true)
}

func (h *Handler) updateStoreMenuReference(c *gin.Context) {
	userId := c.Param("user_id")
	storeId := c.Param("store_id")
	menuId := c.Param("menu_id")

	if err := h.services.StoreMenus.UpdateMenuReference(c.Request.Context(), userId, storeId, menuId); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, true)
}

func (h *Handler) deleteStoreMenuReference(c *gin.Context) {
	userId := c.Param("user_id")
	storeId := c.Param("store_id")

	if err := h.services.StoreMenus.DeleteMenuReference(c.Request.Context(), userId, storeId); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, true)
}

func (h *Handler) getMenuByStoreId(c *gin.Context) {
	userId := c.Param("user_id") // 商家無userId
	storeId := c.Param("store_id")
	languageIdStr := c.Query("language")
	languageId, err := strconv.Atoi(languageIdStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "language parameter is missing or invalid syntax"})
		return
	}

	menu, err := h.services.StoreMenus.GetMenuByStoreId(c.Request.Context(), userId, storeId, languageId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, menu)
}
