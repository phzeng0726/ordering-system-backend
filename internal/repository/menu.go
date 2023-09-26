package repository

import (
	"context"
	"errors"
	"ordering-system-backend/internal/domain"

	"gorm.io/gorm"
)

type MenusRepo struct {
	db *gorm.DB
}

func NewMenusRepo(db *gorm.DB) *MenusRepo {
	return &MenusRepo{
		db: db,
	}
}

func (r *MenusRepo) Create(ctx context.Context, menu domain.Menu) error {
	if err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", menu.UserId).First(&domain.User{}).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("user id not found")
			}
			return err
		}

		// 確認該User存在，才可新增執行後續
		if err := tx.Create(&menu).Error; err != nil {
			return err
		}

		for _, mi := range menu.MenuItems {
			if err := tx.Create(&mi).Error; err != nil {
				return err
			}

			mapping := domain.MenuItemMapping{
				MenuId:     menu.Id,
				MenuItemId: mi.Id,
			}

			if err := tx.Create(&mapping).Error; err != nil {
				return err
			}

		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (r *MenusRepo) Update(ctx context.Context, menu domain.Menu) error {
	// tx := r.db.Begin()
	// defer func() {
	// 	if tx.Error != nil {
	// 		tx.Rollback()
	// 	} else {
	// 		tx.Commit()
	// 	}
	// }()

	// if err := tx.Model(&domain.Menu{}).Where("id = ?", m.Id).Updates(map[string]interface{}{
	// 	"title":       m.Title,
	// 	"description": m.Description,
	// 	"is_hide":     m.IsHide,
	// }).Error; err != nil {
	// 	return err
	// }

	// if err := tx.Where("menu_id = ?", m.Id).Delete(&domain.MenuItemMapping{}).Error; err != nil {
	// 	return err
	// }

	// for _, mi := range m.MenuItems {
	// 	if err := tx.Create(&mi).Error; err != nil {
	// 		return err
	// 	}
	// 	mapping := domain.MenuItemMapping{
	// 		MenuId:     m.Id,
	// 		MenuItemId: mi.Id,
	// 	}
	// 	if err := tx.Create(&mapping).Error; err != nil {
	// 		return err
	// 	}
	// }

	return nil
}

func (r *MenusRepo) Delete(ctx context.Context, userId string, menuId string) error {
	// tx := r.db.Begin()
	// defer func() {
	// 	if tx.Error != nil {
	// 		tx.Rollback()
	// 	} else {
	// 		tx.Commit()
	// 	}
	// }()

	// if err := tx.Where("menu_id = ?", menuId).Delete(&domain.MenuItemMapping{}).Error; err != nil {
	// 	return err
	// }

	// if err := tx.Where("id = ?", menuId).Delete(&domain.Menu{}).Error; err != nil {
	// 	return err
	// }

	return nil
}

func (r *MenusRepo) GetAllByUserId(ctx context.Context, userId string) ([]domain.Menu, error) {
	var menus []domain.Menu
	var menuItemMappings []domain.MenuItemMapping

	if err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", userId).First(&domain.User{}).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("user id not found")
			}
			return err
		}

		// 確認該User存在，才可新增執行後續
		if err := tx.Where("user_id = ?", userId).Find(&menus).Error; err != nil {
			return err
		}

		if len(menus) == 0 {
			return errors.New("menu list is empty")
		}

		if err := tx.Preload("Menu").Preload("MenuItem.Category").Where("menu_id IN (SELECT id FROM menus WHERE user_id = ?)", userId).Find(&menuItemMappings).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return menus, err
	}

	menuItemsIdMap := make(map[string][]domain.MenuItem)
	for _, mim := range menuItemMappings {
		menuItemsIdMap[mim.MenuId] = append(menuItemsIdMap[mim.MenuId], mim.MenuItem)
	}

	for i, menu := range menus {
		menus[i].MenuItems = menuItemsIdMap[menu.Id]
	}

	return menus, nil
}

func (r *MenusRepo) GetById(ctx context.Context, userId string, menuId string) (domain.Menu, error) {
	var menu domain.Menu
	var menuItemMappings []domain.MenuItemMapping
	if err := r.db.WithContext(ctx).Preload("Menu").Preload("MenuItem.Category").Where("menu_id = ?", menuId).Find(&menuItemMappings).Error; err != nil {
		return menu, err
	}

	if len(menuItemMappings) == 0 {
		err := errors.New("menu with items not found")
		return menu, err
	}

	menu = menuItemMappings[0].Menu
	for _, mim := range menuItemMappings {
		menu.MenuItems = append(menu.MenuItems, mim.MenuItem)
	}

	return menu, nil
}
