package database

import (
	"database/sql"
	"fmt"
	"log"
	"ordering-system-backend/config"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Manager struct {
	db *sql.DB
}

var DB *sql.DB

func GetDBConnection() string {
	appConfig := config.Env

	// 設定資料庫連線字串
	dsn := fmt.Sprintf("%s:%s@/%s", appConfig.UserName, appConfig.Password, appConfig.DBName)

	return dsn
}

func Connect() {
	dsn := GetDBConnection()

	// 連線資料庫
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	DB.SetConnMaxLifetime(time.Duration(10) * time.Second)
	DB.SetMaxIdleConns(5)
	DB.SetMaxOpenConns(2)

	fmt.Println("Connected to MySQL database!")
}


// func Open() (*Manager, error) {
// 	appConfig := config.Env

// 	// 設定資料庫連線字串
// 	dsn := fmt.Sprintf("%s:%s@/%s", appConfig.UserName, appConfig.Password, appConfig.DBName)
// 	db, err := sql.Open("mysql", dsn)

// 	db.SetConnMaxLifetime(time.Duration(10) * time.Second)
// 	db.SetMaxIdleConns(5)
// 	db.SetMaxOpenConns(2)
// 	return &Manager{db}, err
// }

// func (m *Manager) Close() error {
// 	if m.db != nil {
// 		return m.db.Close()
// 	}
// 	return nil
// }
