package models

import (
	"errors"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/yusufatalay/RentACar/database"
	"gorm.io/gorm"
)

type Location struct {
	ID               uint              `gorm:"primaryKey" json:"id"`
	Name             string            `json:"name" validate:"required,string,min=2,max=32"`
	Active           bool              `json:"active" validate:"required,boolean"`
	Offices          []Office          `gorm:"foreignkey:LocationID" json:"offices"`
	CarsReservations []CarsReservation `gorm:"foreignkey:LocationID" json:"cars_reservations"`
}

type LocationIdentifier struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

type SuccessfullActiveLocations struct {
	Message string
	Data    []LocationIdentifier
}

func ValidateLocation(location *Location) []validator.FieldError {
	var locationErrors []validator.FieldError
	err := validate.Struct(location)
	if err != nil {
		log.Printf("Error: %v", err)
		for _, err := range err.(validator.ValidationErrors) {
			locationErrors = append(locationErrors, err)
		}
	}
	return locationErrors
}

func (location *Location) BeforeCreate(tx *gorm.DB) (err error) {
	errs := ValidateLocation(location)
	if len(errs) > 0 {
		return errors.New("could not save not valid location")
	}
	return
}

func CreateLocation(location *Location) error {
	err := database.DBConn.Create(&location).Error
	if err != nil {
		log.Printf("Error: %v", err)
		return err
	}
	return nil
}

func UpdateLocation(location *Location) error {
	err := database.DBConn.Save(&location).Error
	if err != nil {
		log.Printf("Error: %v", err)
		return err
	}
	return nil
}

func GetLocation(id uint) (*Location, error) {
	location := &Location{}
	err := database.DBConn.First(&location, id).Error
	if err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}
	return location, nil
}

func GetAllLocations() ([]LocationIdentifier, error) {
	var locations []LocationIdentifier
	err := database.DBConn.Find(&locations).Error
	if err != nil {
		return nil, err
	}
	return locations, nil
}

func GetActiveLocations() ([]LocationIdentifier, error) {
	var locations []LocationIdentifier
	err := database.DBConn.Where("active = ?", true).Find(&locations).Error
	if err != nil {
		return nil, err
	}
	return locations, nil
}

func IsLocationActive(id uint) (bool, error) {
	var location Location
	err := database.DBConn.First(&location, id).Error
	if err != nil {
		return false, err
	}
	return location.Active, nil
}
