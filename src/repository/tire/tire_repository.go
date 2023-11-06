package tire

import (
	"github.com/go-pg/pg/v10"
	"hvalfangst/imperative-golang-gin-api/src/model" // Import the model package with the correct path
	"log"
)

func CreateTire(db *pg.DB, tire *model.Tire) error {
	_, err := db.Model(tire).Insert()
	if err != nil {
		log.Printf("Error creating tire: %v", err)
		return err
	}
	return nil
}

func GetTireByID(db *pg.DB, tireID int64) (*model.Tire, error) {
	tire := &model.Tire{}
	err := db.Model(tire).Where("id = ?", tireID).Select()
	if err != nil {
		log.Printf("Error retrieving tire by ID: %v", err)
		return nil, err
	}
	return tire, nil
}

func UpdateTire(db *pg.DB, tireID int64, tire *model.Tire) error {
	_, err := db.Model(tire).Where("id = ?", tireID).Update()
	if err != nil {
		log.Printf("Error updating tire: %v", err)
		return err
	}
	return nil
}

func DeleteTireByID(db *pg.DB, tireID int64) error {
	tire := &model.Tire{ID: tireID}

	_, err := db.Model(tire).WherePK().Delete()
	if err != nil {
		log.Printf("Error deleting tire by ID: %v", err)
		return err
	}
	return nil
}

func GetTiresByCarID(db *pg.DB, carID int64) ([]*model.Tire, error) {
	var tires []*model.Tire
	err := db.Model(&tires).Where("car_id = ?", carID).Select()
	if err != nil {
		log.Printf("Error retrieving tires by car ID: %v", err)
		return nil, err
	}
	return tires, nil
}
