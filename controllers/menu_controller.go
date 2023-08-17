package controllers

import (
	"net/http"
	"ordering-system-backend/models"
	s "ordering-system-backend/services"

	"github.com/gin-gonic/gin"
)

func GetMenus(c *gin.Context) {
	id := c.Param("id")
	menus, err := s.GetMenus(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, menus)
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
