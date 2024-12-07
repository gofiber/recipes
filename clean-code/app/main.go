package main

import (
	"context"
	"log"

	"app/datasources"
	"app/datasources/database"
	"app/server"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conf := NewConfiguration()
	db, err := database.NewDatabase(ctx, conf.DatabaseURL)
	if err != nil {
		log.Fatalf("failed to create database: %v", err)
	}
	defer db.CloseConnections()

	app := server.NewServer(ctx, &datasources.DataSources{DB: db})
	log.Fatal(app.Listen(":" + conf.Port))
}
