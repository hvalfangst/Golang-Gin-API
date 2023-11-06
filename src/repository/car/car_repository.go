package car

import (
	"github.com/go-pg/pg/v10"
	"hvalfangst/imperative-golang-gin-api/src/model"
	"log"
)

func CreateCar(db *pg.DB, car *model.Car) error {
	_, err := db.Model(car).Insert()
	if err != nil {
		log.Printf("Error creating car: %v", err)
		return err
	}
	return nil
}

func GetCarByID(db *pg.DB, carID int64) (*model.Car, error) {
	car := &model.Car{}
	err := db.Model(car).Where("id = ?", carID).Select()
	if err != nil {
		log.Printf("Error retrieving car by ID: %v", err)
		return nil, err
	}
	return car, nil
}

func UpdateCar(db *pg.DB, carID int64, car *model.Car) error {
	_, err := db.Model(car).Where("id = ?", carID).Update()
	if err != nil {
		log.Printf("Error updating car: %v", err)
		return err
	}
	return nil
}

func DeleteCarByID(db *pg.DB, carID int64) error {
	car := &model.Car{ID: carID}

	_, err := db.Model(car).WherePK().Delete()
	if err != nil {
		log.Printf("Error deleting car by ID: %v", err)
		return err
	}
	return nil
}

func GetCarsByOwnerID(db *pg.DB, ownerID int64) ([]*model.Car, error) {
	var cars []*model.Car
	err := db.Model(&cars).Where("owner_id = ?", ownerID).Select()
	if err != nil {
		log.Printf("Error retrieving cars by owner ID: %v", err)
		return nil, err
	}
	return cars, nil
}
