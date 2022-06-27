package api

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/yusufatalay/RentACar/models"
)

// ReserveACar godoc
// @Summary    	Reserve an available car
// @Description Send reservation request for a car if its successfull then car is reserved to the given user
// @Tags         reservations
// @Accept       json
// @Produce      json
// @Param        account  body      models.CarReservationModel true  "Reserve a car"
// @Success      200      {object} 	models.SuccessfullReservation
// @Failure 	 204 	  {object}  models.Error
// @Failure      400      {object}  models.Error
// @Failure      404      {object}  models.Error
// @Failure      500      {object} 	models.Error
// @Router       /reservations [post]
func ReserveACar(c *fiber.Ctx) error {

	var payload models.CarReservationModel
	var err error
	if err = c.BodyParser(&payload); err != nil {
		log.Printf("Error: %v", err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"Code":    models.TechnicalErrors[0].Code,
			"Message": models.TechnicalErrors[0].Message,
		})
	}
	err = models.CreateReservation(&payload)
	if err != nil {
		log.Printf("Error: %v", err)
		// based on the error, return appropriate error code and message
		switch err.Error() {
		case "location":
			return c.Status(http.StatusNoContent).JSON(fiber.Map{
				"Code":    models.BusinessErrors[2].Code,
				"Message": models.BusinessErrors[2].Message,
			})
		case "office":
			return c.Status(http.StatusNoContent).JSON(fiber.Map{
				"Code":    models.BusinessErrors[1].Code,
				"Message": models.BusinessErrors[1].Message,
			})
		case "no_vacancy":
			return c.Status(http.StatusNoContent).JSON(fiber.Map{
				"Code":    models.BusinessErrors[5].Code,
				"Message": models.BusinessErrors[5].Message,
			})
		default:
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"Code":    models.TechnicalErrors[1].Code,
				"Message": models.TechnicalErrors[1].Message,
			})
		}
	}

	err = c.Status(http.StatusCreated).JSON(fiber.Map{
		"Message": "Car reservation created successfully",
		"Data":    payload,
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

// GetAllReservations godoc
// @Summary      List all reservations in the system
// @Description   List all reservations in the system
// @Tags        	reservations
// @Accept       json
// @Produce      json
// @Success      200  {object} 	models.SuccessfullAllReservations
// @Failure      400  {object}  models.Error
// @Failure      404  {object}  models.Error
// @Failure      500  {object} 	models.Error
// @Router       /reservations/all [get]
func GetAllReservations(c *fiber.Ctx) error {
	var payload []models.CarsReservation
	var err error
	payload, err = models.GetAllReservations()
	if err != nil {
		log.Printf("Error: %v", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"Code":    models.TechnicalErrors[1].Code,
			"Message": models.TechnicalErrors[1].Message,
		})
	}
	if len(payload) == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"Code":    models.BusinessErrors[6].Code,
			"Message": models.BusinessErrors[6].Message,
		})
	}

	err = c.Status(http.StatusOK).JSON(fiber.Map{
		"Message": "All reservations retrieved successfully",
		"Data":    payload,
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
