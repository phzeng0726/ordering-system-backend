package v1

import (
	"net/http"
	"ordering-system-backend/internal/service"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initFCMTokensRoutes(api *gin.RouterGroup) {
	fcmTokens := api.Group("/fcm-tokens")
	{
		fcmTokens.POST("", h.createToken)
		fcmTokens.DELETE("", h.deleteToken)
		fcmTokens.GET("", h.getToken)

	}
}

type createTokenInput struct {
	UserId string `json:"userId" binding:"required"`
	Token  string `json:"token" binding:"required"`
}

type deleteTokenInput struct {
	UserId string `json:"userId" binding:"required"`
	Token  string `json:"token" binding:"required"`
}

func (h *Handler) createToken(c *gin.Context) {
	var inp createTokenInput

	if err := c.BindJSON(&inp); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := h.services.FCMTokens.Create(c.Request.Context(), service.CreateTokenInput{
		UserId: inp.UserId,
		Token:  inp.Token,
	}); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, true)
}

func (h *Handler) getToken(c *gin.Context) {
	userId := c.Query("userId")

	token, err := h.services.FCMTokens.GetByUserId(c.Request.Context(), userId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, token)
}

func (h *Handler) deleteToken(c *gin.Context) {
	var inp deleteTokenInput

	if err := c.BindJSON(&inp); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := h.services.FCMTokens.Delete(c.Request.Context(), service.DeleteTokenInput{
		UserId: inp.UserId,
		Token:  inp.Token,
	}); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, true)
}
