package service

import (
	"net/http"
	"ordering-system-backend/repository"

	"github.com/gin-gonic/gin"
)

type OTPService struct {
	repo repository.OTP
}

func NewOTPService(repo repository.OTP) *OTPService {
	return &OTPService{repo: repo}
}

// token, email
func (s *OTPService) Create(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, true)
}

// token, password
func (s *OTPService) Verify(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, true)
}
