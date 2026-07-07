// Command api is a Fiber HTTP server that enqueues background jobs onto Asynq.
//
// The API never does slow work inline. POST /enqueue drops a task onto a Redis
// queue and returns immediately; a separate worker process (see ./worker)
// picks it up. This is the same split you'd use for sending email, processing
// images, or any task that shouldn't block the request.
package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/hibiken/asynq"

	"github.com/gofiber/recipes/asynq/task"
)

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func main() {
	redisAddr := getEnv("REDIS_ADDR", "localhost:6379")

	// The client is safe for concurrent use; create one and reuse it.
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr})
	defer client.Close()

	app := fiber.New()
	app.Use(logger.New())

	// POST /enqueue with a JSON body: {"user_id": "...", "email": "..."}.
	// The payload is read from the body rather than the query string, so the
	// email address doesn't land in the request-line access log.
	app.Post("/enqueue", func(c fiber.Ctx) error {
		var body task.EmailWelcomePayload
		if err := c.Bind().Body(&body); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "invalid JSON body")
		}
		if body.Email == "" || body.UserID == "" {
			return fiber.NewError(fiber.StatusBadRequest, "user_id and email are required")
		}

		t, err := task.NewEmailWelcome(body.UserID, body.Email)
		if err != nil {
			return err
		}

		// Options are per-enqueue: which queue, how many retries, when to run.
		info, err := client.Enqueue(t,
			asynq.Queue("default"),
			asynq.MaxRetry(5),
		)
		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{
			"enqueued": true,
			"task_id":  info.ID,
			"queue":    info.Queue,
		})
	})

	log.Fatal(app.Listen(":3000"))
}
