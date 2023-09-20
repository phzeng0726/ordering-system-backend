package repository

import (
	"errors"
	"ordering-system-backend/internal/domain"

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
	res := r.db.Model(&store).Where("id = ?", s.Id).Updates(&s)

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

	// 使用 Preload 一次載入所有的OpeningHours，避免N+1問題
	if err := r.db.Preload("StoreOpeningHours").Find(&stores).Error; err != nil {
		return nil, err
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

	if err := r.db.Where("store_id = ?", storeId).Find(&store.StoreOpeningHours).Error; err != nil {
		return store, err
	}

	return store, nil
}
