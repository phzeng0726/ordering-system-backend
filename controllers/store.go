package controllers

import (
	"fmt"
	"net/http"
	s "ordering-system-backend/services"

	"github.com/gin-gonic/gin"
)

func GetStores(c *gin.Context) {
	stores, err := s.GetStores()
	if err != nil {
		fmt.Println(err)
	}
	c.IndentedJSON(http.StatusOK, stores)
}

func GetStoreById(c *gin.Context) {
	id := c.Param("id")
	stores, err := s.GetStoreById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Store not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, stores)
}
