package main

import (
	"github.com/gin-gonic/gin"
	"hvalfangst/imperative-golang-gin-api/src/configuration"
	"hvalfangst/imperative-golang-gin-api/src/db"
	CarsRoute "hvalfangst/imperative-golang-gin-api/src/route"
	"log"
)

func main() {
	r := gin.Default()

	// Fetch JSON based on key 'db' for file 'configuration.json' to be used as Db
	conf, err := configuration.Get("db")
	if err != nil {
		log.Fatalf("Error reading configuration file: %v", err)
	}

	// Connect to the database based on Configuration derived from 'configuration.json'
	database := db.ConnectDatabase(conf.(configuration.Db))
	defer db.CloseDatabase(database)

	// Serve context resources under routes '/users', '/wines', '/tokens' and '/token-activities'
	CarsRoute.ConfigureRoute(r, database)

	// Run the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
