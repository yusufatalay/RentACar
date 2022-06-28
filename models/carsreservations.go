package models

import (
	"errors"
	"log"
	"time"

	"github.com/yusufatalay/RentACar/database"
)

type CarsReservation struct {
	ID               uint      `gorm:"primaryKey;auto_increment" json:"id"`
	CarID            uint      `json:"car_id" validate:"required"`
	LeaserID         uint      `json:"leaser_id" validate:"required"`
	LocationID       uint      `json:"location_id" validate:"required"`
	ReservationBegin time.Time `json:"reservation_begin" validate:"required"`
	ReservationEnd   time.Time `json:"reservation_end"  validate:"required"`
}

//type CarReservationModel struct {
//	CarID            uint      `json:"car_id" validate:"required"`
//	LeaserID         uint      `json:"leaser_id" validate:"required"`
//	LocationID       uint      `json:"location_id" validate:"required"`
//	ReservationBegin time.Time `json:"reservation_begin" validate:"required"`
//	ReservationEnd   time.Time `json:"reservation_end"  validate:"required"`
//}

type SuccessfullAllReservations struct {
	Message string
	Data    []CarsReservation
}

type SuccessfullReservation struct {
	Message string
	Data    CarsReservation
}

// check if there already exist a reservation on given time interval
func IsReservationAvailable(reservation *CarsReservation) (bool, error) {
	var reservations []CarsReservation
	err := database.DBConn.Where("location_id = ? AND reservation_begin < ? AND reservation_end > ?", reservation.LocationID, reservation.ReservationEnd, reservation.ReservationBegin).Find(&reservations).Error
	if err != nil {
		log.Printf("Error: %v", err)
		return false, err
	}
	if len(reservations) > 0 {
		return false, nil
	}
	return true, nil
}

// we need to check the constraints again just in case
func CreateReservation(reservation *CarsReservation) error {

	var flag bool
	var err error

	flag, err = IsLocationActive(reservation.LocationID)

	if err != nil {
		log.Printf("Error: %v", err)
		return err
	}
	if !flag {
		return errors.New("location")
	}

	flag, err = IsOfficeOpen(reservation.CarID, reservation.ReservationBegin)

	if err != nil {
		log.Printf("Error: %v", err)
		return err
	}
	if !flag {
		return errors.New("office")
	}

	flag, err = IsReservationAvailable(reservation)

	if err != nil {
		log.Printf("Error: %v", err)
		return err
	}

	if !flag {
		return errors.New("no_vacancy")
	}

	err = database.DBConn.Create(&reservation).Error
	if err != nil {
		log.Printf("Error: %v", err)
		return err
	}
	return nil
}

func GetAllReservations() ([]CarsReservation, error) {
	var reservations []CarsReservation
	err := database.DBConn.Find(&reservations).Error
	if err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}
	return reservations, nil
}
