package tests

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/yusufatalay/RentACar/database"
	"github.com/yusufatalay/RentACar/models"
)

func TestWhenThereIsNoReservations(t *testing.T) {
	// clean out the database
	WipeDatabase()
	// at the start database is empty so there is no reservations
	result, err := models.GetAllReservations()

	// err should be nil
	assert.Nil(t, err, "there should be no error but got %v", err)

	// length of the result should be 0
	assert.Equal(t, 0, len(result), "there should be no reservations")
}

// I know it's not best practice for writing tests, howerver I am running short in time
func TestWhenThereAreReservations(t *testing.T) {
	// in order to create a reservation we need a car, customer and office
	// clean out the database
	WipeDatabase()
	var location models.Location
	var office models.Office
	var car models.Car
	var customer models.Customer
	// create an active location
	err := database.DBConn.Create(&models.Location{Name: "Location 1", Active: "true"}).Error
	assert.Nil(t, err, "there should be no error but got %v", err)

	err = database.DBConn.Raw("SELECT id FROM locations WHERE name = ?", "Location 1").Scan(&location).Error
	assert.Nil(t, err, "there should be no error but got %v", err)

	// create the office
	err = database.DBConn.Create(&models.Office{Vendor: "Office 1", LocationID: location.ID, OpeningHour: time.Now(), ClosingHour: time.Now().Add(24 * time.Hour)}).Error
	assert.Nil(t, err, "there should be no error but got %v", err)
	// get office's ID
	err = database.DBConn.Raw("SELECT id FROM offices WHERE vendor = ?", "Office 1").Scan(&office).Error
	assert.Nil(t, err, "there should be no error but got %v", err)

	// create the car
	err = database.DBConn.Create(&models.Car{Vendor: "Office 1", OfficeID: office.ID, Fuel: "Benzin", Transmission: "Auto", Name: "Car1"}).Error
	assert.Nil(t, err, "there should be no error but got %v", err)
	// get car id
	err = database.DBConn.Raw("SELECT id FROM cars WHERE vendor = ?", "Office 1").Scan(&car).Error

	assert.Nil(t, err, "there should be no error but got %v", err)
	// create the customer
	err = database.DBConn.Create(&models.Customer{FirstName: "CustomerFN", LastName: "CustomerLN", TCID: "12342345324", DOB: time.Now(), PhoneNumber: "+16175551212"}).Error

	assert.Nil(t, err, "there should be no error but got %v", err)
	// get customer id
	err = database.DBConn.Raw("SELECT id FROM customers WHERE first_name = ?", "CustomerFN").Scan(&customer).Error
	assert.Nil(t, err, "there should be no error but got %v", err)

	// make a reservation

	payload := models.CarsReservation{
		CarID:            car.ID,
		LeaserID:         customer.ID,
		LocationID:       location.ID,
		ReservationBegin: time.Now().Add(1 * time.Hour),
		ReservationEnd:   time.Now().Add(2 * time.Hour),
	}

	err = models.CreateReservation(&payload)

	assert.Nil(t, err, "there should be no error but got %v", err)

	// now we should have 1 reservation

	result, err := models.GetAllReservations()

	assert.Nil(t, err, "there should be no error but got %v", err)

	assert.Equal(t, 1, len(result), "there should be 1 reservation")

	// wipe the database
	WipeDatabase()

}
