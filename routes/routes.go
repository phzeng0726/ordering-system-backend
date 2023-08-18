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
	router.GET("/stores/:store_id", c.GetStoreById)
	router.DELETE("/stores/:store_id", c.DeleteStore)

	// Menu
	router.GET("/stores/:store_id/menus", c.GetMenus)             // 得到menu列表
	router.GET("/stores/:store_id/menus/:menu_id", c.GetMenuById) // 得到詳細menu資訊
	router.POST("/stores/:store_id/menus", c.CreateMenus)         // 創建menu
}
