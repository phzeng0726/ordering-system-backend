package v1

import "github.com/gin-gonic/gin"

func (h *Handler) initOTPRoutes(api *gin.RouterGroup) {
	otp := api.Group("/otp")
	{
		otp.POST("/create", h.services.OTP.Create) // createOTP 建立OTP
		otp.POST("/verify", h.services.OTP.Verify) // verifyOTP 驗證OTP
	}
}
