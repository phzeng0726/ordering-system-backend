package v1

import "github.com/gin-gonic/gin"

func (h *Handler) initUserRoutes(api *gin.RouterGroup) {
	user := api.Group("/user")
	{
		user.GET("/login", h.services.Users.GetByEmail) // 透過Email確認user有沒有存在
	}
}
