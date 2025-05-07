package cmd

import (
	"aws-ses-sender/config"
	"aws-ses-sender/model"
	"context"
	"log"
	"strconv"
	"time"
)

// RunScheduler runs the scheduler
// It schedules the email sending requests to be processed by the sender
func RunScheduler(ctx context.Context) {
	db := config.GetDB()
	sendPerSecStr := config.GetEnv("EMAIL_RATE", "14")
	sendPerSec, _ := strconv.Atoi(sendPerSecStr)
	sendPerMin := sendPerSec * 60
	batchSize := 1000

	ticker := time.NewTicker(1 * time.Minute)
	for range ticker.C {
		for i := 0; i < sendPerMin; i += batchSize {
			now := time.Now()
			reqs := make([]*model.Request, 0, batchSize)
			err := db.Raw(`
				WITH locked_requests AS (
					SELECT id
					FROM email_requests
					WHERE status = ? AND (scheduled_at <= ? OR scheduled_at IS NULL) AND deleted_at IS NULL
					ORDER BY id ASC
					LIMIT ?
					FOR UPDATE SKIP LOCKED
				)
				UPDATE email_requests
				SET status = ?, updated_at = ?
				FROM locked_requests
				WHERE email_requests.id = locked_requests.id
				RETURNING email_requests.*;
			`,
				model.EmailMessageStatusCreated,
				now,
				batchSize,
				model.EmailMessageStatusProcessing,
				now,
			).Scan(&reqs).Error

			if err != nil {
				log.Printf("Update Returning Error: %v", err)
			} else if len(reqs) > 0 {
				for _, req := range reqs {
					reqChan <- req
				}
			} else {
				break
			}
		}
	}
}
