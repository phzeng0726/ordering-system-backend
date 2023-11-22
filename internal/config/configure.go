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
	keyPort              = "PORT"
	keyHost              = "HOST"
	keyOTPSenderEmail    = "OTP_SENDER_EMAIL"
	keyOTPSenderPassword = "OTP_SENDER_PASSWORD"
)

type AppConfig struct {
	DBName            string
	DBUser            string
	DBPassword        string
	Host              string
	Port              string
	OTPSenderEmail    string
	OTPSenderPassword string
}

func InitConfig() {
	Env = &AppConfig{
		DBName:            os.Getenv(keyDBName),
		DBUser:            os.Getenv(keyDBUser),
		DBPassword:        os.Getenv(keyDBPass),
		Host:              os.Getenv(keyHost),
		Port:              os.Getenv(keyPort),
		OTPSenderEmail:    os.Getenv(keyOTPSenderEmail),
		OTPSenderPassword: os.Getenv(keyOTPSenderPassword),
	}
}
