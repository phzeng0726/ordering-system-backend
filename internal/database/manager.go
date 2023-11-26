package database

import (
	"fmt"
	"log"
	"ordering-system-backend/internal/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	var dsn string
	appConfig := config.Env

	// 設定資料庫連線字串，Host為空代表在Cloud Run跑，否則在local跑
	if appConfig.IsOnCloud {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", appConfig.DBUser, appConfig.DBPassword, appConfig.DBHost, appConfig.DBPort, appConfig.DBName)
	} else {
		dsn = fmt.Sprintf("%s:%s@/%s?parseTime=true", appConfig.DBUser, appConfig.DBPassword, appConfig.DBName)
	}

	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Connected to MySQL database!")
	return conn
}
