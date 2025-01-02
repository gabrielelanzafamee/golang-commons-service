package main

import (
	"log"

	"github.com/gabrielelanzafamee/golang-commons-service/db"
	"github.com/gabrielelanzafamee/golang-commons-service/utils"
)

func main() {
	// Initialize Logger
	utils.InitLogger()
	utils.Logger.Info("Starting Common Service")

	// Load Configurations
	utils.LoadConfig()
	utils.Logger.Info("Configurations loaded")

	// Run Database Migrations
	db.RunMigrations()
	utils.Logger.Info("Database migrations completed")

	log.Println("Common Service is up and running!")
}
