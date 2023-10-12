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

func (s *MenusService) Create(ctx context.Context, input CreateMenuInput) (string, error) {
	var menuItems []domain.MenuItem
	for _, mi := range input.MenuItems {
		menuItems = append(
			menuItems,
			domain.MenuItem{
				Title:       mi.Title,
				Description: mi.Description,
				Quantity:    mi.Quantity,
				Price:       mi.Price,
				CategoryId:  mi.CategoryId,
			},
		)
	}

	menu := domain.Menu{
		Id:          uuid.New().String(),
		UserId:      input.UserId,
		Title:       input.Title,
		Description: input.Description,
		MenuItems:   menuItems,
	}

	if err := s.repo.Create(ctx, menu); err != nil {
		return menu.Id, err
	}
	return menu.Id, nil
}

func (s *MenusService) Update(ctx context.Context, input UpdateMenuInput) error {
	var menuItems []domain.MenuItem
	for _, mi := range input.MenuItems {
		menuItems = append(
			menuItems,
			domain.MenuItem{
				Title:       mi.Title,
				Description: mi.Description,
				Quantity:    mi.Quantity,
				Price:       mi.Price,
				CategoryId:  mi.CategoryId,
			},
		)
	}

	menu := domain.Menu{
		Id:          input.MenuId,
		UserId:      input.UserId,
		Title:       input.Title,
		Description: input.Description,
		MenuItems:   menuItems,
	}

	if err := s.repo.Update(ctx, menu); err != nil {
		return err
	}
	return nil
}

func (s *MenusService) Delete(ctx context.Context, userId string, menuId string) error {
	return nil
}

func (s *MenusService) GetAllByUserId(ctx context.Context, userId string, languageId int) ([]domain.Menu, error) {
	menus, err := s.repo.GetAllByUserId(ctx, userId, languageId)
	if err != nil {
		return menus, err
	}
	return menus, nil
}

func (s *MenusService) GetById(ctx context.Context, userId string, menuId string, languageId int) (domain.Menu, error) {
	menu, err := s.repo.GetById(ctx, userId, menuId, languageId)
	if err != nil {
		return menu, err
	}
	return menu, nil
}
