package main

import (
	"log"
	"ordering-system-backend/config"
	"ordering-system-backend/database"
	"ordering-system-backend/domain"

	"gorm.io/gorm"
)

func fetchData(db *gorm.DB) {
	var menu domain.Menu
	if err := db.Preload("MenuItemMapping").First(&menu, 1).Error; err != nil {
		log.Println("failed to load menu with items")
	}
	// menuItem := domain.MenuItem{
	// 	Id:             1,
	// 	MenuCategoryId: 1,
	// }
	// if err := db.Preload("MenuCategory").First(&menuItem).Error; err != nil {
	// 	log.Println("failed to load menuItem")
	// }
	log.Println(menu)
}

func main() {
	config.InitConfig()

	db := database.Connect()
	fetchData(db)
	// 	repos := repository.NewRepositories(db)
	// 	services := service.NewServices(service.Deps{
	// 		Repos: repos,
	// 	})

	// 	router := gin.Default()
	// 	routesSetup := routes.NewRoutesSetup(router, services)
	// 	routes.SetUpRoutes(routesSetup)
	// 	router.Run("localhost:" + config.Env.Port)
}
