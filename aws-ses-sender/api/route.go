package api

import (
	"github.com/gofiber/fiber/v3"
)

// setV1Routes V1 Routes
func setV1Routes(app *fiber.App) {
	v1 := app.Group("/v1")
	{
		// Messages
		// Handlers run in the order they are listed, and these handlers never call
		// c.Next(), so apiKeyAuth has to come first or it would never run.
		v1.Post("/messages", apiKeyAuth, createMessageHandler)
		// Topics
		v1.Get("/topics/:topicId", apiKeyAuth, getResultCountHandler)
		// Events
		v1.Get("/events/open", createOpenEventHandler)
		v1.Get("/events/counts/sent", apiKeyAuth, getSentCountHandler)
		v1.Post("/events/results", createResultEventHandler(newSNSVerifier()))
	}
}
