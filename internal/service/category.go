package service

import (
	"context"
	"ordering-system-backend/internal/domain"
	"ordering-system-backend/internal/repository"
)

type CategoriesService struct {
	repo repository.Categories
}

func NewCategoriesService(repo repository.Categories) *CategoriesService {
	return &CategoriesService{repo: repo}
}

func (s *CategoriesService) GetAllByUserId(ctx context.Context, userId string, languageId int) ([]domain.Category, error) {
	categories, err := s.repo.GetAllByUserId(ctx, userId, languageId)
	if err != nil {
		return categories, err
	}
	return categories, nil
}
