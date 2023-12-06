package v1

import (
	"net/http"
	"ordering-system-backend/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initUserRoutes(api *gin.RouterGroup) {
	// 不帶有userId
	user := api.Group("/users")
	{
		// Auth
		user.POST("", h.createUser)                   // 創建User
		user.GET("/login", h.login)                   // 透過Email確認user有沒有存在
		user.POST("/reset-password", h.resetPassword) // 重設密碼

		// Others
		user.PATCH("/:user_id", h.updateUser)
		user.DELETE("/:user_id", h.deleteUser)
		user.GET("/:user_id", h.getUserById)
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

type updateUserInput struct {
	FirstName  string `json:"firstName" binding:"required"`
	LastName   string `json:"lastName" binding:"required"`
	LanguageId int    `json:"languageId" binding:"required"`
}

type resetPasswordInput struct {
	UserId   string `json:"userId" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Tags Users
// @Description Create User
// @Accept json
// @Param data body createUserInput true "userType: 用來區分用戶類別，StoreEase商家為0, OrderEase客戶為1<br><br>languageId: 區分多語系，en為1, zh為2<br><br>"
// @Produce json
// @Success 200 {string} string userId
// @Router /users [post]
func (h *Handler) createUser(c *gin.Context) {
	var inp createUserInput
	if err := c.BindJSON(&inp); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	userId, err := h.services.Users.Create(c.Request.Context(), service.CreateUserInput{
		Email:      inp.Email,
		Password:   inp.Password,
		UserType:   inp.UserType,
		FirstName:  inp.FirstName,
		LastName:   inp.LastName,
		LanguageId: inp.LanguageId,
	})

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, userId)
}

// @Tags Users
// @Description Update User
// @Accept json
// @Param data body updateUserInput true "languageId: 區分多語系，en為1, zh為2<br><br>"
// @Produce json
// @Success 200 {string} string userId
// @Router /users [patch]
func (h *Handler) updateUser(c *gin.Context) {
	var inp updateUserInput
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

// @Tags Users
// @Description Delete User
// @Param user_id path string true "User Id"
// @Produce json
// @Success 200 {string} string userId
// @Router /users/{userId} [delete]
func (h *Handler) deleteUser(c *gin.Context) {
	userId := c.Param("user_id")
	if err := h.services.Users.Delete(c.Request.Context(), userId); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, true)
}

// @Tags Users
// @Description 以Firebase uid或Email獲取UserId，可在使用者登入前就使用來區分頁面要跳轉到登入或註冊
// @Param userType query int true "用來區分用戶類別，StoreEase商家為0, OrderEase客戶為1"
// @Param method query string true "可選擇以 uid 或 email 進行query"
// @Param uid query string false "method為uid時不可為空"
// @Param email query string false "method為email時不可為空"
// @Produce json
// @Success 200 {string} string userId
// @Router /users/login [get]
func (h *Handler) login(c *gin.Context) {
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

// @Tags Users
// @Description Get user's data by id
// @Param user_id path string true "User Id"
// @Produce json
// @Success 200 {object} domain.User
// @Router /users/{userId} [get]
func (h *Handler) getUserById(c *gin.Context) {
	userId := c.Param("user_id")
	user, err := h.services.Users.GetById(c.Request.Context(), userId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

// @Tags Users
// @Description Reset user's password
// @Accept json
// @Param data body resetPasswordInput true "JSON data"
// @Produce json
// @Success 200 {string} string userId
// @Router /users/reset-password [post]
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
