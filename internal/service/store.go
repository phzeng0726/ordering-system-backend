package service

import (
	"net/http"
	"ordering-system-backend/internal/domain"
	"ordering-system-backend/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type StoresService struct {
	repo repository.Stores
}

func NewStoresService(repo repository.Stores) *StoresService {
	return &StoresService{repo: repo}
}

func (s *StoresService) Create(c *gin.Context) {
	var newStore domain.Store
	if err := c.BindJSON(&newStore); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	uuid := uuid.New()
	newStore.Id = uuid.String()

	err := s.repo.Create(newStore)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, newStore)
}

func (s *StoresService) Update(c *gin.Context) {
	var newStore domain.Store
	id := c.Param("store_id")

	if err := c.BindJSON(&newStore); err != nil {
		return
	}

	newStore.Id = id
	err := s.repo.Update(newStore)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, newStore)
}

func (s *StoresService) Delete(c *gin.Context) {
	id := c.Param("store_id")

	err := s.repo.Delete(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, true)
}

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
