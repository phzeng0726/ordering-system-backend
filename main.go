package main

import (
	"ordering-system-backend/db"
	"ordering-system-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()

	router := gin.Default()
	routes.SetUpRoutes(router)

	port := "8080"
	router.Run("localhost:" + port)

}
