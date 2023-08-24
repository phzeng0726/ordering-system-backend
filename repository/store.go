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
	// sql := "INSERT INTO store (id, name, description, email, phone, is_open)" +
	// 	" VALUES (?, ?, ?, ?, ?, ?)"

	// _, err := r.db.Exec(sql, s.Id, s.Name, s.Description, s.Email, s.Phone, s.IsOpen

	// if err != nil {
	// 	return err
	// }
	res := r.db.Create(&s)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (r *StoresRepo) Update(s domain.Store) error {
	// sql := "UPDATE store" +
	// 	" SET `name` = ?, `description` = ?, email = ?, phone = ?, is_open = ?" +
	// 	" WHERE id = ?"

	// _, err := r.db.Exec(sql, s.Name, s.Description, s.Email, s.Phone, s.IsOpen, s.Id)

	// if err != nil {
	// 	return err
	// }
	// return nil
	var store domain.Store
	res := r.db.Model(&store).Where("id = ?", s.Id).Updates(domain.Store{
		Name:        s.Name,
		Description: s.Description,
		Email:       s.Email,
		Phone:       s.Phone,
		IsOpen:      s.IsOpen,
	})

	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (r *StoresRepo) Delete(storeId string) error {
	// sql := "DELETE FROM store WHERE id = ?"
	// _, err := r.db.Exec(sql, storeId)

	// if err != nil {
	// 	return err
	// }

	// return nil
	var store domain.Store
	res := r.db.Where("id = ?", storeId).Delete(&store)

	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (r *StoresRepo) GetAll() ([]domain.Store, error) {
	// var stores []domain.Store
	// sql := "SELECT * FROM store"
	// rows, err := r.db.Query(sql)
	// if err != nil {
	// 	return stores, err
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	var store domain.Store
	// 	err := rows.Scan(
	// 		&store.Id,
	// 		&store.Name,
	// 		&store.Description,
	// 		&store.Email,
	// 		&store.Phone,
	// 		&store.IsOpen,
	// 	)
	// 	if err != nil {
	// 		return stores, err
	// 	}
	// 	stores = append(stores, store)
	// }
	var stores []domain.Store

	res := r.db.Find(&stores)
	if res.Error != nil {
		return nil, res.Error
	}
	return stores, nil
}

func (r *StoresRepo) GetById(storeId string) (domain.Store, error) {
	// var store domain.Store

	// sql := "SELECT *" +
	// 	" FROM store" +
	// 	" WHERE id = ?" +
	// 	" LIMIT 1"
	// rows, err := r.db.Query(sql, storeId)
	// if err != nil {
	// 	return store, err
	// }
	// defer rows.Close()

	// found := false // 用於標記是否找到符合條件的資料行
	// for rows.Next() {
	// 	found = true // 找到資料行
	// 	err := rows.Scan(
	// 		&store.Id,
	// 		&store.Name,
	// 		&store.Description,
	// 		&store.Email,
	// 		&store.Phone,
	// 		&store.IsOpen,
	// 	)
	// 	if err != nil {
	// 		return store, err
	// 	}
	// }

	// if !found {
	// 	err = errors.New("store not found")
	// 	return store, err
	// }

	// return store, nil
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
