package repository

import (
	"database/sql"
	"ordering-system-backend/models"
)

type Stores interface {
	Create(s models.Store) error
	GetAll() ([]models.Store, error)
	GetById(storeId string) (models.Store, error)
}

type Menus interface {
	Create(m models.Menu) error
	Update(m models.Menu) error
	GetAll(storeId string) ([]models.Menu, error)
	GetById(storeId string, menuId int) (models.Menu, error)
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
