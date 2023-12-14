package v1

import (
	"net/http"
	"ordering-system-backend/internal/service"
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
	}
}

type createCategoryInput struct {
	Title      string `json:"title" binding:"required"`
	Identifier string `json:"identifier"`
}

type updateCategoryInput struct {
	Title      string `json:"title" binding:"required"`
	Identifier string `json:"identifier"`
}

// @Tags Categories
// @Description Create category
// @Produce json
// @Accept json
// @Param data body createCategoryInput true "JSON data"
// @Param user_id path string true "User id"
// @Success 200 {boolean} result
// @Router /users/{user_id}/categories [post]
func (h *Handler) createCategory(c *gin.Context) {
	var inp createCategoryInput
	userId := c.Param("user_id")

	if err := c.BindJSON(&inp); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := h.services.Categories.Create(c.Request.Context(), userId, service.CreateCategoryInput{
		Title:      inp.Title,
		Identifier: inp.Identifier,
	}); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, true)
}

// @Tags Categories
// @Description Update category
// @Produce json
// @Accept json
// @Param data body updateCategoryInput true "JSON data"
// @Param user_id path string true "User id"
// @Param category_id path int true "Category id"
// @Success 200 {boolean} result
// @Router /users/{user_id}/categories/{category_id} [patch]
func (h *Handler) updateCategory(c *gin.Context) {
	var inp updateCategoryInput
	userId := c.Param("user_id")
	categoryIdStr := c.Param("category_id")
	categoryId, err := strconv.Atoi(categoryIdStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "category parameter is missing or invalid syntax"})
		return
	}

	if err := c.BindJSON(&inp); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := h.services.Categories.Update(c.Request.Context(), userId, service.UpdateCategoryInput{
		CategoryId: categoryId,
		Title:      inp.Title,
		Identifier: inp.Identifier,
	}); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, true)
}

// @Tags Categories
// @Description Delete category
// @Produce json
// @Param user_id path string true "User id"
// @Param category_id path int true "Category id"
// @Success 200 {boolean} result
// @Router /users/{user_id}/categories/{category_id} [delete]
func (h *Handler) deleteCategory(c *gin.Context) {
	userId := c.Param("user_id")
	categoryIdStr := c.Param("category_id")
	categoryId, err := strconv.Atoi(categoryIdStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "category parameter is missing or invalid syntax"})
		return
	}

	if err := h.services.Categories.Delete(c.Request.Context(), userId, categoryId); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, true)
}

// @Tags Categories
// @Description Get all categories by user id
// @Produce json
// @Param user_id path string true "User id"
// @Param language query int true "Language"
// @Success 200 {array} domain.Category
// @Router /users/{user_id}/categories [get]
func (h *Handler) getAllCategoriesByUserId(c *gin.Context) {
	userId := c.Param("user_id")
	languageIdStr := c.Query("language")
	languageId, err := strconv.Atoi(languageIdStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "language parameter is missing or invalid syntax"})
		return
	}

	categories, err := h.services.Categories.GetAllByUserId(c.Request.Context(), userId, languageId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, categories)
}
