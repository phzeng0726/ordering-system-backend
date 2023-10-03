package v1

import (
	"net/http"
	"ordering-system-backend/internal/domain"
	"ordering-system-backend/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
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
		h.initUserStoresRoutes(user)
		h.initUserMenusRoutes(user)
	}
}

type createUserInput struct {
	Email      string `json:"email" binding:"required"`
	Password   string `json:"password" binding:"required"`
	UserType   *int   `json:"userType" binding:"required"`
	FirstName  string `json:"firstName" binding:"required"`
	LastName   string `json:"lastName" binding:"required"`
	LanguageId int    `json:"languageId" binding:"required"`
}

type resetPasswordInput struct {
	UserId   string `json:"userId" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) createUser(c *gin.Context) {
	var inp createUserInput
	if err := c.BindJSON(&inp); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := h.services.Users.Create(c.Request.Context(), service.CreateUserInput{
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

	c.IndentedJSON(http.StatusOK, true)
}

func (h *Handler) deleteUser(c *gin.Context) {
	userId := c.Param("user_id")
	if err := h.services.Users.Delete(c.Request.Context(), userId); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, true)
}

func (h *Handler) getUserByEmail(c *gin.Context) {
	userTypeStr := c.Query("userType")
	method := c.Query("method") // email or uid
	userId := ""
	userType, err := strconv.Atoi(userTypeStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "userType parameter is missing or invalid syntax"})
		return
	}

	if method == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "method parameter is missing"})
		return
	}

	if method == "email" {
		email := c.Query("email")
		if email == "" {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "email parameter is missing"})
			return
		}

		userId, err = h.services.Users.GetByEmail(c.Request.Context(), email, userType)

		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
	} else if method == "uid" {
		uid := c.Query("uid")
		if uid == "" {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "uid parameter is missing"})
			return
		}

		userId, err = h.services.Users.GetByUid(c.Request.Context(), uid, userType)

		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
	}

	c.IndentedJSON(http.StatusOK, userId)
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

	c.IndentedJSON(http.StatusOK, true)
}
