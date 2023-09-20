package main

import (
	"ordering-system-backend/internal/config"
	"ordering-system-backend/internal/database"
	delivery "ordering-system-backend/internal/delivery/http"
	"ordering-system-backend/internal/repository"
	"ordering-system-backend/internal/service"
)

func main() {
	config.InitConfig()
	conn := database.Connect()

	repos := repository.NewRepositories(conn)
	services := service.NewServices(service.Deps{
		Repos: repos,
	})

	handlers := delivery.NewHandler(services)
	router := handlers.Init()
	if config.Env.Host == "" {
		config.Env.Host = "localhost"
	}
	router.Run(config.Env.Host + ":" + config.Env.Port)
}
