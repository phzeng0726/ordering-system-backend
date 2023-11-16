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
		menus.GET("/:category_id", h.getCategoryById)
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
