package database

import (
	"log"
	"os"

	"ent-mysql/ent"
)

var DBConn *ent.Client

func ConnectDb() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "user:password@tcp(localhost:3306)/dbName?charset=utf8mb4&parseTime=True&loc=Local"
	client, err := ent.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database \n", err)
		os.Exit(1)
	}
	log.Println("Connect")
	DBConn = client
}
