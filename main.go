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

	conn := database.Connect()
	// db, err := conn.DB()
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer db.Close()

	repos := repository.NewRepositories(conn)
	services := service.NewServices(service.Deps{
		Repos: repos,
	})

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.Use(routes.Middleware(conn))
	routesSetup := routes.NewHandler(router, services)
	routes.SetUpRoutes(routesSetup)
	router.Run("localhost:" + config.Env.Port)
}
