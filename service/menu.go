package service

import (
	"net/http"
	"ordering-system-backend/domain"
	"ordering-system-backend/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MenusService struct {
	repo repository.Menus
}

func NewMenusService(repo repository.Menus) *MenusService {
	return &MenusService{repo: repo}
}

func (s *MenusService) Create(c *gin.Context) {
	var newMenu domain.Menu

	if err := c.BindJSON(&newMenu); err != nil {
		return
	}

	err := s.repo.Create(newMenu)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, newMenu)
}

func (s *MenusService) Update(c *gin.Context) {
	var newMenu domain.Menu
	if err := c.BindJSON(&newMenu); err != nil {
		return
	}

	err := s.repo.Update(newMenu)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, newMenu)
}

func (s *MenusService) Delete(c *gin.Context) {
	storeId := c.Param("store_id")
	menuIdStr := c.Param("menu_id")
	menuId, err := strconv.Atoi(menuIdStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid menu_id"})
		return
	}

	err = s.repo.Delete(storeId, menuId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, nil)
}

func (s *MenusService) GetAll(c *gin.Context) {
	storeId := c.Param("store_id")
	menus, err := s.repo.GetAll(storeId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, menus)
}

func (s *MenusService) GetById(c *gin.Context) {
	storeId := c.Param("store_id")
	menuIdStr := c.Param("menu_id")
	menuId, err := strconv.Atoi(menuIdStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid menu_id"})
		return
	}

	menus, err := s.repo.GetById(storeId, menuId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, menus)
}
