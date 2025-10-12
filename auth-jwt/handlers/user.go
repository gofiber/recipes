package handlers

import (
	"strconv"

	"auth-jwt-gorm/database"
	"auth-jwt-gorm/models"

	"github.com/gofiber/fiber/v2"
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
	n, err := strconv.Atoi(id)
	if err != nil {
		return false
	}
	claims := t.Claims.(jwt.MapClaims)
	uid := int(claims["user_id"].(float64))
	return uid == n
}

func getUserByID(id string) (*models.User, error) {
	db := database.DB
	var user models.User
	if err := db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUser get a user
func (uh *UserHandler) GetUser(c *fiber.Ctx) error {
	user, err := uh.userRepo.GetUserByID(c.Params("id"))
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
func (uh *UserHandler) CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Review your input",
			"data":    err,
		})
	}

	if _, err := uh.userRepo.GetUserByEmail(user.Email); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "User already exists",
			"data":    err,
		})
	}

	hash, err := hashPassword(user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Couldn't hash password",
			"data":    err,
		})
	}

	user.Password = hash
	if _, err := uh.userRepo.CreateUser(user.Email, user.Username, user.Password); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Couldn't create user",
			"data":    err,
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
func (uh *UserHandler) UpdateUser(c *fiber.Ctx) error {
	var input struct {
		Names string `json:"names"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Review your input",
			"data":    err,
		})
	}

	id := c.Params("id")
	if !validToken(c.Locals("user").(*jwt.Token), id) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid token id",
			"data":    nil,
		})
	}

	user, err := getUserByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "User not found",
			"data":    nil,
		})
	}

	user.Names = input.Names
	database.DB.Save(user)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "User successfully updated",
		"data":    user,
	})
}

// DeleteUser delete user
func (uh *UserHandler) DeleteUser(c *fiber.Ctx) error {
	var input struct {
		Password string `json:"password"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Review your input",
			"data":    err,
		})
	}

	id := c.Params("id")
	if !validToken(c.Locals("user").(*jwt.Token), id) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid token id",
			"data":    nil,
		})
	}

	
	 err := uh.userRepo.DeleteUser(id)
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
