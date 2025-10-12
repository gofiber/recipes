package handlers

import (
	"errors"
	"time"

	"auth-jwt-gorm/services"

	"github.com/gofiber/fiber/v2"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type RefreshResponse struct {
	Token string `json:"token"`
}

// AuthHandler contains HTTP handlers for authentication
type AuthHandler struct {
	authService *services.AuthService
}

// NewAuthHandler creates a new auth handler
func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// Login get user and password
func (ah *AuthHandler) Login(c *fiber.Ctx) error {
	var input LoginRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Error on login request",
			"data":    nil,
		})
	}
	token, err := ah.authService.Login(input.Email, input.Password)
	if err != nil {
		if errors.Is(err, services.ErrInvalidCredentials) {

			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  "error",
				"message": "Invalid credentials",
				"data":    nil,
			})
		} else {

			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": "Internal Server Error",
				"data":    nil,
			})
		}
	}

	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 72),
		HTTPOnly: true,
	})

	// Return the token
	response := LoginResponse{Token: token}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Success login",
		"data":    response,
	})
}

func (ah *AuthHandler) Logout(c *fiber.Ctx) error {
	// Clear cookie
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	})

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Success logout",
		"data":    nil,
	})
}

func (ah *AuthHandler) Register(c *fiber.Ctx) error {
	var input RegisterRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Error on register request",
			"data":    nil,
		})
	}

	user, err := ah.authService.Register(input.Email, input.Username, input.Password)
	if err != nil {
		if errors.Is(err, services.ErrEmailInUse) {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"status":  "error",
				"message": "Email already in use",
				"data":    nil,
			})
		}
	}

	newUser := RegisterResponse{
		Email:    user.Email,
		Username: user.Username,
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Success register",
		"data":    newUser,
	})
}

func (ah *AuthHandler) RefreshToken(c *fiber.Ctx) error {
	var input RefreshRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request payload",
			"data":    nil,
		})
	}
	token, err := ah.authService.RefreshAccessToken(input.RefreshToken)
	if err != nil {
		if errors.Is(err, services.ErrInvalidToken) || errors.Is(err, services.ErrExpiredToken) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  "error",
				"message": "Invalid or expired refresh token",
				"data":    nil,
			})
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": "Internal server error",
				"data":    nil,
			})
		}
	}
	// Clear cookie
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 72),
		HTTPOnly: true,
	})

	response := RefreshResponse{Token: token}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Success refresh token",
		"data":    response,
	})
}
