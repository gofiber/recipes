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
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
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

	// Use global middlewares.
	app.Use(cache.New())
	app.Use(cors.New())
	app.Use(compress.New())
	app.Use(logger.New())
	app.Use(recover.New())

	// Create repositories.
	cityRepository := city.NewCityRepository(mariadb)
	userRepository := user.NewUserRepository(mariadb)

	// Create all of our services.
	cityService := city.NewCityService(cityRepository)
	userService := user.NewUserService(userRepository)

	// Prepare our endpoints for the API.
	misc.NewMiscHandler(app.Group("/api/v1"))
	auth.NewAuthHandler(app.Group("/api/v1/auth"))
	city.NewCityHandler(app.Group("/api/v1/cities"), cityService)
	user.NewUserHandler(app.Group("/api/v1/users"), userService)

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
