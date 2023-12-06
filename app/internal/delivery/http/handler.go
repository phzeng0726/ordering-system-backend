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

	router.GET("/ping", h.ping)

	// Swagger 文件
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

// PingExample
// @Tags Get Started
// @Summary Get user by ID
// @Description Get a user by its ID
// @ID get-user-by-id
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {string} User
// @Router /users/{id} [get]
func (h *Handler) ping(g *gin.Context) {
	g.JSON(http.StatusOK, "pong")
}
