package service

import (
	"context"
	"net/http"
	"ordering-system-backend/internal/domain"
	"ordering-system-backend/internal/repository"
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

func (s *UsersService) Create(ctx context.Context, input CreateUserInput) error {
	uuid := uuid.New()
	userAccount := domain.UserAccount{
		Id:       input.UserId,
		UidCode:  uuid.String(),
		Email:    input.Email,
		UserType: input.UserType,
	}

	user := domain.User{
		Id:         input.UserId,
		FirstName:  input.FirstName,
		LastName:   input.LastName,
		LanguageId: input.LanguageId,
	}

	if err := s.repo.Create(ctx, userAccount, user, input.Password); err != nil {
		return err
	}
	return nil
}

func (s *UsersService) Update(c *gin.Context) {
	var newUser domain.User
	id := c.Param("user_id")

	if err := c.BindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	newUser.Id = id
	err := s.repo.Update(newUser)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, newUser)
}

func (s *UsersService) Delete(c *gin.Context) {
	userId := c.Param("user_id")

	err := s.repo.Delete(userId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, true)
}

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

func (s *UsersService) GetById(c *gin.Context) {
	id := c.Param("user_id")
	user, err := s.repo.GetById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, user)
}

func (s *UsersService) ResetPassword(c *gin.Context) {
	var ur domain.UserRequest

	if err := c.BindJSON(&ur); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := s.repo.ResetPassword(ur)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, true)
}
