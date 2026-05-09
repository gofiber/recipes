package database

import (
	"log"
	"os"

	"gorm-postgres/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

// ConnectDb opens the PostgreSQL connection and runs auto-migration.
// Set DB_DSN env var to override the default DSN.
// Default DSN: host=localhost user=postgres password='' dbname=go-db port=5432 sslmode=disable
func ConnectDb() {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		dsn = "host=localhost user=postgres password='' dbname=go-db port=5432 sslmode=disable"
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations")
	db.AutoMigrate(&models.Book{})

	DB = Dbinstance{
		Db: db,
	}
}
