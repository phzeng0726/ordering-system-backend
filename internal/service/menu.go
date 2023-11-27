package service

import (
	"context"
	"ordering-system-backend/internal/domain"
	"ordering-system-backend/internal/repository"

	"github.com/google/uuid"
)

type MenusService struct {
	menusRepo repository.Menus
	usersRepo repository.Users
}

func NewMenusService(menusRepo repository.Menus, usersRepo repository.Users) *MenusService {
	return &MenusService{
		menusRepo: menusRepo,
		usersRepo: usersRepo,
	}
}

func (s *MenusService) Create(ctx context.Context, input CreateMenuInput) (string, error) {
	var menuItems []domain.MenuItem
	newMenuId := uuid.New().String()

	// 確認使用者是否存在
	if _, err := s.usersRepo.GetById(ctx, input.UserId); err != nil {
		return newMenuId, err
	}

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
		Id:          newMenuId,
		UserId:      input.UserId,
		Title:       input.Title,
		Description: input.Description,
		MenuItems:   menuItems,
	}

	if err := s.menusRepo.Create(ctx, menu); err != nil {
		return newMenuId, err
	}

	return newMenuId, nil
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

	if err := s.menusRepo.Update(ctx, menu); err != nil {
		return err
	}
	return nil
}

func (s *MenusService) Delete(ctx context.Context, userId string, menuId string) error {
	if err := s.menusRepo.Delete(ctx, userId, menuId); err != nil {
		return err
	}
	return nil
}

func menuDataClean(menu *domain.Menu) {
	// 沒有menuItems的時候，回傳空的slice
	if len(menu.MenuItemMappings) == 0 {
		menu.MenuItems = []domain.MenuItem{}
	} else {
		// 否則將MenuItem資料Preload的各個項目，撈出需要的變數，處理成需要的格式
		var menuItems []domain.MenuItem
		for j, mim := range menu.MenuItemMappings {
			tempItem := menu.MenuItemMappings[j].MenuItem
			tempItem.ImageBytes = mim.MenuItem.Image.BytesData
			tempItem.Category.Title = mim.MenuItem.Category.CategoryLanguage.Title
			menuItems = append(menuItems, tempItem)
		}

		// 將該menu的menuItems取代為處理過的資料
		menu.MenuItems = menuItems
	}
}

func (s *MenusService) GetAllByUserId(ctx context.Context, userId string, languageId int) ([]domain.Menu, error) {
	menus := make([]domain.Menu, 0)

	// 確認使用者是否存在
	if _, err := s.usersRepo.GetById(ctx, userId); err != nil {
		return menus, err
	}

	// 撈出menus
	menus, err := s.menusRepo.GetAllByUserId(ctx, userId, languageId)
	if err != nil {
		return menus, err
	}

	// 進行資料處理
	for i := range menus {
		menuDataClean(&menus[i])
	}

	return menus, nil
}

func (s *MenusService) GetById(ctx context.Context, userId string, menuId string, languageId int) (domain.Menu, error) {
	var menu domain.Menu

	// 確認使用者是否存在
	if _, err := s.usersRepo.GetById(ctx, userId); err != nil {
		return menu, err
	}

	// 撈出menu
	menu, err := s.menusRepo.GetByMenuId(ctx, menuId, languageId)
	if err != nil {
		return menu, err
	}

	// 進行資料處理
	menuDataClean(&menu)

	return menu, nil
}
