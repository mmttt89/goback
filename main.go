package main

import (
	config "goback/src/constants"
	"goback/src/database"
	"goback/src/model"
	route "goback/src/routes"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	db := database.ConnectDB()
	migrateErr := db.AutoMigrate(&model.User{})
	if migrateErr != nil {
		// Handle error
		log.Println("there is migration error")

	}
	r := gin.Default()
	route.SetupRoutes(r, db)

	log.Println("before running the app")
	r.Run(config.HostAddress)
	log.Println("running on port 8080")
}
