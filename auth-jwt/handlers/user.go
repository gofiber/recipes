package handlers

import (
	"auth-jwt-gorm/models"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

// UserHandler contains HTTP handlers for users.
type UserHandler struct {
	userRepo         *models.UserRepository
	refreshTokenRepo *models.RefreshTokenRepository
}

// NewUserHandler creates a new user handler.
func NewUserHandler(userRepo *models.UserRepository, refreshTokenRepo *models.RefreshTokenRepository) *UserHandler {
	return &UserHandler{
		userRepo:         userRepo,
		refreshTokenRepo: refreshTokenRepo,
	}
}

func validToken(t *jwt.Token, id string) bool {
	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return false
	}
	sub, ok := claims["sub"].(string)
	if !ok {
		return false
	}
	return sub == id
}

func parseUserID(c fiber.Ctx) (uint, string, error) {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		return 0, "", c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid user ID",
			"data":    nil,
		})
	}

	return uint(id), strconv.Itoa(id), nil
}

// GetUser get a user
func (uh *UserHandler) GetUser(c fiber.Ctx) error {
	id, idString, err := parseUserID(c)
	if err != nil {
		return err
	}

	tok, ok := c.Locals("user").(*jwt.Token)
	if !ok || !validToken(tok, idString) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid token id",
			"data":    nil,
		})
	}

	user, err := uh.userRepo.GetUserByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "No user found with ID",
			"data":    nil,
		})
	}
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "User found",
		"data":    user,
	})
}

// UpdateUser update user
func (uh *UserHandler) UpdateUser(c fiber.Ctx) error {
	var input struct {
		Names string `json:"names"`
	}
	if err := c.Bind().Body(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Review your input",
			"data":    err,
		})
	}

	id, idString, err := parseUserID(c)
	if err != nil {
		return err
	}

	tok, ok := c.Locals("user").(*jwt.Token)
	if !ok || !validToken(tok, idString) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid token id",
			"data":    nil,
		})
	}

	updatedUser, err := uh.userRepo.UpdateNames(id, input.Names)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "User update failed",
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "User successfully updated",
		"data":    updatedUser,
	})
}

// DeleteUser delete user
func (uh *UserHandler) DeleteUser(c fiber.Ctx) error {
	id, idString, err := parseUserID(c)
	if err != nil {
		return err
	}

	tok, ok := c.Locals("user").(*jwt.Token)
	if !ok || !validToken(tok, idString) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid token id",
			"data":    nil,
		})
	}

	if err := uh.refreshTokenRepo.RevokeAllUserTokens(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to revoke user tokens",
			"data":    nil,
		})
	}

	err = uh.userRepo.DeleteUser(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "User not found",
			"data":    nil,
		})
	}

	c.Locals("user", nil)
	c.ClearCookie("jwt")

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "User successfully deleted",
		"data":    nil,
	})
}
