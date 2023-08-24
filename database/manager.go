package database

import (
	"fmt"
	"log"
	"ordering-system-backend/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	appConfig := config.Env
	// 設定資料庫連線字串
	dsn := fmt.Sprintf("%s:%s@/%s", appConfig.UserName, appConfig.Password, appConfig.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	// err = db.Ping()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// db.SetConnMaxLifetime(time.Duration(10) * time.Second)
	// db.SetMaxIdleConns(5)
	// db.SetMaxOpenConns(2)

	fmt.Println("Connected to MySQL database!")
	return db
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
