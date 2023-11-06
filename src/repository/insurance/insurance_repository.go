package insurance

import (
	"github.com/go-pg/pg/v10"
	"hvalfangst/imperative-golang-gin-api/src/model"
	"log"
)

func CreateInsurance(db *pg.DB, insurance *model.Insurance) error {
	_, err := db.Model(insurance).Insert()
	if err != nil {
		log.Printf("Error creating insurance: %v", err)
		return err
	}
	return nil
}

func GetInsuranceByID(db *pg.DB, insuranceID int64) (*model.Insurance, error) {
	insurance := &model.Insurance{}
	err := db.Model(insurance).Where("id = ?", insuranceID).Select()
	if err != nil {
		log.Printf("Error retrieving insurance by ID: %v", err)
		return nil, err
	}
	return insurance, nil
}

func GetInsurancesByCarID(db *pg.DB, carID int64) ([]*model.Insurance, error) {
	var insurances []*model.Insurance
	err := db.Model(&insurances).Where("car_id = ?", carID).Select()
	if err != nil {
		log.Printf("Error retrieving insurances by car ID: %v", err)
		return nil, err
	}
	return insurances, nil
}
