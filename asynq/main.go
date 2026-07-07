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

	// POST /enqueue?email=a@b.com&user=123
	// Enqueues a welcome-email job and returns the task id.
	app.Post("/enqueue", func(c fiber.Ctx) error {
		email := c.Query("email")
		userID := c.Query("user")
		if email == "" || userID == "" {
			return fiber.NewError(fiber.StatusBadRequest, "email and user are required")
		}

		t, err := task.NewEmailWelcome(userID, email)
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
