package database

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
)

// NewConnection initializes a Firestore client using Application Default Credentials (ADC).
// On Cloud Functions / Cloud Run, ADC is provided automatically by the runtime environment.
// For local development, run: gcloud auth application-default login
func NewConnection() *firestore.Client {
	ctx := context.Background()

	// No explicit credentials option needed — ADC resolves credentials automatically.
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		log.Fatalf("functions.init: NewApp %v\n", err)
	}

	db, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("functions.init: Database init : %v\n", err)
	}

	return db
}
