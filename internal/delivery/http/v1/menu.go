package v1

import (
	"fmt"
	"net/http"
	"ordering-system-backend/internal/service"

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
	// storeId := c.Param("store_id")
	// menuIdStr := c.Param("menu_id")
	// menuId, err := strconv.Atoi(menuIdStr)
	// if err != nil {
	// 	c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid menu_id"})
	// 	return
	// }

	// // 確認是否store有該menu，有的話才會做下一步，避免別人把她的menu刪掉
	// menus, err := s.repo.GetById(storeId, menuId)
	// if err != nil || menus.Id == 0 {
	// 	if menus.Id == 0 {
	// 		err = errors.New("menu not found for this store")
	// 	}
	// 	c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	// 	return
	// }

	// err = s.repo.Delete(storeId, menuId)
	// if err != nil {
	// 	c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	// 	return
	// }

	// c.IndentedJSON(http.StatusOK, nil)
}

func (h *Handler) getAllByUserId(c *gin.Context) {
	userId := c.Param("user_id")

	menus, err := h.services.Menus.GetAllByUserId(c.Request.Context(), userId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, menus)
}

func (h *Handler) getById(c *gin.Context) {
	userId := c.Param("user_id")
	menuId := c.Param("menu_id")

	fmt.Println(menuId)
	menu, err := h.services.Menus.GetById(c.Request.Context(), userId, menuId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, menu)
}
