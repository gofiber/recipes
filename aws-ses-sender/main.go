package main

import (
	"aws-ses-sender-go/api"
	"aws-ses-sender-go/cmd"
	"aws-ses-sender-go/config"

	"github.com/getsentry/sentry-go"
)

func main() {
	// Sentry
	_ = sentry.Init(sentry.ClientOptions{
		Dsn: config.GetEnv("SENTRY_DSN"),
	})

	// Message Scheduler
	go cmd.RunScheduler()

	// Email Sender
	go cmd.RunSender()

	// HTTP Server
	api.Run()
}
