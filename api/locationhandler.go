package api

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/yusufatalay/RentACar/models"
)

func GetActiveLocations(c *fiber.Ctx) error {
	var activelocations []models.Location
	var err error
	activelocations, err = models.GetActiveLocations()

	if err != nil {
		log.Printf("Error: %v", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"Code":    models.TechnicalErrors[1].Code,
			"Message": models.TechnicalErrors[1].Message,
		})
	}

	if len(activelocations) == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"Code":    models.BusinessErrors[4].Code,
			"Message": models.BusinessErrors[4].Message,
		})
	}

	err = c.Status(http.StatusOK).JSON(fiber.Map{
		"Message": "Active locations retrieved successfully",
		"Data":    activelocations,
	})

	if err != nil {
		log.Printf("Error: %v", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"Code":    models.TechnicalErrors[2].Code,
			"Message": models.TechnicalErrors[2].Message,
		})
	}
	return nil
}
