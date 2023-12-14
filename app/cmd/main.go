package main

import (
	docs "ordering-system-backend/docs"
	"ordering-system-backend/internal/config"
	"ordering-system-backend/internal/database"
	delivery "ordering-system-backend/internal/delivery/http"
	"ordering-system-backend/internal/repository"
	"ordering-system-backend/internal/service"
)

// swag init -g ./cmd/main.go -o ./docs
func main() {
	config.InitConfig()
	conn := database.Connect()

	repos := repository.NewRepositories(conn)
	services := service.NewServices(service.Deps{
		Repos: repos,
	})

	handlers := delivery.NewHandler(services)
	router := handlers.Init()

	if config.Env.Port == "" {
		config.Env.Port = "8080"
	}

	docs.SwaggerInfo.BasePath = "/api/v1"

	// Host沒有填的時候就是Cloud (GCP上不需要填Host)
	router.Run(config.Env.Host + ":" + config.Env.Port)
}
