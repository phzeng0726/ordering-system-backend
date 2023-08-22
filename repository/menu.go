package repository

import (
	"database/sql"
	"fmt"
	"ordering-system-backend/models"
	"ordering-system-backend/utils"
)

type MenusRepo struct {
	db *sql.DB
}

func NewMenusRepo(db *sql.DB) *MenusRepo {
	return &MenusRepo{
		db: db,
	}
}

func scanMenusRow(rows *sql.Rows) ([]models.Menu, error) {
	var menus []models.Menu
	var err error

	for rows.Next() {
		var menu models.Menu
		var createAtStr string // 創建一個字串來暫存日期時間字串

		err = rows.Scan(
			&menu.Id,
			&menu.StoreId,
			&menu.Title,
			&menu.Description,
			&menu.IsHide,
			&createAtStr, // 接收日期時間字串
		)
		if err != nil {
			return menus, err
		}

		menu.CreateAt, err = utils.DateTimeConverter(createAtStr)
		if err != nil {
			return menus, err
		}
		menus = append(menus, menu)
	}

	return menus, err
}

func (r *MenusRepo) GetMenus(storeId string) ([]models.Menu, error) {
	var menus []models.Menu

	sql := "SELECT *" +
		" FROM menu" +
		" WHERE store_id = ?"
	rows, err := r.db.Query(sql, storeId)

	if err != nil {
		fmt.Println(err)
		return menus, err
	}
	defer rows.Close()

	menus, err = scanMenusRow(rows)

	if err != nil {
		return menus, err
	}

	return menus, err
}
