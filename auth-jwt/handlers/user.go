package handlers

import (
	"auth-jwt-gorm/models"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// AuthHandler contains HTTP handlers for authentication
type UserHandler struct {
	userRepo *models.UserRepository
}

// NewAuthHandler creates a new auth handler
func NewUserHandler(userRepo *models.UserRepository) *UserHandler {
	return &UserHandler{
		userRepo: userRepo,
	}
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
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

// GetUser get a user
func (uh *UserHandler) GetUser(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}
	user, err := uh.userRepo.GetUserByID(uint(id))
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

// CreateUser new user
func (uh *UserHandler) CreateUser(c fiber.Ctx) error {
	var user models.User
	if err := c.Bind().Body(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Review your input",
			"data":    err,
		})
	}

	if _, err := uh.userRepo.GetUserByEmail(user.Email); err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "User already exists",
			"data":    nil,
		})
	}

	hash, err := hashPassword(user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Couldn't hash password",
			"data":    nil,
		})
	}

	user.Password = hash
	if _, err := uh.userRepo.CreateUser(user.Email, user.Username, user.Password); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Couldn't create user",
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Created user",
		"data": fiber.Map{
			"username": user.Username,
			"email":    user.Email,
		},
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

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	tok, ok := c.Locals("user").(*jwt.Token)
	if !ok || !validToken(tok, strconv.Itoa(id)) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid token id",
			"data":    nil,
		})
	}

	user, err := uh.userRepo.GetUserByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "User not found",
			"data":    nil,
		})
	}

	user.Names = input.Names
	updatedUser, err := uh.userRepo.UpdateUser(uint(id), *user)
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
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	tok, ok := c.Locals("user").(*jwt.Token)
	if !ok || !validToken(tok, strconv.Itoa(id)) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid token id",
			"data":    nil,
		})
	}

	err = uh.userRepo.DeleteUser(uint(id))
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
