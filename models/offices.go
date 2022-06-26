package models

import (
	"errors"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/yusufatalay/RentACar/database"
	"gorm.io/gorm"
)

type Office struct {
	ID                uint                `gorm:"primaryKey" json:"id"`
	Vendor            string              `json:"vendor" validate:"required,string,min=2,max=32"`
	LocationID        uint                `json:"location_id" validate:"required"`
	OpeningHour       time.Time           `json:"opening_hour" validate:"required,datetime"`
	ClosingHour       time.Time           `json:"closing_hour" validate:"required,datetime"`
	Cars              []Car               `gorm:"foreignkey:OfficeID" json:"cars"`
	OfficesWorkingDay []OfficesWorkingDay `gorm:"foreignkey:OfficeID" json:"offices_working_day"`

	CreatedAt int64 `gorm:"autoCreateTime"`
	UpdatedAt int64 `gorm:"autoUpdateTime:milli"`
}

func ValidateOffice(office *Office) []validator.FieldError {
	var officeErrors []validator.FieldError
	err := validate.Struct(office)
	if err != nil {
		log.Printf("Error: %v", err)
		for _, err := range err.(validator.ValidationErrors) {
			officeErrors = append(officeErrors, err)
		}
	}
	return officeErrors
}

func (office *Office) BeforeCreate(tx *gorm.DB) (err error) {
	errs := ValidateOffice(office)
	if len(errs) > 0 {
		return errors.New("could not save not valid office")
	}
	return
}

func CreateOffice(office *Office) error {
	err := database.DBConn.Create(&office).Error
	if err != nil {
		log.Printf("Error: %v", err)
		return err
	}
	return nil
}

func GetOffice(id uint) (*Office, error) {
	office := &Office{}
	err := database.DBConn.First(&office, id).Error
	if err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}
	return office, nil
}
func GetAllOffices() ([]Office, error) {
	offices := []Office{}
	err := database.DBConn.Find(&offices).Error
	if err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}
	return offices, nil
}

func UpdateOffice(office *Office) error {
	err := database.DBConn.Save(&office).Error
	if err != nil {
		log.Printf("Error: %v", err)
		return err
	}
	return nil
}

func IsOfficeOpen(id uint, reservationBegin time.Time) (bool, error) {
	office, err := GetOffice(id)
	reservationDay := reservationBegin.Weekday().String()
	var isOpen bool
	if err != nil {
		log.Printf("Error: %v", err)
		return false, err
	}

	// check if today is working day
	for _, day := range office.OfficesWorkingDay {
		if day.Day == GetDayIndex(reservationDay) {
			isOpen = true
			break
		}
	}
	if office.OpeningHour.Before(reservationBegin) && office.ClosingHour.After(reservationBegin) {
		isOpen = true
	}
	return isOpen, nil
}
