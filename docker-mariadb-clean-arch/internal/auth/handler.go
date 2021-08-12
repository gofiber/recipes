package auth

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

// Create an authentication handler. Leave this empty, as we have no domains nor use-cases.
// In my opinion, authentication is an implementation detail (framework layer).
type AuthHandler struct{}

// Creates a new authentication handler.
func NewAuthHandler(authRoute fiber.Router) {
	handler := &AuthHandler{}

	// Declare routing for specific routes.
	authRoute.Post("/login", handler.signInUser)
	authRoute.Post("/logout", handler.signOutUser)
	authRoute.Get("/private", JWTMiddleware(), handler.privateRoute)
}

// Signs in a user and gives them a JWT.
func (h *AuthHandler) signInUser(c *fiber.Ctx) error {
	// Create a struct so the request body can be mapped here.
	type loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// Create a struct for our custom JWT payload.
	type jwtClaims struct {
		User string `json:"user"`
		jwt.StandardClaims
	}

	// Get request body.
	request := &loginRequest{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	// If both username and password are incorrect, do not allow access.
	if request.Username != os.Getenv("API_USERNAME") || request.Password != os.Getenv("API_PASSWORD") {
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"status":  "fail",
			"message": "Wrong username or password!",
		})
	}

	// Send back JWT as a cookie.
	// HTTPOnly cookie is disabled - it causes problems with JWT middleware of Fiber.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwtClaims{
		os.Getenv("API_USERNAME"),
		jwt.StandardClaims{
			Audience:  "docker-mariadb-clean-arch-users",
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    "docker-mariadb-clean-arch",
		},
	})
	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    signedToken,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 24),
		Secure:   false,
		HTTPOnly: true,
	})

	// Send also response.
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"token":  signedToken,
	})
}

// Logs out user and removes their JWT.
func (h *AuthHandler) signOutUser(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    "loggedOut",
		Path:     "/",
		Expires:  time.Now().Add(time.Second * 10),
		Secure:   false,
		HTTPOnly: true,
	})

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Logged out successfully!",
	})
}

// A single private route that only logged in users can access.
func (h *AuthHandler) privateRoute(c *fiber.Ctx) error {
	jwtData := c.Locals("user").(*jwt.Token)
	claims := jwtData.Claims.(jwt.MapClaims)
	user := claims["user"]

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Welcome to the private route!",
		"user":    user,
	})
}
