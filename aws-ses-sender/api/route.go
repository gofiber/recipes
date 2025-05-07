package api

import (
	"github.com/gofiber/fiber/v3"
)

// setV1Routes V1 Routes
func setV1Routes(app *fiber.App) {
	// Messages
	app.Post("/v1/messages", createMessageHandler)
	// Topics
	app.Get("/v1/topics/:topicId", getResultCountHandler)
	// Events
	app.Get("/v1/events/open", createOpenEventHandler)
	app.Get("/v1/events/counts/sent", getSentCountHandler)
	app.Post("/v1/events/results", createResultEventHandler)
}
