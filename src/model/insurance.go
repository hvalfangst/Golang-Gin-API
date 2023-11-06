package model

import "time"

type Insurance struct {
	ID           int64     `json:"id"`
	CarID        int64     `json:"car_id"`
	PolicyNumber string    `json:"policy_number"`
	Provider     string    `json:"provider"`
	StartDate    time.Time `json:"start_date"`
	EndDate      time.Time `json:"end_date"`
	_            struct{}  `pg:"_schema:insurances"`
}
