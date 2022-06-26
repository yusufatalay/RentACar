package api

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/yusufatalay/RentACar/models"
)

func CreateCar(c *fiber.Ctx) error {
	var payload models.Car
	var err error
	if err = c.BodyParser(&payload); err != nil {
		log.Printf("Error: %v", err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"Code":    models.TechnicalErrors[0].Code,
			"Message": models.TechnicalErrors[0].Message,
		})
	}

	// write the car to the database
	if err = models.SaveCar(&payload); err != nil {
		log.Printf("Error: %v", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"Code":    models.TechnicalErrors[1].Code,
			"Message": models.TechnicalErrors[1].Message,
		})
	}

	if err = c.Status(http.StatusCreated).JSON(fiber.Map{
		"Message": "Car created successfully",
		"Data":    payload,
	}); err != nil {
		log.Printf("Error: %v", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"Code":    models.TechnicalErrors[2].Code,
			"Message": models.TechnicalErrors[2].Message,
		})
	}
	return nil
}

func GetAvailableCars(c *fiber.Ctx) error {
	var payload models.CarAvailabilityIdentifier
	var availablecars []models.Car
	var err error

	if err = c.BodyParser(&payload); err != nil {
		log.Printf("Error: %v", err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"Code":    models.TechnicalErrors[0].Code,
			"Message": models.TechnicalErrors[0].Message,
		})
	}

	availablecars, err = models.GetAvailableCars(&payload)

	if err != nil {
		log.Printf("Error: %v", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"Code":    models.TechnicalErrors[1].Code,
			"Message": models.TechnicalErrors[1].Message,
		})
	}

	if len(availablecars) == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"Code":    models.BusinessErrors[3].Code,
			"Message": models.BusinessErrors[3].Message,
		})
	}

	if err = c.Status(http.StatusOK).JSON(fiber.Map{
		"Message": "Available cars found",
		"Data":    availablecars,
	}); err != nil {
		log.Printf("Error: %v", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"Code":    models.TechnicalErrors[2].Code,
			"Message": models.TechnicalErrors[2].Message,
		})
	}

	return nil
}
