package service

import (
	"net/http"
	"ordering-system-backend/domain"
	"ordering-system-backend/repository"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UsersService struct {
	repo repository.Users
}

func NewUsersService(repo repository.Users) *UsersService {
	return &UsersService{repo: repo}
}

func (s *UsersService) Create(c *gin.Context) {
	var newUserReq domain.UserRequest
	if err := c.BindJSON(&newUserReq); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	uuid := uuid.New()

	err := s.repo.Create(uuid.String(), newUserReq)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, newUserReq)
}

// email
func (s *UsersService) GetByEmail(c *gin.Context) {
	email := c.Query("email")
	userTypeStr := c.Query("userType")
	userType, err := strconv.Atoi(userTypeStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	userId, err := s.repo.GetByEmail(email, userType)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if userId == "" {
		c.IndentedJSON(http.StatusOK, false)
		return
	}

	c.IndentedJSON(http.StatusOK, userId)
}
