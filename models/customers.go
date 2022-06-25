package models

import (
	"time"
)

type Customer struct {
	ID              uint              `gorm:"primaryKey" json:"id"`
	FirstName       string            `json:"first_name"`
	LastName        string            `json:"last_name"`
	TCID            string            `gorm:"primaryKey" json:"tc_id"`
	DOB             time.Time         `json:"date_of_birth"`
	PhoneNumber     int64             `json:"phone_number"`
	CreatedAt       int64             `gorm:"autoCreateTime"`
	UpdatedAt       int64             `gorm:"autoUpdateTime:milli"`
}
