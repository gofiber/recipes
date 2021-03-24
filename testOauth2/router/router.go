package router

import (
	"testOauth2/handlers"
	"testOauth2/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// SetupRoutes prepares required routes
func SetupRoutes(app *fiber.App) {

	app.Use(cors.New(cors.Config{
		// attempt to mitigate CORS issues - pay attention to last /
		AllowOrigins: "http://localhost:8080, http://localhost:8080/, https://api.github.com/user, https://api.github.com/user/,", //
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		Next:         nil,
	}))

	// display a nice call trae on the console
	app.Use(logger.New())

	// if you want to prevent crashes
	app.Use(recover.New())

	// add a standard redirect to the index page
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/index.html", fiber.StatusTemporaryRedirect)
	})

	// return pages from their templates
	app.Get("/:template", handlers.HTMLPages)

	// perform logout - in fact only the local session is destroyed
	app.Get("/logout", middleware.OAUTHDisconnect)

	// display the "forbidden" page - but only if the middleware agrees with it
	app.Get("/protected", middleware.OAUTHProtected, middleware.OAUTHGETHandler)

	// perform the "magic" - exdecutes the whole GitHub authentication routine
	app.Get("/oauth/redirect", middleware.OAUTHRedirect)
}
