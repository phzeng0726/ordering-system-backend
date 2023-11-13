package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initTokenRoutes(api *gin.RouterGroup) {
	tokens := api.Group("/tokens")
	{
		tokens.POST("", h.createToken)
		tokens.GET("", h.getToken)
	}
}

func (h *Handler) createToken(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, true)
}

func (h *Handler) getToken(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, true)
}
