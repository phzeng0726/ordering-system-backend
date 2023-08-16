package routes

import (
	c "ordering-system-backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(router *gin.Engine) {
	router.GET("/stores", c.GetStores)
	router.GET("/store/:id", c.GetStoreById)

}
