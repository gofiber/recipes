package main

import (
	"docker-mariadb-clean-arch/internal/auth"
	"docker-mariadb-clean-arch/internal/city"
	"docker-mariadb-clean-arch/internal/infrastructure"
	"docker-mariadb-clean-arch/internal/misc"
	"docker-mariadb-clean-arch/internal/user"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Try to connect to our database as the initial part.
	mariadb, err := infrastructure.ConnectToMariaDB()
	if err != nil {
		log.Fatal("Database connection error: $s", err)
	}

	// Creates a new Fiber instance.
	app := fiber.New(fiber.Config{
		AppName:      "Docker MariaDB Clean Arch",
		ServerHeader: "Fiber",
	})

	// Prepare to group all routes to '/api/v1'.
	api := app.Group("/api/v1")

	// Create repositories.
	cityRepository := city.NewCityRepository(mariadb)
	userRepository := user.NewUserRepository(mariadb)

	// Create all of our services.
	cityService := city.NewCityService(cityRepository)
	userService := user.NewUserService(userRepository)

	// Prepare endpoints for 'auth' routes.
	authRoute := api.Group("/auth")
	auth.NewAuthHandler(authRoute)

	// Prepare endpoints for 'City' entity.
	cityRoute := api.Group("/cities")
	city.NewCityHandler(cityRoute, cityService, auth.JWTMiddleware(), auth.GetDataFromJWT)

	// Prepare endpoints for 'miscellaneous' routes, such as health-check, etc.
	misc.NewMiscHandler(api)

	// Prepare endpoints and dependency injection for 'User' entity.
	userRoute := api.Group("/users")
	user.NewUserHandler(userRoute, userService)

	// Prepare an endpoint for 'Not Found'.
	app.All("*", func(c *fiber.Ctx) error {
		errorMessage := fmt.Sprintf("Route '%s' does not exist in this API!", c.OriginalURL())

		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"status":  "fail",
			"message": errorMessage,
		})
	})

	// Listen to port 8080.
	log.Fatal(app.Listen(":8080"))
}
