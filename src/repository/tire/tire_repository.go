package tire

import (
	"context"
	"github.com/uptrace/bun"
	"hvalfangst/golang-gin-api-with-bun/src/model"
	"log"
	"time"
)

func GetTiresByCarID(db *bun.DB, carID int64) ([]*model.Tire, error) {
	// Initialize a slice for tires.
	var tires []*model.Tire

	// Create a context with timeout or cancellation if needed.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Execute the query.
	err := db.NewSelect().Model(&tires).Where("car_id = ?", carID).Scan(ctx)

	if err != nil {
		log.Printf("Error retrieving tires by car ID: %v", err)
		return nil, err
	}

	// If no error occurred, return the tires.
	return tires, nil
}
