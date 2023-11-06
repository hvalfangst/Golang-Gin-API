package db

import (
	"github.com/go-pg/pg/v10"
	"hvalfangst/imperative-golang-gin-api/src/configuration"
	"log"
)

func ConnectDatabase(config configuration.Db) *pg.DB {
	opts := &pg.Options{
		User:     config.User,
		Password: config.Password,
		Addr:     config.Address,
		Database: config.Database,
	}
	return pg.Connect(opts)
}

func CloseDatabase(db *pg.DB) {
	if db == nil {
		return
	}

	err := db.Close()
	if err != nil {
		log.Printf("Error closing database connection: %v", err)
	}
}
