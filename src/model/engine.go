package model

import "time"

type Engine struct {
	ID                int64     `json:"id"`
	CarID             int64     `json:"car_id"`
	Type              string    `json:"type"`
	Displacement      string    `json:"displacement"`
	Horsepower        int       `json:"horsepower"`
	ManufacturingDate time.Time `json:"manufacturing_date"`
	_                 struct{}  `pg:"_schema:engines"`
}
