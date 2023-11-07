package engine

import (
	"context"
	"github.com/uptrace/bun"
	"hvalfangst/golang-gin-api-with-bun/src/model"
	"log"
	"time"
)

func GetEngineByCarID(db *bun.DB, carID int64) (*model.Engine, error) {
	// Initialize an engine instance.
	engine := &model.Engine{}

	// Create a context with timeout or cancellation if needed.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Execute the query.
	err := db.NewSelect().Model(engine).Where("car_id = ?", carID).Scan(ctx)

	if err != nil {
		log.Printf("Error retrieving engine by car ID: %v", err)
		return nil, err
	}

	// If no error occurred, return the engine.
	return engine, nil
}
