package database

import (
	"log"
	"time"

	"github.com/yusufatalay/RentACar/database"
	"github.com/yusufatalay/RentACar/models"
)

func SeedOffices() {
	var offices = []models.Office{
		{Vendor: "Vendor1", LocationID: 1, OpeningHour: time.Now(), ClosingHour: time.Now().Add(time.Hour * 9)},
		{Vendor: "Vendor2", LocationID: 2, OpeningHour: time.Now(), ClosingHour: time.Now().Add(time.Hour * 9)},
		{Vendor: "Vendor3", LocationID: 3, OpeningHour: time.Now(), ClosingHour: time.Now().Add(time.Hour * 9)},
		{Vendor: "Vendor4", LocationID: 4, OpeningHour: time.Now(), ClosingHour: time.Now().Add(time.Hour * 9)},
	}
	err := database.DBConn.Create(&offices).Error
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
