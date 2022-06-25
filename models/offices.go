package models

import (
	"time"
)

type Office struct {
	ID                uint                `gorm:"primaryKey" json:"id"`
	Vendor            string              `json:"vendor"`
	LocationID        uint                `json:"location_id"`
	OpeningHour       time.Time           `json:"opening_hour"`
	ClosingHour       time.Time           `json:"closing_hour"`
	Cars              []Car               `gorm:"foreignkey:OfficeID" json:"cars"`
	OfficesWorkingDay []OfficesWorkingDay `gorm:"foreignkey:OfficeID" json:"offices_working_day"`

	CreatedAt int64 `gorm:"autoCreateTime"`
	UpdatedAt int64 `gorm:"autoUpdateTime:milli"`
}
