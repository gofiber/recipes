package main

import (
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

	// Creates a new Fiber instance and group router to '/api/v1' subroutes.
	app := fiber.New()
	api := app.Group("/api/v1")

	// Prepare endpoints for 'miscellaneous' routes, such as health-check, etc.
	misc.NewMiscHandler(api)

	// Prepare endpoints and dependency injection for 'User' entity.
	userRoute := api.Group("/users")
	userRepository := user.NewUserRepository(mariadb)
	userService := user.NewUserService(userRepository)
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
