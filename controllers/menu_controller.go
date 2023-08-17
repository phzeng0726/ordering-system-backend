package controllers

import (
	"net/http"
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
