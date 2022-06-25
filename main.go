package main

import (
	"github.com/yusufatalay/RentACar/database"
	_ "github.com/yusufatalay/RentACar/models"
)

func main() {

	db, err := database.DBConn.DB()
	if err != nil {
		panic(err)
	}

	defer db.Close()
}
