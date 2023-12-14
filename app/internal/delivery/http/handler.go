package http

import (
	"net/http"
	v1 "ordering-system-backend/internal/delivery/http/v1"
	"ordering-system-backend/internal/service"

	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Init() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.Use(
		gin.Recovery(),
		corsMiddleware,
	)

	h.initAPI(router)

	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	handlerV1 := v1.NewHandler(h.services)
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}

	// 唯一一個在 Swagger 文件上不可使用，因為BasePath非/api/v1
	router.GET("/ping", h.ping)

	// Swagger 文件
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

// @Tags Get Started
// @Description 測試API是否成功運作
// @Produce json
// @Success 200 {string} Pong
// @Router /ping [get]
func (h *Handler) ping(g *gin.Context) {
	g.JSON(http.StatusOK, "pong")
}
