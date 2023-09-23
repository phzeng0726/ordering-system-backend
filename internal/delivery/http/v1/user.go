package v1

import (
	"net/http"
	"ordering-system-backend/internal/domain"
	"ordering-system-backend/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)

func (h *Handler) initUserRoutes(api *gin.RouterGroup) {
	// 不帶有userId
	userAuth := api.Group("/users")
	{
		userAuth.POST("", h.createUser)                                  // 創建User
		userAuth.GET("/login", h.services.Users.GetByEmail)              // 透過Email確認user有沒有存在
		userAuth.POST("/reset-password", h.services.Users.ResetPassword) // 重設密碼
	}

	// 帶有userId
	user := api.Group("/users/:user_id")
	{
		user.PATCH("", h.services.Users.Update)
		user.DELETE("", h.services.Users.Delete)
		user.GET("", h.services.Users.GetById)
		h.initUserStoreRoutes(user)
	}
}

type createUserInput struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	UserType   int    `json:"user_type"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	LanguageId int    `json:"language_id"`
}

func (h *Handler) createUser(c *gin.Context) {
	var inp createUserInput
	if err := c.BindJSON(&inp); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	uuid := uuid.New()
	if err := h.services.Users.Create(c.Request.Context(), service.CreateUserInput{
		UserId:     uuid.String(),
		Email:      inp.Email,
		Password:   inp.Password,
		UserType:   inp.UserType,
		FirstName:  inp.FirstName,
		LastName:   inp.LastName,
		LanguageId: inp.LanguageId,
	}); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, inp)
}

func (h *Handler) Update(c *gin.Context) {
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

func (h *Handler) Delete(c *gin.Context) {
	userId := c.Param("user_id")

	err := s.repo.Delete(userId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, true)
}

func (h *Handler) GetByEmail(c *gin.Context) {
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

func (h *Handler) GetById(c *gin.Context) {
	id := c.Param("user_id")
	user, err := s.repo.GetById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, user)
}

func (h *Handler) ResetPassword(c *gin.Context) {
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
