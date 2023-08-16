package routes

import (
	c "ordering-system-backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(router *gin.Engine) {
	router.GET("/stores", c.GetStores)
	router.POST("/stores", c.CreateStore)
	router.PATCH("/stores", c.UpdateStore)

	router.GET("/store/:id", c.GetStoreById)
	router.DELETE("/store/:id", c.DeleteStore)

}
