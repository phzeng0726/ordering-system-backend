package repository

import (
	"database/sql"
	"ordering-system-backend/models"
)

type Stores interface {
	GetStores() ([]models.Store, error)
	GetStoreById(storeId string) (models.Store, error)
	Create(s models.Store) error
}

type Menus interface {
	GetMenus(storeId string) ([]models.Menu, error)
	GetMenuById(storeId string, menuId int) (models.Menu, error)
	Create(m models.Menu) error
	Update(m models.Menu) error
}

type Repositories struct {
	Menus  Menus
	Stores Stores
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		Menus:  NewMenusRepo(db),
		Stores: NewStoresRepo(db),
	}
}
