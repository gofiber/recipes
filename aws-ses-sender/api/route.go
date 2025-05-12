package api

import (
	"github.com/gofiber/fiber/v3"
)

// setV1Routes V1 Routes
func setV1Routes(app *fiber.App) {
	v1 := app.Group("/v1")
	{
		// Messages
		v1.Post("/messages", createMessageHandler, apiKeyAuth)
		// Topics
		v1.Get("/topics/:topicId", getResultCountHandler, apiKeyAuth)
		// Events
		v1.Get("/events/open", createOpenEventHandler)
		v1.Get("/events/counts/sent", getSentCountHandler, apiKeyAuth)
		v1.Post("/events/results", createResultEventHandler)
	}
}
