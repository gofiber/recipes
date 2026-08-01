package main

import (
	"log"
	"time"

	"github.com/carlmjohnson/gateway"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
	"github.com/gofiber/fiber/v3/middleware/static"

	"github.com/gofiber/recipes/svelte-netlify/handler"
)

func main() {
	app := fiber.New()
	app.Get("/*", static.New("./public"))
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendFile("index")
	})
	app.Get("/api/:ip", handler.CacheRequest(10*time.Minute), handler.GeoLocation)

	// gateway translates the API Gateway proxy event Netlify delivers into an
	// http.Request, so the Fiber app is exposed as a plain http.Handler.
	//
	// Despite the name, the first argument is not a listen address and is never
	// parsed as one. It is only used as a fallback Host for the constructed
	// request when the event does not carry one. "n/a" follows the library's own
	// example.
	log.Fatal(gateway.ListenAndServe("n/a", adaptor.FiberApp(app)))
}
