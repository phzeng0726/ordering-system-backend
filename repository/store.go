package repository

import (
	"database/sql"
	"errors"
	"ordering-system-backend/domain"
)

type StoresRepo struct {
	db *sql.DB
}

func NewStoresRepo(db *sql.DB) *StoresRepo {
	return &StoresRepo{
		db: db,
	}
}

func (r *StoresRepo) Create(s domain.Store) error {
	sql := "INSERT INTO store (id, name, description, email, phone, is_open)" +
		" VALUES (?, ?, ?, ?, ?, ?)"

	_, err := r.db.Exec(sql, s.Id, s.Name, s.Description, s.Email, s.Phone, s.IsOpen)

	if err != nil {
		return err
	}

	return nil
}

func (r *StoresRepo) Update(s domain.Store) error {
	sql := "UPDATE store" +
		" SET `name` = ?, `description` = ?, email = ?, phone = ?, is_open = ?" +
		" WHERE id = ?"

	_, err := r.db.Exec(sql, s.Name, s.Description, s.Email, s.Phone, s.IsOpen, s.Id)

	if err != nil {
		return err
	}

	return nil
}
func (r *StoresRepo) Delete(storeId string) error {
	sql := "DELETE FROM store WHERE id = ?"
	_, err := r.db.Exec(sql, storeId)

	if err != nil {
		return err
	}

	return nil
}

func (r *StoresRepo) GetAll() ([]domain.Store, error) {
	var stores []domain.Store

	sql := "SELECT * FROM store"
	rows, err := r.db.Query(sql)
	if err != nil {
		return stores, err
	}
	defer rows.Close()

	for rows.Next() {
		var store domain.Store
		err := rows.Scan(
			&store.Id,
			&store.Name,
			&store.Description,
			&store.Email,
			&store.Phone,
			&store.IsOpen,
		)
		if err != nil {
			return stores, err
		}
		stores = append(stores, store)
	}

	return stores, nil
}

func (r *StoresRepo) GetById(storeId string) (domain.Store, error) {
	var store domain.Store

	sql := "SELECT *" +
		" FROM store" +
		" WHERE id = ?" +
		" LIMIT 1"
	rows, err := r.db.Query(sql, storeId)
	if err != nil {
		return store, err
	}
	defer rows.Close()

	found := false // 用於標記是否找到符合條件的資料行
	for rows.Next() {
		found = true // 找到資料行
		err := rows.Scan(
			&store.Id,
			&store.Name,
			&store.Description,
			&store.Email,
			&store.Phone,
			&store.IsOpen,
		)
		if err != nil {
			return store, err
		}
	}

	if !found {
		err = errors.New("store not found")
		return store, err
	}

	return store, nil
}
