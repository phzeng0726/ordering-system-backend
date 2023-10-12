package v1

import (
	"fmt"
	"net/http"
	"ordering-system-backend/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initUserMenusRoutes(api *gin.RouterGroup) {
	menus := api.Group("/menus")
	{
		menus.POST("", h.createMenu)
		menus.PATCH("/:menu_id", h.updateMenu)
		menus.DELETE("/:menu_id", h.deleteMenu)
		menus.GET("", h.getAllByUserId)
		menus.GET("/:menu_id", h.getById)
	}
}

type createMenuInput struct {
	StoreId     string                `json:"storeId" binding:"required"`
	Title       string                `json:"title" binding:"required"`
	Description string                `json:"description"`
	MenuItems   []createMenuItemInput `json:"menuItems" binding:"required"`
}

type createMenuItemInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Quantity    *int   `json:"quantity" binding:"required"`
	Price       *int   `json:"price" binding:"required"`
	CategoryId  *int   `json:"categoryId" binding:"required"`
}

func (h *Handler) createMenu(c *gin.Context) {
	var inp createMenuInput
	userId := c.Param("user_id")

	if err := c.BindJSON(&inp); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var menuItems []service.CreateMenuItemInput
	for _, mi := range inp.MenuItems {
		menuItems = append(menuItems, service.CreateMenuItemInput{
			Title:       mi.Title,
			Description: mi.Description,
			Quantity:    *mi.Quantity,
			Price:       *mi.Price,
			CategoryId:  *mi.CategoryId,
		})
	}

	menuId, err := h.services.Menus.Create(c.Request.Context(), service.CreateMenuInput{
		UserId:      userId,
		StoreId:     inp.StoreId,
		Title:       inp.Title,
		Description: inp.Description,
		MenuItems:   menuItems,
	})

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, menuId)
}

func (h *Handler) updateMenu(c *gin.Context) {
	var inp createMenuInput
	userId := c.Param("user_id")
	menuId := c.Param("menu_id")

	fmt.Println(menuId)
	if err := c.BindJSON(&inp); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var menuItems []service.CreateMenuItemInput
	for _, mi := range inp.MenuItems {
		menuItems = append(menuItems, service.CreateMenuItemInput{
			Title:       mi.Title,
			Description: mi.Description,
			Quantity:    *mi.Quantity,
			Price:       *mi.Price,
			CategoryId:  *mi.CategoryId,
		})
	}

	err := h.services.Menus.Update(c.Request.Context(), service.UpdateMenuInput{
		UserId:      userId,
		MenuId:      menuId,
		StoreId:     inp.StoreId,
		Title:       inp.Title,
		Description: inp.Description,
		MenuItems:   menuItems,
	})

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, menuId)
}

func (h *Handler) deleteMenu(c *gin.Context) {
	userId := c.Param("user_id")
	menuId := c.Param("menu_id")

	if err := h.services.Menus.Delete(c.Request.Context(), userId, menuId); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, menuId)
}

func (h *Handler) getAllByUserId(c *gin.Context) {
	userId := c.Param("user_id")
	languageIdStr := c.Query("languageId")
	languageId, err := strconv.Atoi(languageIdStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "languageId parameter is missing or invalid syntax"})
		return
	}

	menus, err := h.services.Menus.GetAllByUserId(c.Request.Context(), userId, languageId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, menus)
}

func (h *Handler) getById(c *gin.Context) {
	userId := c.Param("user_id")
	menuId := c.Param("menu_id")
	languageIdStr := c.Query("languageId")
	languageId, err := strconv.Atoi(languageIdStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "languageId parameter is missing or invalid syntax"})
		return
	}

	menu, err := h.services.Menus.GetById(c.Request.Context(), userId, menuId, languageId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, menu)
}
