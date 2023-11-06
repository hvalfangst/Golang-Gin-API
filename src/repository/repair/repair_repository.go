package repair

import (
	"github.com/go-pg/pg/v10"
	"hvalfangst/imperative-golang-gin-api/src/model" // Import the model package with the correct path
	"log"
)

func CreateRepair(db *pg.DB, repair *model.Repair) error {
	_, err := db.Model(repair).Insert()
	if err != nil {
		log.Printf("Error creating repair: %v", err)
		return err
	}
	return nil
}

func GetRepairByID(db *pg.DB, repairID int64) (*model.Repair, error) {
	repair := &model.Repair{}
	err := db.Model(repair).Where("id = ?", repairID).Select()
	if err != nil {
		log.Printf("Error retrieving repair by ID: %v", err)
		return nil, err
	}
	return repair, nil
}

func UpdateRepair(db *pg.DB, repairID int64, repair *model.Repair) error {
	_, err := db.Model(repair).Where("id = ?", repairID).Update()
	if err != nil {
		log.Printf("Error updating repair: %v", err)
		return err
	}
	return nil
}

func DeleteRepairByID(db *pg.DB, repairID int64) error {
	repair := &model.Repair{ID: repairID}

	_, err := db.Model(repair).WherePK().Delete()
	if err != nil {
		log.Printf("Error deleting repair by ID: %v", err)
		return err
	}
	return nil
}

func GetRepairsByCarID(db *pg.DB, carID int64) ([]*model.Repair, error) {
	var repairs []*model.Repair
	err := db.Model(&repairs).Where("car_id = ?", carID).Select()
	if err != nil {
		log.Printf("Error retrieving repairs by car ID: %v", err)
		return nil, err
	}
	return repairs, nil
}
