package main

import (
	"log"

	"github.com/Didik2584/task-5-pbi-btpns-Didik_kurniawan/database"
	"github.com/Didik2584/task-5-pbi-btpns-Didik_kurniawan/router"
)

func main() {
	// Initialize the database connection
	db := database.InitDB()

	// Check if the DB connection was successful
	if db == nil {
		log.Fatal("Failed to connect to the database")
	}

	// Run database migrations (no return value expected)
	database.MigrateDB(db)

	// Initialize the router (adjust the arguments if needed)
	r := router.RouteInit()

	// Start the server on port 9000
	if err := r.Run(":9000"); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
