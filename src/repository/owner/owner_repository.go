package owner

import (
	"context"
	"github.com/uptrace/bun"
	"hvalfangst/golang-gin-api-with-bun/src/model"
	"log"
)

func CreateOwner(db *bun.DB, owner *model.Owner) error {
	_, err := db.NewInsert().Model(owner).Exec(context.Background())
	if err != nil {
		log.Printf("Error creating owner: %v", err)
		return err
	}
	return nil
}

func GetOwnerByID(db *bun.DB, ownerID int64) (*model.Owner, error) {
	owner := new(model.Owner)
	err := db.NewSelect().Model(owner).Where("id = ?", ownerID).Scan(context.Background())
	if err != nil {
		log.Printf("Error retrieving owner by ID: %v", err)
		return nil, err
	}
	return owner, nil
}

func UpdateOwner(db *bun.DB, ownerID int64, owner *model.Owner) error {
	_, err := db.NewUpdate().Model(owner).Where("id = ?", ownerID).Exec(context.Background())
	if err != nil {
		log.Printf("Error updating owner: %v", err)
		return err
	}
	return nil
}

func DeleteOwnerByID(db *bun.DB, ownerID int64) error {
	_, err := db.NewDelete().Model(&model.Owner{ID: ownerID}).WherePK().Exec(context.Background())
	if err != nil {
		log.Printf("Error deleting owner by ID: %v", err)
		return err
	}
	return nil
}
