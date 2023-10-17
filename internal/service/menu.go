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
				Image: domain.Image{
					BytesData: mi.ImageBytes,
				},
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
				Image: domain.Image{
					BytesData: mi.ImageBytes,
				},
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
	if err := s.repo.Delete(ctx, userId, menuId); err != nil {
		return err
	}
	return nil
}

func (s *MenusService) GetAllByUserId(ctx context.Context, userId string, languageId int) ([]domain.Menu, error) {
	menus := make([]domain.Menu, 0)
	menuItemMappings, err := s.repo.GetAllByUserId(ctx, userId, languageId)
	if err != nil {
		return menus, err
	}

	menuItemsIdMap := make(map[string][]domain.MenuItem)
	menuMap := make(map[string]struct{}) // 使用map追蹤已經處理過的 menu

	for _, mim := range menuItemMappings {
		// 檢查是否已經處理過該 menu
		if _, ok := menuMap[mim.Menu.Id]; !ok {
			menuMap[mim.Menu.Id] = struct{}{}
			menus = append(menus, mim.Menu)
		}

		// key: menuId, value: menuItems
		mim.MenuItem.ImageBytes = mim.MenuItem.Image.BytesData
		mim.MenuItem.Category.Title = mim.MenuItem.Category.CategoryLanguage.Title
		menuItemsIdMap[mim.MenuId] = append(menuItemsIdMap[mim.MenuId], mim.MenuItem)
	}

	// 將 menuItems 加入 menu 中
	for i, menu := range menus {
		menus[i].MenuItems = menuItemsIdMap[menu.Id]
	}

	return menus, nil
}

func (s *MenusService) GetById(ctx context.Context, userId string, menuId string, languageId int) (domain.Menu, error) {
	var menu domain.Menu
	menuItemMappings, err := s.repo.GetById(ctx, userId, menuId, languageId)
	if err != nil {
		return menu, err
	}

	menu = menuItemMappings[0].Menu
	for _, mim := range menuItemMappings {
		mim.MenuItem.ImageBytes = mim.MenuItem.Image.BytesData
		mim.MenuItem.Category.Title = mim.MenuItem.Category.CategoryLanguage.Title
		menu.MenuItems = append(menu.MenuItems, mim.MenuItem)
	}

	return menu, nil
}
