package user

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v3"
)

// Represents our handler with our use-case / service.
// requestTimeout bounds a single data layer call, not a whole request.
// Routes that run the existence middleware before the handler perform two
// independent calls and can therefore wait up to twice this long in total.
// Without it a hung query would block indefinitely, since the request
// context alone is only cancelled when the client disconnects.
const requestTimeout = 10 * time.Second

type UserHandler struct {
	userService UserService
}

// Creates a new handler.
func NewUserHandler(userRoute fiber.Router, us UserService) {
	// Create a handler based on our created service / use-case.
	handler := &UserHandler{
		userService: us,
	}

	// Declare routing endpoints for general routes.
	userRoute.Get("", handler.getUsers)
	userRoute.Post("", handler.createUser)

	// Declare routing endpoints for specific routes.
	userRoute.Get("/:userID", handler.getUser)
	userRoute.Put("/:userID", handler.checkIfUserExistsMiddleware, handler.updateUser)
	userRoute.Delete("/:userID", handler.checkIfUserExistsMiddleware, handler.deleteUser)
}

// Gets all users.
func (h *UserHandler) getUsers(c fiber.Ctx) error {
	// Get all users.
	ctx, cancel := context.WithTimeout(c.Context(), requestTimeout)
	defer cancel()

	users, err := h.userService.GetUsers(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	// Return results.
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   users,
	})
}

// Gets a single user.
func (h *UserHandler) getUser(c fiber.Ctx) error {
	// Fetch parameter.
	targetedUserID, err := fiber.Params[int](c, "userID"), error(nil)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "Please specify a valid user ID!",
		})
	}

	// Get one user.
	ctx, cancel := context.WithTimeout(c.Context(), requestTimeout)
	defer cancel()

	user, err := h.userService.GetUser(ctx, targetedUserID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	// Return results.
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   user,
	})
}

// Creates a single user.
func (h *UserHandler) createUser(c fiber.Ctx) error {
	// Initialize variables.
	user := &User{}

	// Parse request body.
	if err := c.Bind().Body(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	// Create one user.
	ctx, cancel := context.WithTimeout(c.Context(), requestTimeout)
	defer cancel()

	err := h.userService.CreateUser(ctx, user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	// Return result.
	return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"status":  "success",
		"message": "User has been created successfully!",
	})
}

// Updates a single user.
func (h *UserHandler) updateUser(c fiber.Ctx) error {
	// Initialize variables.
	user := &User{}
	targetedUserID := c.Locals("userID").(int)

	// Parse request body.
	if err := c.Bind().Body(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	// Update one user.
	ctx, cancel := context.WithTimeout(c.Context(), requestTimeout)
	defer cancel()

	err := h.userService.UpdateUser(ctx, targetedUserID, user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	// Return result.
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "User has been updated successfully!",
	})
}

// Deletes a single user.
func (h *UserHandler) deleteUser(c fiber.Ctx) error {
	// Initialize previous user ID.
	targetedUserID := c.Locals("userID").(int)

	// Delete one user.
	ctx, cancel := context.WithTimeout(c.Context(), requestTimeout)
	defer cancel()

	err := h.userService.DeleteUser(ctx, targetedUserID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	// Return 204 No Content.
	return c.SendStatus(fiber.StatusNoContent)
}
