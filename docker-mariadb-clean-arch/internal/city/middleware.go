package city

import (
	"context"

	"github.com/gofiber/fiber/v3"
)

// If city does not exist, do not allow one to access the API.
func (h *CityHandler) checkIfCityExistsMiddleware(c fiber.Ctx) error {
	// Fetch parameter.
	targetedCityID, err := fiber.Params[int](c, "cityID"), error(nil)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "Please specify a valid city ID!",
		})
	}

	// Check if city exists.
	// cancel() is called directly instead of deferred: this is middleware, so
	// a deferred call would only run after c.Next() returns and would keep the
	// timer alive for the whole downstream chain. The row is fully scanned
	// before the call returns, so releasing the context here is safe.
	ctx, cancel := context.WithTimeout(c.Context(), requestTimeout)
	searchedCity, err := h.cityService.FetchCity(ctx, targetedCityID)
	cancel()
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
