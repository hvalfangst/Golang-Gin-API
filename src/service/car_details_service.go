package service

import (
	"github.com/go-pg/pg/v10"
	"hvalfangst/imperative-golang-gin-api/src/model"
	CarRepository "hvalfangst/imperative-golang-gin-api/src/repository/car"
	EngineRepository "hvalfangst/imperative-golang-gin-api/src/repository/engine"
	InsuranceRepository "hvalfangst/imperative-golang-gin-api/src/repository/insurance"
	OwnerRepository "hvalfangst/imperative-golang-gin-api/src/repository/owner"
	RepairRepository "hvalfangst/imperative-golang-gin-api/src/repository/repair"
	TireRepository "hvalfangst/imperative-golang-gin-api/src/repository/tire"
)

func GetCarDetails(db *pg.DB, carID int64) (*model.CarDetails, error) {
	car, err := CarRepository.GetCarByID(db, carID)
	if err != nil {
		return nil, err
	}

	owner, err := OwnerRepository.GetOwnerByID(db, carID)
	if err != nil {
		return nil, err
	}

	tires, err := TireRepository.GetTiresByCarID(db, carID)
	if err != nil {
		return nil, err
	}

	insurance, err := InsuranceRepository.GetInsurancesByCarID(db, carID)
	if err != nil {
		return nil, err
	}

	repairs, err := RepairRepository.GetRepairsByCarID(db, carID)
	if err != nil {
		return nil, err
	}

	engine, err := EngineRepository.GetEngineByCarID(db, carID)
	if err != nil {
		return nil, err
	}

	return &model.CarDetails{
		Car:       car,
		Owner:     owner,
		Tires:     tires,
		Insurance: insurance,
		Repairs:   repairs,
		Engine:    engine,
	}, nil
}
