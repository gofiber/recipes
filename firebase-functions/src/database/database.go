package database

import (
	"context"
	"encoding/json"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"

	"google.golang.org/api/option"
)

func NewConnection() *firestore.Client {
	ctx := context.Background()

	sa := option.WithCredentialsJSON(credentials())
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalf("functions.init: NewApp %v\n", err)
	}

	db, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("functions.init: Database init : %v\n", err)
	}

	return db
}

func credentials() []byte {
	// TODO: Replace with your Credentials
	data := map[string]interface{}{
		"type":                        "",
		"project_id":                  "",
		"private_key_id":              "",
		"private_key":                 "",
		"client_email":                "",
		"client_id":                   "",
		"auth_uri":                    "",
		"token_uri":                   "",
		"auth_provider_x509_cert_url": "",
		"client_x509_cert_url":        "",
		"universe_domain":             "",
	}

	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	return bytes
}
