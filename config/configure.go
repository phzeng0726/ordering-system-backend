package config

var (
	Env *AppConfig
)

type AppConfig struct {
	UserName string
	Password string
	DBName   string
	Port     string
}

func InitConfig() {
	Env = &AppConfig{
		UserName: "root",
		Password: "0000",
		DBName:   "ordering_db",
		Port:     "8080",
	}
}
