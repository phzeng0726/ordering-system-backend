package service

import (
	"context"
	"ordering-system-backend/internal/domain"
	"ordering-system-backend/internal/repository"
)

type CategoriesService struct {
	categoriesRepo repository.Categories
	usersRepo      repository.Users
}

func NewCategoriesService(categoriesRepo repository.Categories, usersRepo repository.Users) *CategoriesService {
	return &CategoriesService{
		categoriesRepo: categoriesRepo,
		usersRepo:      usersRepo,
	}
}

func (s *CategoriesService) Create(ctx context.Context, userId string, input CreateCategoryInput) error {
	isDefault := false

	category := domain.Category{
		Identifier: input.Identifier,
		IsDefault:  &isDefault,
	}

	// Create的category沒有language id
	categoryLanguage := domain.CategoryLanguage{
		Title: input.Title,
	}

	categoryUserMapping := domain.CategoryUserMapping{
		UserId: userId,
	}

	if err := s.categoriesRepo.Create(ctx, category, categoryLanguage, categoryUserMapping); err != nil {
		return err
	}

	return nil
}

func (s *CategoriesService) Update(ctx context.Context, userId string, input UpdateCategoryInput) error {
	category := domain.Category{
		Id:         input.CategoryId,
		Identifier: input.Identifier,
	}

	categoryLanguage := domain.CategoryLanguage{
		CategoryId: input.CategoryId,
		Title:      input.Title,
	}

	if err := s.categoriesRepo.Update(ctx, userId, category, categoryLanguage); err != nil {
		return err
	}

	return nil
}

func (s *CategoriesService) Delete(ctx context.Context, userId string, categoryId int) error {
	if err := s.categoriesRepo.Delete(ctx, userId, categoryId); err != nil {
		return err
	}

	return nil
}

func (s *CategoriesService) GetAllByUserId(ctx context.Context, userId string, languageId int) ([]domain.Category, error) {
	var categories []domain.Category

	// 確認用戶存在
	if _, err := s.usersRepo.GetById(ctx, userId); err != nil {
		return categories, err
	}

	categoryUserMappings, err := s.categoriesRepo.GetAllByUserId(ctx, userId, languageId)
	if err != nil {
		return categories, err
	}

	for _, cum := range categoryUserMappings {
		cum.Category.Title = cum.Category.CategoryLanguage.Title
		categories = append(categories, cum.Category)
	}

	return categories, nil
}
