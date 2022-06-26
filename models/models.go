package models

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/yusufatalay/RentACar/database"
)

// a global validator variable (Singleton was recommended)
var validate *validator.Validate

func init() {
	//TODO implement logging mechanism

	// initiate the validator
	validate = validator.New()
	err := database.DBConn.AutoMigrate(&Location{}, &Office{}, &Car{}, &OfficesWorkingDay{}, &Error{}, &CarsReservation{}, &Customer{})

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

}
