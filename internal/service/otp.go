package service

import (
	"context"
	"net/http"
	"ordering-system-backend/internal/domain"
	"ordering-system-backend/internal/repository"

	"github.com/gin-gonic/gin"
)

type OTPService struct {
	repo repository.OTP
}

func NewOTPService(repo repository.OTP) *OTPService {
	return &OTPService{repo: repo}
}

func (s *OTPService) CreateTesting(ctx context.Context, input CreateOTPInput) error {
	if err := s.repo.CreateTesting(ctx, input.Token, input.Password); err != nil {
		return err
	}
	return nil
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
