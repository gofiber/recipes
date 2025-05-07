package main

import (
	"aws-ses-sender/api"
	"aws-ses-sender/cmd"
	"aws-ses-sender/config"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/getsentry/sentry-go"
)

func main() {
	// Sentry
	if err := sentry.Init(sentry.ClientOptions{
		Dsn: config.GetEnv("SENTRY_DSN"),
	}); err != nil {
		log.Printf("Sentry initialization failed: %v", err)
	}
	// Ensure flush before exit
	defer sentry.Flush(2 * time.Second)

	// Create a context that will be canceled on interrupt
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Set up a channel to listen for interrupt signals
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cancel()
	}()

	// Message Scheduler
	go cmd.RunScheduler(ctx)

	// Email Sender
	go cmd.RunSender(ctx)

	// HTTP Server
	api.Run(ctx)
}
