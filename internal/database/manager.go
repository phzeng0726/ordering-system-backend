package database

import (
	"fmt"
	"log"
	"ordering-system-backend/internal/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	appConfig := config.Env
	// 設定資料庫連線字串
	dsn := fmt.Sprintf("%s:%s@/%s?parseTime=true", appConfig.DBUser, appConfig.DBPassword, appConfig.DBName)

	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Connected to MySQL database!")
	return conn
}
