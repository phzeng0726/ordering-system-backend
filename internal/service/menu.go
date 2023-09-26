package service

import (
	"context"
	"ordering-system-backend/internal/domain"
	"ordering-system-backend/internal/repository"

	"github.com/google/uuid"
)

type MenusService struct {
	repo repository.Menus
}

func NewMenusService(repo repository.Menus) *MenusService {
	return &MenusService{repo: repo}
}

func (s *MenusService) Create(ctx context.Context, menu domain.Menu) (string, error) {
	menu.Id = uuid.New().String()
	if err := s.repo.Create(ctx, menu); err != nil {
		return menu.Id, err
	}

	return menu.Id, nil
}

func (s *MenusService) Update(ctx context.Context, menu domain.Menu) error {
	return nil
}

func (s *MenusService) Delete(ctx context.Context, userId string, menuId int) error {
	return nil
}

func (s *MenusService) GetAllByUserId(ctx context.Context, userId string) ([]domain.Menu, error) {
	return []domain.Menu{}, nil
}

func (s *MenusService) GetById(ctx context.Context, userId string, menuId int) (domain.Menu, error) {
	return domain.Menu{}, nil
}
