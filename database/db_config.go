package database

import (
	"fmt"
	"log"

	"github.com/Didik2584/task-5-pbi-btpns-Didik_kurniawan/helpers"
	"github.com/Didik2584/task-5-pbi-btpns-Didik_kurniawan/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

// InitDB initializes the database connection and assigns it to the global db variable.
func InitDB() *gorm.DB {
	// Get the current environment stage (development, production, or testing)
	stage := helpers.GetAsString("STAGE", "development")
	var path string

	// Load environment variables based on the stage
	if stage == "testing" {
		path = "../.env"
	} else {
		path = ".env"
	}

	// Load the environment variables from the specified path
	helpers.LoadEnv(path)

	// Construct the database URI from environment variables
	dbURI := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		helpers.GetAsString("DB_USER", "postgres"),
		helpers.GetAsString("DB_PASSWORD", "postgres"),
		helpers.GetAsString("DB_HOST", "localhost"),
		helpers.GetAsInt("DB_PORT", 5432),
		helpers.GetAsString("DB_NAME", "postgres"),
	)

	// Open a connection to the database
	var err error
	db, err = gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		// Log the error and exit if the connection fails
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// Return the initialized db instance
	return db
}

// MigrateDB runs the migrations for the models (User, Photo) based on the current environment stage.
func MigrateDB(db *gorm.DB) {
	// Get the current environment stage
	stage := helpers.GetAsString("STAGE", "development")

	// Only run migrations in development or production environments
	if stage == "development" || stage == "production" {
		// Run database migrations for the User and Photo models
		db.Debug().AutoMigrate(&models.User{}, &models.Photo{})
	}
}

// GetDB returns the initialized global db instance.
func GetDB() *gorm.DB {
	return db
}
