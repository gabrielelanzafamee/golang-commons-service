package db

import (
	"log"

	"gorm.io/gorm"
)

// RunMigrations runs database migrations
func RunMigration(db *gorm.DB, model interface{}) {
	err := db.AutoMigrate(model)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Database migrations completed")
}
