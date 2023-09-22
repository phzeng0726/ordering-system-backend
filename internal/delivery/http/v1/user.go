package v1

import "github.com/gin-gonic/gin"

func (h *Handler) initUserRoutes(api *gin.RouterGroup) {
	// 不帶有userId
	userAuth := api.Group("/users")
	{
		userAuth.POST("", h.services.Users.Create)          // 創建User
		userAuth.GET("/login", h.services.Users.GetByEmail) // 透過Email確認user有沒有存在
	}

	// 帶有userId
	user := api.Group("/users/:user_id")
	{
		user.PATCH("", h.services.Users.Update)
		user.GET("", h.services.Users.GetById)
		h.initUserStoreRoutes(user)
	}
}
