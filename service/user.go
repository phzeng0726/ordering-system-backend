package service

import (
	"net/http"
	"ordering-system-backend/repository"

	"github.com/gin-gonic/gin"
)

type UsersService struct {
	repo repository.Users
}

func NewUsersService(repo repository.Users) *UsersService {
	return &UsersService{repo: repo}
}

func (s *UsersService) Create(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, false)
}

// email
func (s *UsersService) GetByEmail(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, false)
}
