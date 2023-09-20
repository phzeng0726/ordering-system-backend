package main

import (
	"ordering-system-backend/config"
	"ordering-system-backend/database"
	delivery "ordering-system-backend/delivery/http"
	"ordering-system-backend/repository"
	"ordering-system-backend/service"
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
