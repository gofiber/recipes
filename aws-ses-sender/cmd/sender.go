package cmd

import (
	"aws-ses-sender/config"
	"aws-ses-sender/model"
	"aws-ses-sender/pkg/aws"
	"context"
	"strconv"
	"time"
)

var reqChan = make(chan *model.Request, 1000)

// RunSender runs the email sender
// It consumes the email sending requests from the channel and sends them to the AWS SES
func RunSender(ctx context.Context) {
	rateStr := config.GetEnv("EMAIL_RATE", "14")
	rate, _ := strconv.Atoi(rateStr)
	sesClient, err := aws.NewSESClient(ctx)
	if err != nil {
		panic(err)
	}

	db := config.GetDB()
	ticker := time.NewTicker(1 * time.Second / time.Duration(rate))
	for range ticker.C {
		req := <-reqChan
		go func(r *model.Request) {
			// Add code for the open event at the end of the body
			serverHost := config.GetEnv("SERVER_HOST", "http://localhost:3000")
			content := r.Content
			content += `<img src="` + serverHost + `/v1/events/open?requestId=` + strconv.Itoa(int(r.ID)) + `">`
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			msgId, err := sesClient.SendEmail(
				ctx,
				int(r.ID),
				&r.Subject,
				&content,
				[]string{r.To},
			)
			status := model.EmailMessageStatusSent
			errMsg := ""
			if err != nil {
				status = model.EmailMessageStatusFailed
				errMsg = err.Error()
			}
			db.Model(&model.Request{}).
				Where("id = ?", r.ID).
				Updates(model.Request{
					MessageId: msgId,
					Status:    status,
					Error:     errMsg,
				})
		}(req)
	}
}
