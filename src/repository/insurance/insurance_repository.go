package insurance

import (
	"context"
	"github.com/uptrace/bun"
	"hvalfangst/golang-gin-api-with-bun/src/model"
	"log"
	"time"
)

func GetInsurancesByCarID(db *bun.DB, carID int64) ([]*model.Insurance, error) {
	// Initialize a slice for insurances.
	var insurances []*model.Insurance

	// Create a context with timeout or cancellation if needed.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Execute the query.
	err := db.NewSelect().Model(&insurances).Where("car_id = ?", carID).Scan(ctx)

	if err != nil {
		log.Printf("Error retrieving insurances by car ID: %v", err)
		return nil, err
	}

	// If no error occurred, return the insurances.
	return insurances, nil
}
