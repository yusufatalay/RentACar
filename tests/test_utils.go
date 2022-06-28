package tests

import (
	"github.com/yusufatalay/RentACar/database"
	"github.com/yusufatalay/RentACar/models"
)

func WipeDatabase() {
	database.DBConn.Where("1 =1").Delete(&models.CarsReservation{})
	database.DBConn.Where("1 =1").Delete(&models.Customer{})
	database.DBConn.Where("1 =1").Delete(&models.Car{})
	database.DBConn.Where("1 =1").Delete(&models.OfficesWorkingDay{})
	database.DBConn.Where("1 =1").Delete(&models.Office{})
	database.DBConn.Where("1 =1").Delete(&models.Location{})

}
