package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var DB *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "user"
	password = "password"
	dbname   = "fiber_demo"
)

func ConnectDB() {
	// Connect to the database
	var connStr string = fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Check if the connection is successful
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	boil.SetDB(db)
	DB = db
	fmt.Println("😁 Connected to database")
}
