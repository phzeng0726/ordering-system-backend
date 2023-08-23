package service

import (
	"net/http"
	"ordering-system-backend/models"
	"ordering-system-backend/repository"

	"github.com/gin-gonic/gin"
)

type StoresService struct {
	repo repository.Stores
}

func NewStoresService(repo repository.Stores) *StoresService {
	return &StoresService{repo: repo}
}

func (s *StoresService) Create(c *gin.Context) {
	var newStore models.Store

	if err := c.BindJSON(&newStore); err != nil {
		return
	}
	err := s.repo.Create(newStore)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, nil)
}

func (s *StoresService) Update(c *gin.Context) {}

func (s *StoresService) Delete(c *gin.Context) {}

func (s *StoresService) GetAll(c *gin.Context) {
	stores, err := s.repo.GetAll()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, stores)
}

func (s *StoresService) GetById(c *gin.Context) {
	id := c.Param("store_id")
	stores, err := s.repo.GetById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, stores)
}
