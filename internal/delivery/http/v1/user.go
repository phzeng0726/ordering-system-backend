package v1

import "github.com/gin-gonic/gin"

func (h *Handler) initUserRoutes(api *gin.RouterGroup) {
	user := api.Group("/user")
	{
		user.POST("", h.services.Users.Create)           // 創建User
		user.PATCH("/:user_id", h.services.Users.Update) // 創建User
		user.GET("/login", h.services.Users.GetByEmail)  // 透過Email確認user有沒有存在
		user.GET("/:user_id", h.services.Users.GetById)
	}
}
