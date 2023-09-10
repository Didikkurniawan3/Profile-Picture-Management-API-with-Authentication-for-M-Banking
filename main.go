package main

import (
	"log"

	"github.com/Didik2584/task-5-pbi-btpns-Didik_kurniawan/database"
	"github.com/Didik2584/task-5-pbi-btpns-Didik_kurniawan/router"
)

func main() {

	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	err = database.MigrateDB(db)
	if err != nil {
		log.Fatalf("Failed to migrate the database: %v", err)
	}

	r := router.RouteInit(db)
	err = r.Run(":9000")
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}