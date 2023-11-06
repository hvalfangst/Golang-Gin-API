package engine

import (
	"github.com/go-pg/pg/v10"
	"hvalfangst/imperative-golang-gin-api/src/model"
	"log"
)

func CreateEngine(db *pg.DB, engine *model.Engine) error {
	_, err := db.Model(engine).Insert()
	if err != nil {
		log.Printf("Error creating engine: %v", err)
		return err
	}
	return nil
}

func GetEngineByID(db *pg.DB, engineID int64) (*model.Engine, error) {
	engine := &model.Engine{}
	err := db.Model(engine).Where("id = ?", engineID).Select()
	if err != nil {
		log.Printf("Error retrieving engine by ID: %v", err)
		return nil, err
	}
	return engine, nil
}

func UpdateEngine(db *pg.DB, engineID int64, engine *model.Engine) error {
	_, err := db.Model(engine).Where("id = ?", engineID).Update()
	if err != nil {
		log.Printf("Error updating engine: %v", err)
		return err
	}
	return nil
}

func DeleteEngineByID(db *pg.DB, engineID int64) error {
	engine := &model.Engine{ID: engineID}

	_, err := db.Model(engine).WherePK().Delete()
	if err != nil {
		log.Printf("Error deleting engine by ID: %v", err)
		return err
	}
	return nil
}

func GetEngineByCarID(db *pg.DB, carID int64) (*model.Engine, error) {
	engine := &model.Engine{}
	err := db.Model(engine).Where("car_id = ?", carID).Select()
	if err != nil {
		log.Printf("Error retrieving engines by car ID: %v", err)
		return nil, err
	}
	return engine, nil
}
