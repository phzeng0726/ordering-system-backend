package config

import (
	"os"
	"time"
)

var (
	Env *AppConfig
)

const (
	dbConnMaxIdleTime = 3 * time.Minute
	dbConnMaxLifetime = 6 * time.Minute
	dbDialTimeout     = "20s"
	dbMaxIdleConns    = 20
	dbMaxOpenConns    = 100
	dbSlowThreshold   = 200 * time.Millisecond
	keyDBName         = "DB_NAME"
	keyDBUser         = "DB_USER"
	keyDBPass         = "DB_PASS"
	keyPort           = "PORT"
	keyHost           = "HOST"
)

type AppConfig struct {
	DBName   string
	UserName string
	Password string
	Host     string
	Port     string
}

func InitConfig() {
	Env = &AppConfig{
		DBName:   os.Getenv(keyDBName),
		UserName: os.Getenv(keyDBUser),
		Password: os.Getenv(keyDBPass),
		Host:     os.Getenv(keyHost),
		Port:     os.Getenv(keyPort),
	}
}
