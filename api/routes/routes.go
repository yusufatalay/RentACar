package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yusufatalay/RentACar/api"
	"github.com/gofiber/swagger"
	_ "github.com/yusufatalay/RentACar/docs"
)

func Routes(app *fiber.App) {
	root := app.Group("/api/v1")

	locations := root.Group("/locations")
	locations.Get("/active", api.GetActiveLocations) // list all active locations

	reservations := root.Group("/reservations")
	reservations.Get("/", api.GetAllReservations)      // list all reservations
	reservations.Post("/", api.ReserveACar)            // create a new reservation")
	reservations.Post("/lookup", api.GetAvailableCars) // list all available cars


	// Register swagger's endpoint
	app.Get("/swagger/*", swagger.HandlerDefault) // default
}
