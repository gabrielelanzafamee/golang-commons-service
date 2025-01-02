package db

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ConnectDatabase initializes the database connection
func ConnectDatabase(db *gorm.DB, databaseName string) {
	var err error
	db, err = gorm.Open(sqlite.Open(databaseName), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Database connection established")
}

// CloseDatabase closes the database connection
func CloseDatabase(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to retrieve database instance: %v", err)
	}
	sqlDB.Close()
}
