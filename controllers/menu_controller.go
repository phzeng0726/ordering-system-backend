package controllers

import (
	"net/http"
	"ordering-system-backend/models"
	s "ordering-system-backend/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetMenus(c *gin.Context) {
	storeId := c.Param("store_id")
	menus, err := s.GetMenus(storeId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, menus)
}

func GetMenuById(c *gin.Context) {
	storeId := c.Param("store_id")
	menuIdStr := c.Param("menu_id")
	menuId, err := strconv.Atoi(menuIdStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid menu_id"})
		return
	}

	menus, err := s.GetMenuById(storeId, menuId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, menus)
}

func UpdateMenus(c *gin.Context) {
	var newMenu models.Menu
	if err := c.BindJSON(&newMenu); err != nil {
		return
	}

	err := s.UpdateMenus(newMenu)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, nil)
}

func CreateMenus(c *gin.Context) {
	var newMenu models.Menu

	if err := c.BindJSON(&newMenu); err != nil {
		return
	}

	err := s.CreateMenus(newMenu)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, nil)
}
