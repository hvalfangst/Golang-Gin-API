package service

import (
	"github.com/uptrace/bun"
	"hvalfangst/golang-gin-api-with-bun/src/model"
	CarRepository "hvalfangst/golang-gin-api-with-bun/src/repository/car"
	EngineRepository "hvalfangst/golang-gin-api-with-bun/src/repository/engine"
	InsuranceRepository "hvalfangst/golang-gin-api-with-bun/src/repository/insurance"
	OwnerRepository "hvalfangst/golang-gin-api-with-bun/src/repository/owner"
	RepairRepository "hvalfangst/golang-gin-api-with-bun/src/repository/repair"
	TireRepository "hvalfangst/golang-gin-api-with-bun/src/repository/tire"
)

func GetCarById(db *bun.DB, ID int64) (*model.Car, error) {
	car, err := CarRepository.GetCarByID(db, ID)
	if err != nil {
		return nil, err
	}
	return car, nil
}

func GetCarDetails(db *bun.DB, carID int64) (*model.CarDetails, error) {
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
