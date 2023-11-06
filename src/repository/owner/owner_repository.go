package owner

import (
	"github.com/go-pg/pg/v10"
	"hvalfangst/imperative-golang-gin-api/src/model" // Import the model package with the correct path
	"log"
)

func CreateOwner(db *pg.DB, owner *model.Owner) error {
	_, err := db.Model(owner).Insert()
	if err != nil {
		log.Printf("Error creating owner: %v", err)
		return err
	}
	return nil
}

func GetOwnerByID(db *pg.DB, ownerID int64) (*model.Owner, error) {
	owner := &model.Owner{}
	err := db.Model(owner).Where("id = ?", ownerID).Select()
	if err != nil {
		log.Printf("Error retrieving owner by ID: %v", err)
		return nil, err
	}
	return owner, nil
}

func UpdateOwner(db *pg.DB, ownerID int64, owner *model.Owner) error {
	_, err := db.Model(owner).Where("id = ?", ownerID).Update()
	if err != nil {
		log.Printf("Error updating owner: %v", err)
		return err
	}
	return nil
}

func DeleteOwnerByID(db *pg.DB, ownerID int64) error {
	owner := &model.Owner{ID: ownerID}

	_, err := db.Model(owner).WherePK().Delete()
	if err != nil {
		log.Printf("Error deleting owner by ID: %v", err)
		return err
	}
	return nil
}
