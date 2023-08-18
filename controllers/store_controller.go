package controllers

import (
	"net/http"
	"ordering-system-backend/models"
	s "ordering-system-backend/services"

	"github.com/gin-gonic/gin"
)

// CRUD
func GetStores(c *gin.Context) {
	stores, err := s.GetStores()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, stores)
}

func GetStoreById(c *gin.Context) {
	id := c.Param("store_id")
	stores, err := s.GetStoreById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, stores)
}

func CreateStore(c *gin.Context) {
	var newStore models.Store

	if err := c.BindJSON(&newStore); err != nil {
		return
	}
	err := s.CreateStore(newStore)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, nil)
}

func UpdateStore(c *gin.Context) {}
func DeleteStore(c *gin.Context) {}
