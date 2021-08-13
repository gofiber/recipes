package user

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

// Represents our handler with our use-case / service.
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
func (h *UserHandler) getUsers(c *fiber.Ctx) error {
	// Create cancellable context.
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Get all users.
	users, err := h.userService.GetUsers(customContext)
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
func (h *UserHandler) getUser(c *fiber.Ctx) error {
	// Create cancellable context.
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Fetch parameter.
	targetedUserID, err := c.ParamsInt("userID")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "Please specify a valid user ID!",
		})
	}

	// Get one user.
	user, err := h.userService.GetUser(customContext, targetedUserID)
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
func (h *UserHandler) createUser(c *fiber.Ctx) error {
	// Initialize variables.
	user := &User{}

	// Create cancellable context.
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Parse request body.
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	// Create one user.
	err := h.userService.CreateUser(customContext, user)
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
func (h *UserHandler) updateUser(c *fiber.Ctx) error {
	// Initialize variables.
	user := &User{}
	targetedUserID := c.Locals("userID").(int)

	// Create cancellable context.
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Parse request body.
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	// Update one user.
	err := h.userService.UpdateUser(customContext, targetedUserID, user)
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
func (h *UserHandler) deleteUser(c *fiber.Ctx) error {
	// Initialize previous user ID.
	targetedUserID := c.Locals("userID").(int)

	// Create cancellable context.
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Delete one user.
	err := h.userService.DeleteUser(customContext, targetedUserID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	// Return 204 No Content.
	return c.SendStatus(fiber.StatusNoContent)
}
