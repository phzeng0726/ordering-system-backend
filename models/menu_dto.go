package models

import (
	"database/sql"
)

type MenuDTO struct {
	MenuId       int
	StoreId      string
	Title        string
	Description  string
	IsHide       bool
	MenuItemDTOs []MenuItemDTO
}
type MenuItemDTO struct {
	MenuItemId          int
	MenuItemTitle       string
	MenuItemDescription string
	MenuItemQuantity    int
	MenuItemPrice       int
	MenuCategoryDTO     MenuCategoryDTO
}
type MenuCategoryDTO struct {
	MenuCategoryId    int
	MenuCategoryTitle string
}

func ScanMenuRow(rows *sql.Rows) (MenuDTO, error) {
	var menuDTO MenuDTO
	var menuItemDTOs []MenuItemDTO
	var err error

	for rows.Next() {
		var menuItemDTO MenuItemDTO
		var menuCategoryDTO MenuCategoryDTO

		err = rows.Scan(
			&menuDTO.MenuId,
			&menuDTO.StoreId,
			&menuDTO.Title,
			&menuDTO.Description,
			&menuDTO.IsHide,
			&menuItemDTO.MenuItemId,
			&menuItemDTO.MenuItemTitle,
			&menuItemDTO.MenuItemDescription,
			&menuItemDTO.MenuItemQuantity,
			&menuItemDTO.MenuItemPrice,
			&menuCategoryDTO.MenuCategoryId,
			&menuCategoryDTO.MenuCategoryTitle,
		)
		if err != nil {
			return menuDTO, err
		}
		menuItemDTO.MenuCategoryDTO = menuCategoryDTO
		menuItemDTOs = append(menuItemDTOs, menuItemDTO)
	}

	menuDTO.MenuItemDTOs = menuItemDTOs
	return menuDTO, err
}

func (dto *MenuDTO) ToDomain() Menu {
	var menuItems []MenuItem
	menu := Menu{
		Id:          dto.MenuId,
		StoreId:     dto.StoreId,
		Title:       dto.Title,
		Description: dto.Description,
		IsHide:      dto.IsHide,
	}

	for _, miDto := range dto.MenuItemDTOs {
		menuItem := MenuItem{
			Id:          miDto.MenuItemId,
			Title:       miDto.MenuItemTitle,
			Description: miDto.MenuItemDescription,
			Quantity:    miDto.MenuItemQuantity,
			Price:       miDto.MenuItemPrice,
			MenuCategory: MenuCategory{
				Id:    miDto.MenuCategoryDTO.MenuCategoryId,
				Title: miDto.MenuCategoryDTO.MenuCategoryTitle,
			},
		}
		menuItems = append(menuItems, menuItem)
	}

	menu.MenuItems = menuItems
	return menu
}
