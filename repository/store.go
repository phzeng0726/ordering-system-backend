package repository

import (
	"errors"
	"ordering-system-backend/domain"

	"gorm.io/gorm"
)

type StoresRepo struct {
	db *gorm.DB
}

func NewStoresRepo(db *gorm.DB) *StoresRepo {
	return &StoresRepo{
		db: db,
	}
}

func (r *StoresRepo) Create(s domain.Store) error {
	res := r.db.Create(&s)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (r *StoresRepo) Update(s domain.Store) error {
	var store domain.Store
	res := r.db.Model(&store).Where("id = ?", s.Id).Updates(domain.Store{
		Name:        s.Name,
		Description: s.Description,
		// Email:       s.Email,
		Phone:   s.Phone,
		Address: s.Address,
	})

	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (r *StoresRepo) Delete(storeId string) error {
	var store domain.Store
	res := r.db.Where("id = ?", storeId).Delete(&store)

	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (r *StoresRepo) GetAll() ([]domain.Store, error) {
	var stores []domain.Store

	res := r.db.Find(&stores)
	if res.Error != nil {
		return nil, res.Error
	}
	return stores, nil
}

func (r *StoresRepo) GetById(storeId string) (domain.Store, error) {
	var store domain.Store
	res := r.db.Where("id = ?", storeId).First(&store)

	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return store, errors.New("store not found")
		}
		return store, res.Error
	}

	return store, nil
}
