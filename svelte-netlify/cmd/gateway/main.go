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
	// http.Request, so the Fiber app is exposed as a plain http.Handler. The
	// host argument is unused on Lambda.
	log.Fatal(gateway.ListenAndServe("n/a", adaptor.FiberApp(app)))
}
