package http

import (
	"net/http"
	v1 "ordering-system-backend/internal/delivery/http/v1"
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

func (h *Handler) Init() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.Use(
		gin.Recovery(),
		corsMiddleware,
	)

	// router.Use(Middleware(conn))
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	h.initAPI(router)

	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	handlerV1 := v1.NewHandler(h.services)
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}
