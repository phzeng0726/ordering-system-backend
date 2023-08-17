package routes

import (
	c "ordering-system-backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(router *gin.Engine) {
	// Store
	router.GET("/stores", c.GetStores)
	router.POST("/stores", c.CreateStore)
	router.PATCH("/stores", c.UpdateStore)
	router.GET("/stores/:id", c.GetStoreById)
	router.DELETE("/stores/:id", c.DeleteStore)

	// Menu
	router.GET("/stores/:id/menus", c.GetMenus)

}
