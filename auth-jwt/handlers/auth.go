package handlers

import (
	"errors"
	"net/mail"
	"strconv"
	"strings"
	"time"

	"auth-jwt-gorm/services"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
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
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type RefreshResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
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

func getAccessTokenExpiry(token string) time.Time {
	expiration := time.Now().Add(15 * time.Minute)
	parser := new(jwt.Parser)
	claims := jwt.MapClaims{}

	if _, _, err := parser.ParseUnverified(token, claims); err != nil {
		return expiration
	}

	exp, ok := claims["exp"].(float64)
	if !ok || exp <= 0 {
		return expiration
	}
	return time.Unix(int64(exp), 0)
}

func setJWTCookie(c fiber.Ctx, token string) {
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  getAccessTokenExpiry(token),
		HTTPOnly: true,
		Secure:   c.Protocol() == "https",
		SameSite: "Lax",
	})
}

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// Login get user and password
func (ah *AuthHandler) Login(c fiber.Ctx) error {
	var input LoginRequest
	if err := c.Bind().Body(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Error on login request",
			"data":    nil,
		})
	}

	if strings.TrimSpace(input.Email) == "" || strings.TrimSpace(input.Password) == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Email and password are required",
			"data":    nil,
		})
	}
	// Authenticate the user
	accessToken, refreshToken, err := ah.authService.LoginWithRefresh(input.Email, input.Password, 30*24*time.Hour)
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

	setJWTCookie(c, accessToken)

	// Return the token
	response := LoginResponse{AccessToken: accessToken, RefreshToken: refreshToken}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Success login",
		"data":    response,
	})
}

func (ah *AuthHandler) Logout(c fiber.Ctx) error {
	tok, ok := c.Locals("user").(*jwt.Token)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid authentication token",
			"data":    nil,
		})
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid authentication token claims",
			"data":    nil,
		})
	}

	sub, ok := claims["sub"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid authentication token subject",
			"data":    nil,
		})
	}

	userID, err := strconv.ParseUint(sub, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid authentication token subject",
			"data":    nil,
		})
	}

	if err := ah.authService.RevokeAllUserRefreshTokens(uint(userID)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to revoke refresh tokens on logout",
			"data":    nil,
		})
	}

	// Clear cookie
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
		Secure:   c.Protocol() == "https",
		SameSite: "Lax",
	})

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Success logout",
		"data":    nil,
	})
}

func (ah *AuthHandler) Register(c fiber.Ctx) error {
	var input RegisterRequest
	if err := c.Bind().Body(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Error on register request",
			"data":    nil,
		})
	}
	if strings.TrimSpace(input.Email) == "" || strings.TrimSpace(input.Username) == "" || strings.TrimSpace(input.Password) == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Email, username, and password are required",
			"data":    nil,
		})
	}
	if !isValidEmail(input.Email) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid email",
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
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Error on registering user",
			"data":    nil,
		})
	}

	newUser := RegisterResponse{
		Id:       strconv.FormatUint(uint64(user.ID), 10),
		Email:    user.Email,
		Username: user.Username,
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Success register",
		"data":    newUser,
	})
}

func (ah *AuthHandler) RefreshToken(c fiber.Ctx) error {
	var input RefreshRequest
	if err := c.Bind().Body(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request payload",
			"data":    nil,
		})
	}
	token, newRefreshToken, err := ah.authService.RefreshAccessToken(input.RefreshToken)
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
	setJWTCookie(c, token)

	response := RefreshResponse{Token: token, RefreshToken: newRefreshToken}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Success refresh token",
		"data":    response,
	})
}
