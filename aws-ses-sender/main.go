package main

import (
	"aws-ses-sender/api"
	"aws-ses-sender/cmd"
	"aws-ses-sender/config"
	"log"

	"github.com/getsentry/sentry-go"
)

func main() {
	// Sentry
	if err := sentry.Init(sentry.ClientOptions{
		Dsn: config.GetEnv("SENTRY_DSN"),
	}); err != nil {
		log.Printf("Sentry initialization failed: %v", err)
	}

	// Message Scheduler
	go cmd.RunScheduler()

	// Email Sender
	go cmd.RunSender()

	// HTTP Server
	api.Run()
}
