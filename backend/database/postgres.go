package database

import (
	"fitnes-tracker/config"
	"fitnes-tracker/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDatabase() {
	config.Loadenv()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.Getenv("DB_HOST"), config.Getenv("DB_USER"), config.Getenv("DB_PASSWORD"),
		config.Getenv("DB_NAME"), config.Getenv("DB_PORT"),
	)

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Optional: You can use AutoMigrate for schema changes.
	err = db.AutoMigrate(&models.User{}, &models.Workout{}, &models.Meal{}, &models.Habit{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}

// GetDB returns the database connection instance
func GetDB() *gorm.DB {
	return db
}
