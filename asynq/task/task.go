// Package task defines the job types shared between the API (which enqueues)
// and the worker (which processes). Keeping the payload and type name in one
// place stops the two sides from drifting apart.
package task

import (
	"encoding/json"

	"github.com/hibiken/asynq"
)

// TypeEmailWelcome is the task type name. Asynq routes by this string.
const TypeEmailWelcome = "email:welcome"

// EmailWelcomePayload is what the API sends and the worker receives.
type EmailWelcomePayload struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
}

// NewEmailWelcome builds an enqueueable task from a payload.
func NewEmailWelcome(userID, email string) (*asynq.Task, error) {
	payload, err := json.Marshal(EmailWelcomePayload{UserID: userID, Email: email})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeEmailWelcome, payload), nil
}
