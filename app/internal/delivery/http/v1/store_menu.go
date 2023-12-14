package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initStoreMenusRoutes(api *gin.RouterGroup) {
	storesAndMenus := api.Group("/users/:user_id/stores/:store_id/menus")
	{
		storesAndMenus.GET("", h.getMenuByStoreId) // TODO 疑似不需要了，須跟前端確認
		storesAndMenus.POST("/:menu_id", h.createStoreMenuReference)
		storesAndMenus.PATCH("/:menu_id", h.updateStoreMenuReference)
		storesAndMenus.DELETE("/:menu_id", h.deleteStoreMenuReference) // App沒有實際使用，只是用來方便開發測試用
	}
	storesAndMenusWithoutUser := api.Group("/stores/:store_id/menus")
	{
		storesAndMenusWithoutUser.GET("", h.getMenuByStoreId)
	}
}

// @Tags Store Menus
// @Description Create the reference between store and menu
// @Param user_id path string true "User Id"
// @Param store_id path string true "Store Id"
// @Param menu_id path string true "Menu Id"
// @Produce json
// @Success 200 {boolean} result
// @Router /users/{user_id}/stores/{store_id}/menus/{menu_id} [post]
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

// @Tags Store Menus
// @Description Update the reference between store and menu
// @Param user_id path string true "User Id"
// @Param store_id path string true "Store Id"
// @Param menu_id path string true "Menu Id"
// @Produce json
// @Success 200 {boolean} result
// @Router /users/{user_id}/stores/{store_id}/menus/{menu_id} [patch]
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

// @Tags Store Menus
// @Description Delete the reference between store and menu
// @Param user_id path string true "User Id"
// @Param store_id path string true "Store Id"
// @Param menu_id path string true "Menu Id"
// @Produce json
// @Success 200 {boolean} result
// @Router /users/{user_id}/stores/{store_id}/menus/{menu_id} [delete]
func (h *Handler) deleteStoreMenuReference(c *gin.Context) {
	userId := c.Param("user_id")
	storeId := c.Param("store_id")
	menuId := c.Param("menu_id")

	if err := h.services.StoreMenus.DeleteMenuReference(c.Request.Context(), userId, storeId, menuId); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, true)
}

// @Tags Store Menus
// @Description Get menu by store id
// @Param language query int true "區分多語系，en為1, zh為2"
// @Param userType query int true "用來區分用戶類別，StoreEase商家為0, OrderEase客戶為1"
// @Produce json
// @Success 200 {object} domain.Menu
// @Router /stores/{store_id}/menus [get]
func (h *Handler) getMenuByStoreId(c *gin.Context) {
	userId := c.Param("user_id") //  TODO 疑似不需要了，須跟前端確認，商家無userId
	storeId := c.Param("store_id")
	languageIdStr := c.Query("language")
	languageId, err := strconv.Atoi(languageIdStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "language parameter is missing or invalid syntax"})
		return
	}

	userTypeStr := c.Query("userType")
	userType, err := strconv.Atoi(userTypeStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "userType parameter is missing or invalid syntax"})
		return
	}

	menu, err := h.services.StoreMenus.GetStoreMenuByStoreId(c.Request.Context(), userId, storeId, languageId, userType)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, menu)
}
