package service

import (
	"net/http"
	"ordering-system-backend/domain"
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
	var req domain.OTPRequest
	if err := c.BindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := s.repo.Create(req.Token, req.Email)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, true)
}

// token, password
func (s *OTPService) Verify(c *gin.Context) {
	var req domain.OTPRequest
	if err := c.BindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := s.repo.Verify(req.Token, req.Password)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, true)
}
