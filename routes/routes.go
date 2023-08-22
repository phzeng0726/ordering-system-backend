package routes

import (
	"ordering-system-backend/service"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(router *gin.Engine, s *service.Services) {
	// Store
	router.GET("/stores", s.Stores.GetStores)
	router.GET("/stores/:store_id", s.Stores.GetStoreById)
	router.POST("/stores", s.Stores.Create)
	router.PATCH("/stores", s.Stores.Update)
	router.DELETE("/stores/:store_id", s.Stores.Delete)

	// Menu
	router.GET("/stores/:store_id/menus", s.Menus.GetMenus)             // 得到menu列表
	router.GET("/stores/:store_id/menus/:menu_id", s.Menus.GetMenuById) // 得到詳細menu資訊
	router.POST("/stores/:store_id/menus", s.Menus.Create)              // 創建menu
	router.PATCH("/stores/:store_id/menus", s.Menus.Update)             // 修改menu資訊

}
