package insurance

import (
	"context"
	"github.com/uptrace/bun"
	"hvalfangst/golang-gin-api-with-bun/src/model"
	"log"
)

func CreateInsurance(db *bun.DB, insurance *model.Insurance) error {
	_, err := db.NewInsert().Model(insurance).Exec(context.Background())
	if err != nil {
		log.Printf("Error creating insurance: %v", err)
		return err
	}
	return nil
}

func GetInsuranceByID(db *bun.DB, insuranceID int64) (*model.Insurance, error) {
	insurance := new(model.Insurance)
	err := db.NewSelect().Model(insurance).Where("id = ?", insuranceID).Scan(context.Background())
	if err != nil {
		log.Printf("Error retrieving insurance by ID: %v", err)
		return nil, err
	}
	return insurance, nil
}

func GetInsurancesByCarID(db *bun.DB, carID int64) ([]*model.Insurance, error) {
	var insurances []*model.Insurance
	err := db.NewSelect().Model(&insurances).Where("car_id = ?", carID).Scan(context.Background())
	if err != nil {
		log.Printf("Error retrieving insurances by car ID: %v", err)
		return nil, err
	}
	return insurances, nil
}
