package models

import (
	"errors"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/yusufatalay/RentACar/database"
	"gorm.io/gorm"
)

type Car struct {
	ID           uint   `gorm:"primaryKey;auto_increment"  json:"id"`
	Vendor       string `json:"vendor" validate:"required,min=3,max=32"`
	Fuel         string `json:"fuel" validate:"required,min=2,max=32"`
	Transmission string `json:"transmission" validate:"required,min=2,max=32"`
	Name         string `json:"name" validate:"required,min=2,max=32"`
	OfficeID     uint   `json:"office_id" validate:"required"`
}

func ValidateCar(car *Car) []validator.FieldError {
	var carErrors []validator.FieldError

	err := validate.Struct(car)
	if err != nil {
		log.Printf("Error: %v", err)
		for _, err := range err.(validator.ValidationErrors) {
			carErrors = append(carErrors, err)
		}
	}
	return carErrors
}

func (car *Car) BeforeCreate(tx *gorm.DB) (err error) {
	errs := ValidateCar(car)

	if len(errs) > 0 {
		return errors.New("could not save not valid car")
	}
	return
}

func GetCar(id uint) (*Car, error) {
	car := &Car{}
	err := database.DBConn.First(&car, id).Error
	if err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}
	return car, nil
}

func GetAllCars() ([]Car, error) {
	var cars []Car
	err := database.DBConn.Find(&cars).Error
	if err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}
	return cars, nil
}

func SaveCar(car *Car) error {
	err := database.DBConn.Save(&car).Error
	if err != nil {
		log.Printf("Error: %v", err)
		return err
	}
	return nil
}

type CarAvailabilityIdentifier struct {
	LocationID       uint      `json:"location_id"`
	ReservationBegin time.Time `json:"reservation_begin"`
	ReservationEnd   time.Time `json:"reservation_end"`
}

type SuccessfullAvailableCars struct {
	Message string
	Data    []Car
}

// GetCarAvailability returns a list of cars that are available for the given time and  location
func GetAvailableCars(carAvailabilityIdentifier *CarAvailabilityIdentifier) ([]Car, error) {
	var cars []Car

	// from car we need to go to the office and from office we need to go to location to see if its available
	// we should check if reservation dates holds, location is active and office is in its working day or not

	err := database.DBConn.Joins("JOIN offices ON offices.id = cars.office_id").Joins("JOIN locations ON locations.id = offices.location_id").Joins("JOIN offices_working_days ON offices_working_days.office_id = offices.id ").Where("offices_working_days.day = ? locations.id = ? AND locations.active = ?  AND cars.id NOT IN (SELECT car_id FROM reservations WHERE (reservation_begin BETWEEN ? AND ?) OR (reservation_end BETWEEN ? AND ?))",
		carAvailabilityIdentifier.ReservationBegin.Weekday(),
		carAvailabilityIdentifier.LocationID,
		"true",
		carAvailabilityIdentifier.ReservationBegin,
		carAvailabilityIdentifier.ReservationEnd,
		carAvailabilityIdentifier.ReservationBegin,
		carAvailabilityIdentifier.ReservationEnd).Find(&cars).Error

	if err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}

	return cars, nil
}
