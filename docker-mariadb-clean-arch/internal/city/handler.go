package city

import (
	"context"
	"time"

	"docker-mariadb-clean-arch/internal/auth"

	"github.com/gofiber/fiber/v3"
)

// requestTimeout bounds how long a single request may spend in the data
// layer. Without it a hung query would block the handler indefinitely,
// since the request context alone is only cancelled when the client
// disconnects.
const requestTimeout = 10 * time.Second

// We will inject our dependency - the service - here.
type CityHandler struct {
	cityService CityService
}

// Creates a new handler.
func NewCityHandler(cityRoute fiber.Router, cs CityService) {
	// Create a handler based on our created service / use-case.
	handler := &CityHandler{
		cityService: cs,
	}

	// We will restrict this route with our JWT middleware.
	// You can inject other middlewares if you see fit here.
	cityRoute.Use(auth.JWTMiddleware(), auth.GetDataFromJWT)

	// Routing for general routes.
	cityRoute.Get("", handler.getCities)
	cityRoute.Post("", handler.createCity)

	// Routing for specific routes.
	cityRoute.Get("/:cityID", handler.getCity)
	cityRoute.Put("/:cityID", handler.checkIfCityExistsMiddleware, handler.updateCity)
	cityRoute.Delete("/:cityID", handler.checkIfCityExistsMiddleware, handler.deleteCity)
}

// Handler to get all cities.
func (h *CityHandler) getCities(c fiber.Ctx) error {
	// Get all cities.
	ctx, cancel := context.WithTimeout(c.Context(), requestTimeout)
	defer cancel()

	cities, err := h.cityService.FetchCities(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	// Return results.
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   cities,
	})
}

// Get one city.
func (h *CityHandler) getCity(c fiber.Ctx) error {
	// Fetch parameter.
	targetedCityID, err := fiber.Params[int](c, "cityID"), error(nil)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "Please specify a valid city ID!",
		})
	}

	// Get one city.
	ctx, cancel := context.WithTimeout(c.Context(), requestTimeout)
	defer cancel()

	city, err := h.cityService.FetchCity(ctx, targetedCityID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	// Return results.
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   city,
	})
}

// Creates a single city.
func (h *CityHandler) createCity(c fiber.Ctx) error {
	// Initialize variables.
	city := &City{}
	currentUserID := c.Locals("currentUser").(int)

	// Parse request body.
	if err := c.Bind().Body(city); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	// Create one city.
	ctx, cancel := context.WithTimeout(c.Context(), requestTimeout)
	defer cancel()

	err := h.cityService.BuildCity(ctx, city, currentUserID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	// Return result.
	return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"status":  "success",
		"message": "City has been created successfully!",
	})
}

// Updates a single city.
func (h *CityHandler) updateCity(c fiber.Ctx) error {
	// Initialize variables.
	city := &City{}
	currentUserID := c.Locals("currentUser").(int)
	targetedCityID := c.Locals("cityID").(int)

	// Parse request body.
	if err := c.Bind().Body(city); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	// Update one city.
	ctx, cancel := context.WithTimeout(c.Context(), requestTimeout)
	defer cancel()

	err := h.cityService.ModifyCity(ctx, targetedCityID, city, currentUserID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	// Return result.
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "City has been updated successfully!",
	})
}

// Deletes a single city.
func (h *CityHandler) deleteCity(c fiber.Ctx) error {
	// Initialize previous city ID.
	targetedCityID := c.Locals("cityID").(int)

	// Delete one city.
	ctx, cancel := context.WithTimeout(c.Context(), requestTimeout)
	defer cancel()

	err := h.cityService.DestroyCity(ctx, targetedCityID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	// Return 204 No Content.
	return c.SendStatus(fiber.StatusNoContent)
}
