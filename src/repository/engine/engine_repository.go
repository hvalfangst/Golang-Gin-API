package engine

import (
	"context"
	"github.com/uptrace/bun"
	"hvalfangst/golang-gin-api-with-bun/src/model"
	"log"
)

func CreateEngine(db *bun.DB, engine *model.Engine) error {
	_, err := db.NewInsert().Model(engine).Exec(context.Background())
	if err != nil {
		log.Printf("Error creating engine: %v", err)
		return err
	}
	return nil
}

func GetEngineByID(db *bun.DB, engineID int64) (*model.Engine, error) {
	engine := new(model.Engine)
	err := db.NewSelect().Model(engine).Where("id = ?", engineID).Scan(context.Background())
	if err != nil {
		log.Printf("Error retrieving engine by ID: %v", err)
		return nil, err
	}
	return engine, nil
}

func UpdateEngine(db *bun.DB, engineID int64, engine *model.Engine) error {
	_, err := db.NewUpdate().Model(engine).Where("id = ?", engineID).Exec(context.Background())
	if err != nil {
		log.Printf("Error updating engine: %v", err)
		return err
	}
	return nil
}

func DeleteEngineByID(db *bun.DB, engineID int64) error {
	_, err := db.NewDelete().Model(&model.Engine{ID: engineID}).WherePK().Exec(context.Background())
	if err != nil {
		log.Printf("Error deleting engine by ID: %v", err)
		return err
	}
	return nil
}

func GetEngineByCarID(db *bun.DB, carID int64) (*model.Engine, error) {
	engine := new(model.Engine)
	err := db.NewSelect().Model(engine).Where("car_id = ?", carID).Scan(context.Background())
	if err != nil {
		log.Printf("Error retrieving engines by car ID: %v", err)
		return nil, err
	}
	return engine, nil
}
