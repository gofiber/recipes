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

	conf := newConfiguration()
	db, err := database.NewDatabase(ctx, conf.DatabaseURL)
	if err != nil {
		log.Panicf("failed to create database: %v", err)
	}
	defer db.CloseConnections()

	app := server.NewServer(&datasources.DataSources{DB: db})
	log.Panic(app.Listen(":" + conf.Port))
}
