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
		fcmTokens.GET("", h.getToken) // App沒有實際使用，只是用來方便開發測試用

	}
}

type createTokenInput struct {
	UserId      string `json:"userId" binding:"required"`
	DeviceToken string `json:"token" binding:"required"`
}

type deleteTokenInput struct {
	UserId      string `json:"userId" binding:"required"`
	DeviceToken string `json:"token" binding:"required"`
}

// @Tags FCM Token
// @Description Insert FCM Token with UserId
// @Accept json
// @Param data body createTokenInput true "JSON data"
// @Produce json
// @Success 200 {string} string FCM_Token
// @Router /fcm-tokens [post]
func (h *Handler) createToken(c *gin.Context) {
	var inp createTokenInput

	if err := c.BindJSON(&inp); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := h.services.FCMTokens.Create(c.Request.Context(), service.CreateTokenInput{
		UserId:      inp.UserId,
		DeviceToken: inp.DeviceToken,
	}); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, true)
}

// @Tags FCM Token
// @Description Get FCM Token by UserId
// @Produce json
// @Param userId query string true "userId"
// @Success 200 {string} string FCM_Token
// @Router /fcm-tokens [get]
func (h *Handler) getToken(c *gin.Context) {
	userId := c.Query("userId")

	token, err := h.services.FCMTokens.GetByUserId(c.Request.Context(), userId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, token)
}

// @Tags FCM Token
// @Description Delete FCM Token with UserId and token
// @Accept json
// @Param data body deleteTokenInput true "JSON data"
// @Produce json
// @Success 200 {string} string deletedResult
// @Router /fcm-tokens [delete]
func (h *Handler) deleteToken(c *gin.Context) {
	var inp deleteTokenInput

	if err := c.BindJSON(&inp); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := h.services.FCMTokens.Delete(c.Request.Context(), service.DeleteTokenInput{
		UserId:      inp.UserId,
		DeviceToken: inp.DeviceToken,
	}); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, true)
}
