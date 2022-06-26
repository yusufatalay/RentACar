package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/yusufatalay/RentACar/api/routes"
	"github.com/yusufatalay/RentACar/database"
	_ "github.com/yusufatalay/RentACar/models"
)

var App *fiber.App

func main() {

	App = fiber.New()

	db, err := database.DBConn.DB()
	if err != nil {
		log.Fatalf("Database Error: %v", err)
	}
	// register routes
	routes.Routes(App)

	err = App.Listen(":3000") // app runs at port 3000

	if err != nil {
		log.Fatalf("Cannot start the app: %v", err)
	}

	defer db.Close()
}
