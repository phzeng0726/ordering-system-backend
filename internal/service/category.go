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
	var categories []domain.Category

	categoryUserMappings, err := s.repo.GetAllByUserId(ctx, userId, languageId)
	if err != nil {
		return categories, err
	}

	for _, cum := range categoryUserMappings {
		cum.Category.Title = cum.Category.CategoryLanguage.Title
		categories = append(categories, cum.Category)
	}

	return categories, nil
}
