package controllers

import (
	"net/http"
	"ordering-system-backend/services"

	"github.com/gin-gonic/gin"
)

func GetStores(c *gin.Context) {
	stores := services.GetStores()
	c.IndentedJSON(http.StatusOK, stores)
}
