package config

import (
	"os"
	"time"
)

var (
	Env *AppConfig
)

const (
	dbConnMaxIdleTime    = 3 * time.Minute
	dbConnMaxLifetime    = 6 * time.Minute
	dbDialTimeout        = "20s"
	dbMaxIdleConns       = 20
	dbMaxOpenConns       = 100
	dbSlowThreshold      = 200 * time.Millisecond
	keyDBName            = "DB_NAME"
	keyDBUser            = "DB_USER"
	keyDBPass            = "DB_PASS"
	keyDBHost            = "DB_HOST"
	keyDBPort            = "DB_PORT"
	keyHost              = "HOST"
	keyPort              = "PORT"
	keyOTPSenderEmail    = "OTP_SENDER_EMAIL"
	keyOTPSenderPassword = "OTP_SENDER_PASSWORD"
)

type AppConfig struct {
	DBName            string
	DBUser            string
	DBPassword        string
	DBHost            string
	DBPort            string
	Host              string
	Port              string
	OTPSenderEmail    string
	OTPSenderPassword string
	IsOnCloud         bool
}

func InitConfig() {
	Env = &AppConfig{
		DBName:            os.Getenv(keyDBName),
		DBUser:            os.Getenv(keyDBUser),
		DBPassword:        os.Getenv(keyDBPass),
		DBHost:            os.Getenv(keyDBHost),
		DBPort:            os.Getenv(keyDBPort),
		Host:              os.Getenv(keyHost),
		Port:              os.Getenv(keyPort),
		OTPSenderEmail:    os.Getenv(keyOTPSenderEmail),
		OTPSenderPassword: os.Getenv(keyOTPSenderPassword),
		IsOnCloud:         os.Getenv(keyHost) == "", // Host沒有填的時候就是Cloud (GCP上不需要填Host)
	}
}
