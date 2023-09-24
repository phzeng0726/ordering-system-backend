package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func corsMiddleware(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-Type", "application/json")

	if c.Request.Method != "OPTIONS" {
		fmt.Println("next middleware")
		c.Next()
	} else {
		fmt.Println("abort middleware")
		c.AbortWithStatus(http.StatusOK)
	}
}
