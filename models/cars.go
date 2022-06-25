package models

type Car struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	Fuel         string `json:"fuel"`
	Transmission string `json:"transmission"`
	Name         string `json:"name"`
	OfficeID     uint   `json:"office_id"`
	CreatedAt    int64  `gorm:"autoCreateTime"`
	UpdatedAt    int64  `gorm:"autoUpdateTime:milli"`
}
