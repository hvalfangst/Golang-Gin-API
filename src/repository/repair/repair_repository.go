package repair

import (
	"context"
	"github.com/uptrace/bun"
	"hvalfangst/golang-gin-api-with-bun/src/model"
	"log"
)

func CreateRepair(db *bun.DB, repair *model.Repair) error {
	_, err := db.NewInsert().Model(repair).Exec(context.Background())
	if err != nil {
		log.Printf("Error creating repair: %v", err)
		return err
	}
	return nil
}

func GetRepairByID(db *bun.DB, repairID int64) (*model.Repair, error) {
	repair := new(model.Repair)
	err := db.NewSelect().Model(repair).Where("id = ?", repairID).Scan(context.Background())
	if err != nil {
		log.Printf("Error retrieving repair by ID: %v", err)
		return nil, err
	}
	return repair, nil
}

func UpdateRepair(db *bun.DB, repairID int64, repair *model.Repair) error {
	_, err := db.NewUpdate().Model(repair).Where("id = ?", repairID).Exec(context.Background())
	if err != nil {
		log.Printf("Error updating repair: %v", err)
		return err
	}
	return nil
}

func DeleteRepairByID(db *bun.DB, repairID int64) error {
	_, err := db.NewDelete().Model(&model.Repair{ID: repairID}).WherePK().Exec(context.Background())
	if err != nil {
		log.Printf("Error deleting repair by ID: %v", err)
		return err
	}
	return nil
}

func GetRepairsByCarID(db *bun.DB, carID int64) ([]*model.Repair, error) {
	var repairs []*model.Repair
	err := db.NewSelect().Model(&repairs).Where("car_id = ?", carID).Scan(context.Background())
	if err != nil {
		log.Printf("Error retrieving repairs by car ID: %v", err)
		return nil, err
	}
	return repairs, nil
}
