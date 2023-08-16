package db

import (
	"database/sql"
	"fmt"
	"log"
	"ordering-system-backend/config"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DB *sql.DB
)

func GetDBConnection() string {
	appConfig := config.AppConfig{
		UserName: "root",
		Password: "0000",
		DBName:   "ordering_db",
	}

	// 設定資料庫連線字串
	dsn := fmt.Sprintf("%s:%s@/%s", appConfig.UserName, appConfig.Password, appConfig.DBName)
	fmt.Println(dsn)
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

	// defer DB.Close()

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	DB.SetConnMaxLifetime(time.Duration(10) * time.Second)
	DB.SetMaxIdleConns(5)
	DB.SetMaxOpenConns(2)

	fmt.Println("Connected to MySQL database!")
}
