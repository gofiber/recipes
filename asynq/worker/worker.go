// Command worker consumes tasks enqueued by the API and processes them.
//
// It runs as its own process so it can be scaled and deployed independently of
// the HTTP server. Asynq handles retries, backoff, and the dead-letter queue;
// the handler just does the work and returns an error if it fails, so Asynq
// can retry.
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/hibiken/asynq"

	"github.com/gofiber/recipes/asynq/task"
)

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

// handleEmailWelcome is the processor for TypeEmailWelcome tasks. Returning an
// error tells Asynq to retry (up to the task's MaxRetry); returning nil marks
// it done.
func handleEmailWelcome(_ context.Context, t *asynq.Task) error {
	var p task.EmailWelcomePayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		// A malformed payload will never succeed on retry — skip retries.
		return asynq.SkipRetry
	}

	// Real work goes here (send the email, call an API, etc.).
	log.Printf("sending welcome email to %s (user %s)", p.Email, p.UserID)
	return nil
}

func main() {
	redisAddr := getEnv("REDIS_ADDR", "localhost:6379")

	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisAddr},
		asynq.Config{
			// Total worker goroutines across all queues.
			Concurrency: 10,
			// Weighted priority: "critical" is drained ~6x as often as "low".
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
		},
	)

	mux := asynq.NewServeMux()
	mux.HandleFunc(task.TypeEmailWelcome, handleEmailWelcome)

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run asynq server: %v", err)
	}
}
