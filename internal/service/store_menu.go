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

// TODO Refactor
func (s *StoreMenusService) GetMenuByStoreId(ctx context.Context, userId string, storeId string, languageId int) (domain.Menu, error) {
	var menu domain.Menu
	menuItemMappings, err := s.repo.GetMenuByStoreId(ctx, userId, storeId, languageId)
	if err != nil {
		return menu, err
	}

	if len(menuItemMappings) != 0 {
		menu = menuItemMappings[0].Menu
		for _, mim := range menuItemMappings {
			mim.MenuItem.ImageBytes = mim.MenuItem.Image.BytesData
			mim.MenuItem.Category.Title = mim.MenuItem.Category.CategoryLanguage.Title
			menu.MenuItems = append(menu.MenuItems, mim.MenuItem)
		}

		return menu, nil
	} else {
		// 撈出
		menu, err = s.repo.TempGetAllByUserId(ctx, userId, storeId)
		if err != nil {
			return menu, err
		}

		menu.MenuItems = []domain.MenuItem{}
	}

	return menu, nil
}
