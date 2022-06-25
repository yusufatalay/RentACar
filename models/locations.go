package models

type Location struct {
	ID               uint              `gorm:"primaryKey" json:"id"`
	Name             string            `json:"name"`
	Active           bool              `json:"active"`
	Offices          []Office          `gorm:"foreignkey:LocationID" json:"offices"`
	CarsReservations []CarsReservation `gorm:"foreignkey:LocationID" json:"cars_reservations"`
	CreatedAt        int64             `gorm:"autoCreateTime"`
	UpdatedAt        int64             `gorm:"autoUpdateTime:milli"`
}
