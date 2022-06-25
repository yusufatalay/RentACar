package models

import "time"

type CarsReservation struct {
	ID               uint `gorm:"primaryKey" json:"id"`
	CarID            uint `json:"car_id"`
	LeaserID       uint `json:"leaser_id"`
	LocationID       uint      `json:"location_id"`
	ReservationBegin time.Time `json:"reservation_begin"`
	ReservationEnd   time.Time `json:"reservation_end"`
}
