package main

import (
	"ordering-system-backend/config"
	"ordering-system-backend/database"
	"ordering-system-backend/repository"
	"ordering-system-backend/routes"
	"ordering-system-backend/service"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitConfig()

	db := database.Connect()
	repos := repository.NewRepositories(db)
	services := service.NewServices(service.Deps{
		Repos: repos,
	})

	router := gin.Default()
	routes.SetUpRoutes(router, services)
	router.Run("localhost:" + config.Env.Port)
}
