package main

import (
	"fiber-docker-nginx/api"
	"fiber-docker-nginx/database"
	"log"
)

func main() {
	if !database.ConnectionOK() {
		log.Fatal("Not connected to DB")
		return
	}
	api.Init()
}
