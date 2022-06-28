package models

import (
	"errors"
	"log"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/yusufatalay/RentACar/database"
	"gorm.io/gorm"
)

type Location struct {
	ID               uint              `gorm:"primaryKey;auto_increment" json:"id"`
	Name             string            `json:"name" validate:"required,min=2,max=32"`
	Active           string            `json:"active" validate:"required,boolean"`
	Offices          []Office          `gorm:"foreignKey:LocationID" json:"offices"`
	CarsReservations []CarsReservation `gorm:"foreignKey:LocationID" json:"cars_reservations"`
}

type SuccessfullActiveLocations struct {
	Message string
	Data    []Location
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

func GetAllLocations() ([]Location, error) {
	var locations []Location
	err := database.DBConn.Find(&locations).Error
	if err != nil {
		return nil, err
	}
	return locations, nil
}

func GetActiveLocations() ([]Location, error) {
	var locations []Location
	err := database.DBConn.Where("active = ?", "true").Find(&locations).Error
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
	result, _ := strconv.ParseBool(location.Active)
	return result, nil
}
