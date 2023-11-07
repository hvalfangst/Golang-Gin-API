package car

import (
	"context"
	"github.com/uptrace/bun"
	"hvalfangst/golang-gin-api-with-bun/src/model"
	"log"
	"time"
)

func GetCarByID(db *bun.DB, carID int64) (*model.Car, error) {
	// Initialize a car instance.
	car := &model.Car{}

	// Create a context with timeout or cancellation if needed.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Execute the query.
	err := db.NewSelect().Model(car).Where("id = ?", carID).Scan(ctx)

	if err != nil {
		log.Printf("Error retrieving car by ID: %v", err)
		return nil, err
	}

	// If no error occurred, return the car.
	return car, nil
}
