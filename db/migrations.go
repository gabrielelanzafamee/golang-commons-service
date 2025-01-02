package db

import (
	"log"

	"gorm.io/gorm"
)

// ExampleModel represents a simple database table
type ExampleModel struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
}

// RunMigrations runs database migrations
func RunMigrations() {
	err := DB.AutoMigrate(&ExampleModel{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Database migrations completed")
}
