package tire

import (
	"context"
	"github.com/uptrace/bun"
	"hvalfangst/golang-gin-api-with-bun/src/model"
	"log"
)

func CreateTire(db *bun.DB, tire *model.Tire) error {
	_, err := db.NewInsert().Model(tire).Exec(context.Background())
	if err != nil {
		log.Printf("Error creating tire: %v", err)
		return err
	}
	return nil
}

func GetTireByID(db *bun.DB, tireID int64) (*model.Tire, error) {
	tire := new(model.Tire)
	err := db.NewSelect().Model(tire).Where("id = ?", tireID).Scan(context.Background())
	if err != nil {
		log.Printf("Error retrieving tire by ID: %v", err)
		return nil, err
	}
	return tire, nil
}

func UpdateTire(db *bun.DB, tireID int64, tire *model.Tire) error {
	_, err := db.NewUpdate().Model(tire).Where("id = ?", tireID).Exec(context.Background())
	if err != nil {
		log.Printf("Error updating tire: %v", err)
		return err
	}
	return nil
}

func DeleteTireByID(db *bun.DB, tireID int64) error {
	_, err := db.NewDelete().Model(&model.Tire{ID: tireID}).WherePK().Exec(context.Background())
	if err != nil {
		log.Printf("Error deleting tire by ID: %v", err)
		return err
	}
	return nil
}

func GetTiresByCarID(db *bun.DB, carID int64) ([]*model.Tire, error) {
	var tires []*model.Tire
	err := db.NewSelect().Model(&tires).Where("car_id = ?", carID).Scan(context.Background())
	if err != nil {
		log.Printf("Error retrieving tires by car ID: %v", err)
		return nil, err
	}
	return tires, nil
}
