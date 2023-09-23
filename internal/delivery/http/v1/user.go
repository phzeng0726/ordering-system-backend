package v1

import (
	"net/http"
	"ordering-system-backend/internal/domain"
	"ordering-system-backend/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) initUserRoutes(api *gin.RouterGroup) {
	// 不帶有userId
	userAuth := api.Group("/users")
	{
		userAuth.POST("", h.createUser)                   // 創建User
		userAuth.GET("/login", h.getUserByEmail)          // 透過Email確認user有沒有存在
		userAuth.POST("/reset-password", h.resetPassword) // 重設密碼
	}

	// 帶有userId
	user := api.Group("/users/:user_id")
	{
		user.PATCH("", h.updateUser)
		user.DELETE("", h.deleteUser)
		user.GET("", h.getUserById)
		h.initUserStoreRoutes(user)
	}
}

type createUserInput struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	UserType   int    `json:"userType"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	LanguageId int    `json:"languageId"`
}

type resetPasswordInput struct {
	UserId   string `json:"userId"`
	Password string `json:"password"`
}

func (h *Handler) createUser(c *gin.Context) {
	var inp createUserInput
	if err := c.BindJSON(&inp); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	userId := uuid.New().String()
	if err := h.services.Users.Create(c.Request.Context(), service.CreateUserInput{
		UserId:     userId,
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

func (h *Handler) resetPassword(c *gin.Context) {
	var inp resetPasswordInput
	if err := c.BindJSON(&inp); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := h.services.Users.ResetPassword(c.Request.Context(), service.ResetPasswordInput{
		UserId:   inp.UserId,
		Password: inp.Password,
	}); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, inp)
}

func (h *Handler) getUserByEmail(c *gin.Context) {
	email := c.Query("email")
	userTypeStr := c.Query("userType")
	userType, err := strconv.Atoi(userTypeStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	userId, err := h.services.Users.GetByEmail(c.Request.Context(), email, userType)

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

func (h *Handler) updateUser(c *gin.Context) {
	var inp domain.User
	userId := c.Param("user_id")

	if err := c.BindJSON(&inp); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := h.services.Users.Update(c.Request.Context(), userId, service.UpdateUserInput{
		FirstName:  inp.FirstName,
		LastName:   inp.LastName,
		LanguageId: inp.LanguageId,
	}); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, inp)
}

func (h *Handler) deleteUser(c *gin.Context) {
	userId := c.Param("user_id")
	if err := h.services.Users.Delete(c.Request.Context(), userId); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, true)
}

func (h *Handler) getUserById(c *gin.Context) {
	userId := c.Param("user_id")
	user, err := h.services.Users.GetById(c.Request.Context(), userId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}
