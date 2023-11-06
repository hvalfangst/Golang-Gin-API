package model

import "time"

type Tire struct {
	ID                int64     `json:"id"`
	CarID             int64     `json:"car_id"`
	Brand             string    `json:"brand"`
	Size              string    `json:"size"`
	ManufacturingDate time.Time `json:"manufacturing_date"`
	_                 struct{}  `pg:"_schema:tires"`
}
