package main

import (
	"context"
	"log"

	"app/datasources"
	"app/datasources/database"
	"app/server"
)

func main() {
	ctx := context.Background()
	conf := NewConfiguration()
	dataSources := &datasources.DataSources{
		DB: database.NewDatabase(ctx, conf.DatabaseURL),
	}
	defer dataSources.DB.CloseConnections()

	app := server.NewServer(ctx, dataSources)
	log.Fatal(app.Listen(":" + conf.Port))
}
