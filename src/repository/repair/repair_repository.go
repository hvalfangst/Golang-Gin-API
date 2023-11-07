package repair

import (
	"context"
	"github.com/uptrace/bun"
	"hvalfangst/golang-gin-api-with-bun/src/model"
	"log"
	"time"
)

func GetRepairsByCarID(db *bun.DB, carID int64) ([]*model.Repair, error) {
	// Initialize a slice for repairs.
	var repairs []*model.Repair

	// Create a context with timeout or cancellation if needed.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Execute the query.
	err := db.NewSelect().Model(&repairs).Where("car_id = ?", carID).Scan(ctx)

	if err != nil {
		log.Printf("Error retrieving repairs by car ID: %v", err)
		return nil, err
	}

	// If no error occurred, return the repairs.
	return repairs, nil
}
