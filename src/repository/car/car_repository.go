package car

import (
	"context"
	"github.com/uptrace/bun"
	"hvalfangst/golang-gin-api-with-bun/src/model"
	"log"
)

func CreateCar(db *bun.DB, car *model.Car) error {
	_, err := db.NewInsert().Model(car).Exec(context.Background())
	if err != nil {
		log.Printf("Error creating car: %v", err)
		return err
	}
	return nil
}

func GetCarByID(db *bun.DB, carID int64) (*model.Car, error) {
	car := new(model.Car)
	err := db.NewSelect().Model(car).Where("id = ?", carID).Scan(context.Background())
	if err != nil {
		log.Printf("Error retrieving car by ID: %v", err)
		return nil, err
	}
	return car, nil
}

func UpdateCar(db *bun.DB, carID int64, car *model.Car) error {
	_, err := db.NewUpdate().Model(car).Where("id = ?", carID).Exec(context.Background())
	if err != nil {
		log.Printf("Error updating car: %v", err)
		return err
	}
	return nil
}

func DeleteCarByID(db *bun.DB, carID int64) error {
	_, err := db.NewDelete().Model(&model.Car{ID: carID}).WherePK().Exec(context.Background())
	if err != nil {
		log.Printf("Error deleting car by ID: %v", err)
		return err
	}
	return nil
}

func GetCarsByOwnerID(db *bun.DB, ownerID int64) ([]*model.Car, error) {
	var cars []*model.Car
	err := db.NewSelect().Model(&cars).Where("owner_id = ?", ownerID).Scan(context.Background())
	if err != nil {
		log.Printf("Error retrieving cars by owner ID: %v", err)
		return nil, err
	}
	return cars, nil
}
