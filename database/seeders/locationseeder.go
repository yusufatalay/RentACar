package database

import (
	"log"

	"github.com/yusufatalay/RentACar/database"
	"github.com/yusufatalay/RentACar/models"
)

func SeedLocations() {
	var locations = []models.Location{
		{Name: "Sabiha ", Active: true},
		{Name: "Ankara", Active: false},
		{Name: "İzmir", Active: true},
		{Name: "Bursa", Active: false},
	}

	err := database.DBConn.Create(&locations).Error
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
