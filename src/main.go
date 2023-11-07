package main

import (
	"github.com/gin-gonic/gin"
	"hvalfangst/golang-gin-api-with-bun/src/configuration"
	"hvalfangst/golang-gin-api-with-bun/src/db"
	CarsRoute "hvalfangst/golang-gin-api-with-bun/src/route"
	"log"
)

func main() {
	r := gin.Default()

	// Fetch JSON based on key 'db' for file 'configuration.json' to be used as Db
	conf, err := configuration.Get()
	if err != nil {
		log.Fatalf("Error reading configuration file: %v", err)
	}

	// Connect to the database based on Configuration derived from 'configuration.json'
	database, _ := db.CreateDB(conf.(configuration.Db))

	CarsRoute.ConfigureRoute(r, database)

	// Run the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
