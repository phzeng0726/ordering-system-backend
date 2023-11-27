package service

import (
	"context"
	"errors"
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

// TODO Refactor
func (s *MenusService) GetAllByUserId(ctx context.Context, userId string, languageId int) ([]domain.Menu, error) {
	menus := make([]domain.Menu, 0)

	// 確認使用者是否存在
	if _, err := s.usersRepo.GetById(ctx, userId); err != nil {
		return menus, err
	}

	menuItemMappings, err := s.menusRepo.GetAllByUserId(ctx, userId, languageId)
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

	// 撈出所有
	tempMenus, err := s.menusRepo.TempGetAllByUserId(ctx, userId, languageId)
	if err != nil {
		return menus, err
	}

	// 沒有menuItems的也要放進來
	for _, tm := range tempMenus {
		if _, ok := menuMap[tm.Id]; !ok {
			tm.MenuItems = []domain.MenuItem{}
			menus = append(menus, tm)
		}
	}

	return menus, nil
}

// TODO Refactor
func (s *MenusService) GetById(ctx context.Context, userId string, menuId string, languageId int) (domain.Menu, error) {
	var menu domain.Menu
	menuItemMappings, err := s.menusRepo.GetById(ctx, userId, menuId, languageId)
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
	}

	// TODO 重構
	// 確認使用者是否存在
	if _, err := s.usersRepo.GetById(ctx, userId); err != nil {
		return menu, err
	}

	// 撈出所有
	tempMenus, err := s.menusRepo.TempGetAllByUserId(ctx, userId, languageId)
	if err != nil {
		return menu, err
	}

	if len(tempMenus) != 0 {
		for _, m := range tempMenus {
			if m.Id == menuId {
				m.MenuItems = []domain.MenuItem{}
				return m, nil
			}
		}
	}

	return menu, errors.New("menu not found")
}
