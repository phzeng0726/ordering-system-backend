package config

var (
	Env *AppConfig
)

type AppConfig struct {
	UserName string
	Password string
	DBName   string
	Host     string
	Port     string
}

func InitConfig() {
	Env = &AppConfig{
		UserName: "root",
		Password: "0000",
		DBName:   "ordering_db",
		Host:     "169.254.52.127",
		Port:     "8080",
	}
}
