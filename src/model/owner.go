package model

import "time"

type Owner struct {
	ID          int64     `json:"id"`
	OwnerName   string    `json:"owner_name"`
	ContactInfo string    `json:"contact_info"`
	DateOfBirth time.Time `json:"date_of_birth"`
	_           struct{}  `pg:"_schema:owners"`
}
