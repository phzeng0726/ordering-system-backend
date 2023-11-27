package service

import (
	"context"
	"ordering-system-backend/internal/domain"
	"ordering-system-backend/internal/repository"
)

type StoreMenusService struct {
	storeMenusRepo repository.StoreMenus
	storesRepo     repository.Stores
}

func NewStoreMenusService(storeMenusRepo repository.StoreMenus, storesRepo repository.Stores) *StoreMenusService {
	return &StoreMenusService{
		storeMenusRepo: storeMenusRepo,
		storesRepo:     storesRepo,
	}
}

func (s *StoreMenusService) CreateMenuReference(ctx context.Context, userId string, storeId string, menuId string) error {
	var storeMenuMapping domain.StoreMenuMapping
	storeMenuMapping.StoreId = storeId
	storeMenuMapping.MenuId = menuId

	if err := s.storeMenusRepo.CreateMenuReference(ctx, userId, storeMenuMapping); err != nil {
		return err
	}

	return nil
}

func (s *StoreMenusService) UpdateMenuReference(ctx context.Context, userId string, storeId string, menuId string) error {
	var storeMenuMapping domain.StoreMenuMapping
	storeMenuMapping.StoreId = storeId
	storeMenuMapping.MenuId = menuId

	if err := s.storeMenusRepo.UpdateMenuReference(ctx, userId, storeMenuMapping); err != nil {
		return err
	}

	return nil
}

func (s *StoreMenusService) DeleteMenuReference(ctx context.Context, userId string, storeId string, menuId string) error {
	var storeMenuMapping domain.StoreMenuMapping
	storeMenuMapping.StoreId = storeId
	storeMenuMapping.MenuId = menuId

	if err := s.storeMenusRepo.DeleteMenuReference(ctx, userId, storeMenuMapping); err != nil {
		return err
	}

	return nil
}

func (s *StoreMenusService) GetMenuByStoreId(ctx context.Context, userId string, storeId string, languageId int, userType int) (domain.Menu, error) {
	menu, err := s.storeMenusRepo.GetMenuByStoreId(ctx, userId, storeId, languageId)
	if err != nil {
		return menu, err
	}

	// 撈取商店資訊，供客戶端使用
	if userType == 1 {
		store, err := s.storesRepo.GetByStoreId(ctx, storeId)
		if err != nil {
			return menu, err
		}
		menu.Store = &store
	}

	return menu, nil
}
