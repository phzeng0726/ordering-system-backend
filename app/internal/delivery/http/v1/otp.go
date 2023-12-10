package v1

import (
	"net/http"
	"ordering-system-backend/internal/service"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initOTPRoutes(api *gin.RouterGroup) {
	otp := api.Group("/otp")
	{
		otp.POST("/create", h.create) // 建立OTP
		otp.POST("/verify", h.verify) // 驗證OTP
	}
}

type createOTPInput struct {
	Token string `json:"token" binding:"required"`
	Email string `json:"email" binding:"required"`
}

type verifyOTPInput struct {
	Token    string `json:"token" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Tags OTP
// @Description Create OTP
// @Accept json
// @Param data body createOTPInput true "JSON data"
// @Produce json
// @Success 200 {boolean} result
// @Router /otp/create [post]
func (h *Handler) create(c *gin.Context) {
	var inp createOTPInput
	if err := c.BindJSON(&inp); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := h.services.OTP.Create(c.Request.Context(), service.CreateOTPInput{
		Token: inp.Token,
		Email: inp.Email,
	}); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, true)
}

// @Tags OTP
// @Description Verify OTP
// @Accept json
// @Param data body verifyOTPInput true "JSON data"
// @Produce json
// @Success 200 {boolean} result
// @Router /otp/verify [post]
func (h *Handler) verify(c *gin.Context) {
	var inp verifyOTPInput
	if err := c.BindJSON(&inp); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := h.services.OTP.Verify(c.Request.Context(), service.VerifyOTPInput{
		Token:    inp.Token,
		Password: inp.Password,
	}); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, true)
}
