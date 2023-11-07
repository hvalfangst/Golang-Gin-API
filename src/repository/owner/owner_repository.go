package owner

import (
	"context"
	"github.com/uptrace/bun"
	"hvalfangst/golang-gin-api-with-bun/src/model"
	"log"
	"time"
)

func GetOwnerByID(db *bun.DB, ownerID int64) (*model.Owner, error) {
	// Initialize an owner instance.
	owner := &model.Owner{}

	// Create a context with timeout or cancellation if needed.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Execute the query.
	err := db.NewSelect().Model(owner).Where("id = ?", ownerID).Scan(ctx)

	if err != nil {
		log.Printf("Error retrieving owner by ID: %v", err)
		return nil, err
	}

	// If no error occurred, return the owner.
	return owner, nil
}
