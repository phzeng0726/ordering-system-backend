package v1

import (
	"net/http"
	"ordering-system-backend/internal/domain"

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

func (h *Handler) createMenu(c *gin.Context) {
	var inp domain.Menu
	userId := c.Param("user_id")

	if err := c.BindJSON(&inp); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	inp.UserId = userId

	menuId, err := h.services.Menus.Create(c.Request.Context(), inp)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	inp.Id = menuId
	c.IndentedJSON(http.StatusOK, inp)
}

func (h *Handler) updateMenu(c *gin.Context) {
	// storeId := c.Param("store_id")
	// menuIdStr := c.Param("menu_id")
	// menuId, err := strconv.Atoi(menuIdStr)
	// if err != nil {
	// 	c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid menu_id"})
	// 	return
	// }

	// var newMenu domain.Menu
	// if err := c.BindJSON(&newMenu); err != nil {
	// 	return
	// }

	// newMenu.StoreId = storeId
	// newMenu.Id = menuId

	// // 確認是否store有該menu，有的話才會做下一步，避免別人把她的menu刪掉
	// menus, err := s.repo.GetById(storeId, menuId)
	// if err != nil || menus.Id == 0 {
	// 	if menus.Id == 0 {
	// 		err = errors.New("menu not found for this store")
	// 	}
	// 	c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	// 	return
	// }

	// err = s.repo.Update(newMenu)
	// if err != nil {
	// 	c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	// 	return
	// }

	// c.IndentedJSON(http.StatusOK, newMenu)
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
	// storeId := c.Param("store_id")
	// menus, err := s.repo.GetAll(storeId)
	// if err != nil {
	// 	c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	// 	return
	// }

	// c.IndentedJSON(http.StatusOK, menus)
}

func (h *Handler) getById(c *gin.Context) {
	// storeId := c.Param("store_id")
	// menuIdStr := c.Param("menu_id")
	// menuId, err := strconv.Atoi(menuIdStr)
	// if err != nil {
	// 	c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid menu_id"})
	// 	return
	// }

	// menus, err := s.repo.GetById(storeId, menuId)
	// if err != nil {
	// 	c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	// 	return
	// }

	// c.IndentedJSON(http.StatusOK, menus)
}
