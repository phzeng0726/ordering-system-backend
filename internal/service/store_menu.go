package service

import (
	"context"
	"ordering-system-backend/internal/domain"
	"ordering-system-backend/internal/repository"
)

type StoreMenusService struct {
	repo repository.StoreMenus
}

func NewStoreMenusService(repo repository.StoreMenus) *StoreMenusService {
	return &StoreMenusService{repo: repo}
}

func (s *StoreMenusService) CreateMenuReference(ctx context.Context, userId string, storeId string, menuId string) error {
	var storeMenuMapping domain.StoreMenuMapping
	storeMenuMapping.StoreId = storeId
	storeMenuMapping.MenuId = menuId

	if err := s.repo.CreateMenuReference(ctx, userId, storeMenuMapping); err != nil {
		return err
	}

	return nil
}

func (s *StoreMenusService) UpdateMenuReference(ctx context.Context, userId string, storeId string, menuId string) error {
	var storeMenuMapping domain.StoreMenuMapping
	storeMenuMapping.StoreId = storeId
	storeMenuMapping.MenuId = menuId

	if err := s.repo.UpdateMenuReference(ctx, userId, storeMenuMapping); err != nil {
		return err
	}

	return nil
}

func (s *StoreMenusService) DeleteMenuReference(ctx context.Context, userId string, storeId string) error {
	if err := s.repo.DeleteMenuReference(ctx, userId, storeId); err != nil {
		return err
	}

	return nil
}

func (s *StoreMenusService) GetMenuByStoreId(ctx context.Context, userId string, storeId string, languageId int, userType int) (domain.Menu, error) {
	menu, err := s.repo.GetMenuByStoreId(ctx, userId, storeId, languageId, userType)
	if err != nil {
		return menu, err
	}

	return menu, nil
}
