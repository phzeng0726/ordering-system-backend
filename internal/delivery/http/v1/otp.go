package v1

import (
	"net/http"
	"ordering-system-backend/internal/service"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initOTPRoutes(api *gin.RouterGroup) {
	otp := api.Group("/otp")
	{
		otp.POST("/create-testing", h.createTesting) // createOTP 建立OTP
		otp.POST("/create", h.services.OTP.Create)   // createOTP 建立OTP
		otp.POST("/verify", h.services.OTP.Verify)   // verifyOTP 驗證OTP
	}
}

type createOTPInput struct {
	Token    string `json:"token" binding:"required"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// token, email
func (h *Handler) createTesting(c *gin.Context) {
	var inp createOTPInput
	if err := c.BindJSON(&inp); err != nil {
		return
	}

	if err := h.services.OTP.CreateTesting(c.Request.Context(), service.CreateOTPInput{
		Token:    inp.Token,
		Email:    inp.Email,
		Password: inp.Password,
	}); err != nil {
		return
	}

	c.Status(http.StatusCreated)
}
