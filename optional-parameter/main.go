// ⚡️ Fiber is an Express inspired web framework written in Go with ☕️
// 🤖 Github Repository: https://github.com/gofiber/fiber
// 📌 API Documentation: https://docs.gofiber.io

package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// user list
	users := [...]string{"Alice", "Bob", "Charlie", "David"}

	// Fiber instance
	app := fiber.New()

	// Route to profile
	app.Get("/:id?", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id")) // transform id to array index
		if err != nil || id < 0 || id >= len(users) {
			return c.SendStatus(http.StatusNotFound) // invalid parameter returns 404
		}
		return c.SendString("Hello, " + users[id] + "!") // custom hello message to user with the id
	})

	// Start server
	log.Fatal(app.Listen(":3000"))
}
