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

func InitDB() *gorm.DB {
	stage := helpers.GetAsString("STAGE", "development")
	var path string

	if stage == "testing" {
		path = "../.env"
	} else {
		path = ".env"
	}

	helpers.LoadEnv(path)

	dbURI := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		helpers.GetAsString("DB_USER", "postgres"),
		helpers.GetAsString("DB_PASSWORD", "postgres"),
		helpers.GetAsString("DB_HOST", "localhost"),
		helpers.GetAsInt("DB_PORT", 5432),
		helpers.GetAsString("DB_NAME", "postgres"),
	)

	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}

func MigrateDB(db *gorm.DB) {
	stage := helpers.GetAsString("STAGE", "development")

	if stage == "development" || stage == "production" {
		db.Debug().AutoMigrate(&models.User{}, &models.Photo{})
	}
}

func GetDB() *gorm.DB {
	return db
}
