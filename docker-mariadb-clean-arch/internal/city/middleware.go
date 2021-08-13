package city

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

// If city does not exist, do not allow one to access the API.
func (h *CityHandler) checkIfCityExistsMiddleware(c *fiber.Ctx) error {
	// Create a new customized context.
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Fetch parameter.
	targetedCityID, err := c.ParamsInt("cityID")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "Please specify a valid city ID!",
		})
	}

	// Check if city exists.
	searchedCity, err := h.cityService.FetchCity(customContext, targetedCityID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}
	if searchedCity == nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "There is no city with this ID!",
		})
	}

	// Store in locals for further processing in the real handler.
	c.Locals("cityID", targetedCityID)
	return c.Next()
}
