package models

type OfficesWorkingDay struct {
	ID       uint  `gorm:"primaryKey" json:"id"`
	OfficeID uint  `json:"office_id"`
	Day      uint8 `json:"day"`
}
