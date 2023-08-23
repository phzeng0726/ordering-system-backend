package repository

import (
	"database/sql"
	"ordering-system-backend/domain"
)

type Stores interface {
	Create(s domain.Store) error
	Update(s domain.Store) error
	Delete(storeId string) error
	GetAll() ([]domain.Store, error)
	GetById(storeId string) (domain.Store, error)
}

type Menus interface {
	Create(m domain.Menu) error
	Update(m domain.Menu) error
	Delete(storeId string, menuId int) error
	GetAll(storeId string) ([]domain.Menu, error)
	GetById(storeId string, menuId int) (domain.Menu, error)
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
