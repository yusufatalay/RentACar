package models

import (
	"errors"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/yusufatalay/RentACar/database"
	"gorm.io/gorm"
)

type Customer struct {
	ID          uint      `gorm:"primaryKey;auto_increment" json:"id"`
	FirstName   string    `validate:"required,min=2,max=32" json:"first_name"`
	LastName    string    `validate:"required,min=2,max=32" json:"last_name"`
	TCID        string    `gorm:"unique" validate:"required,min=11,max=11" json:"tc_id"`
	DOB         time.Time `validate:"required" json:"date_of_birth"`
	PhoneNumber string    `validate:"e164,required" json:"phone_number"` // phone numbers must be in E.164 format
}

func ValidateCustomer(customer *Customer) []validator.FieldError {
	var customerErrors []validator.FieldError

	err := validate.Struct(customer)
	if err != nil {
		log.Printf("Error: %v", err)
		for _, err := range err.(validator.ValidationErrors) {
			customerErrors = append(customerErrors, err)
		}
	}
	return customerErrors
}

//BeforeCreate hook will do the validations
func (customer *Customer) BeforeCreate(tx *gorm.DB) (err error) {
	errs := ValidateCustomer(customer)

	if len(errs) > 0 {
		return errors.New("could not save not valid customer")
	}
	return

}

func CreateCustomer(customer *Customer) error {
	err := database.DBConn.Create(&customer).Error
	if err != nil {
		log.Printf("Error: %v", err)
		return err
	}
	return nil
}

func GetCustomer(id uint) (*Customer, error) {
	customer := &Customer{}
	err := database.DBConn.First(&customer, id).Error
	if err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}

	return customer, nil
}

func GetAllCustomers() ([]Customer, error) {
	var customers []Customer
	err := database.DBConn.Find(&customers).Error
	if err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}

	return customers, nil
}

func UpdateCustomer(customer *Customer) error {
	err := database.DBConn.Save(&customer).Error
	if err != nil {
		log.Printf("Error: %v", err)
		return err
	}

	return nil
}
