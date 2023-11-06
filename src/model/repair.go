package model

import "time"

type Repair struct {
	ID         int64     `json:"id"`
	CarID      int64     `json:"car_id"`
	RepairType string    `json:"repair_type"`
	RepairDate time.Time `json:"repair_date"`
	Cost       float64   `json:"cost"`
	_          struct{}  `pg:"_schema:repairs"`
}
