package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func init() {
	var err error
	dsn := "host=localhost user=postgres password=Killer14 dbname=yolcu360DB port=5432 sslmode=disable"
	DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connect to the database: ", err)
	}

}
