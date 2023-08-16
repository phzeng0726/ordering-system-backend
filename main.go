package main

import (
	"ordering-system-backend/config"
	"ordering-system-backend/db"
	"ordering-system-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitConfig()

	db.Connect()

	router := gin.Default()
	routes.SetUpRoutes(router)

	router.Run("localhost:" + config.Env.Port)

}
