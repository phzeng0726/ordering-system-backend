package v1

import (
	"ordering-system-backend/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		// For common
		h.initOTPRoutes(v1)
		h.initUserRoutes(v1)

		// For 商家
		h.initUserCategoryRoutes(v1)
		h.initUserMenusRoutes(v1)
		h.initUserStoresRoutes(v1)
		h.initUserSeatsRoutes(v1)
		h.initStoreMenusRoutes(v1)
		h.initFCMTokensRoutes(v1) // FCM

		// For 客戶
		h.initOrderTicketRoutes(v1)
	}
}
