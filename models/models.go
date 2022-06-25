package models

import (
	"fmt"

	"github.com/yusufatalay/RentACar/database"
)

func init() {
	//TODO implement logging mechanism
	err := database.DBConn.AutoMigrate(&Location{}, &Office{}, &Car{}, &OfficesWorkingDay{}, &Error{}, &CarsReservation{}, &Customer{})
	fmt.Println("models here ")
	if err != nil {
		panic(err)
	}
}
