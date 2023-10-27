package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initUserCategoryRoutes(api *gin.RouterGroup) {
	menus := api.Group("/users/:user_id/categories")
	{
		menus.POST("", h.createCategory)
		menus.PATCH("/:category_id", h.updateCategory)
		menus.DELETE("/:category_id", h.deleteCategory)
		menus.GET("", h.getAllCategoriesByUserId)
		menus.GET("/:category_id", h.getCategoryById)
	}
}

func (h *Handler) createCategory(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, true)
}
func (h *Handler) updateCategory(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, true)
}
func (h *Handler) deleteCategory(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, true)
}
func (h *Handler) getAllCategoriesByUserId(c *gin.Context) {
	userId := c.Param("user_id")
	languageIdStr := c.Query("language")
	languageId, err := strconv.Atoi(languageIdStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "language parameter is missing or invalid syntax"})
		return
	}

	menus, err := h.services.Categories.GetAllByUserId(c.Request.Context(), userId, languageId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, menus)
}
func (h *Handler) getCategoryById(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, true)
}
