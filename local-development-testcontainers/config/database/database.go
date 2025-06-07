package database

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB is the underlying database connection
var DB *gorm.DB

// New creates a new database connection
// Helpful in testing to create a new database connection.
func New(connString string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{
		NowFunc: func() time.Time { return time.Now().Local() },
		Logger:  logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("gorm open: %w", err)
	}

	return db, nil
}

// Connect initiate the database connection and migrate all the tables
func Connect(connString string) {
	db, err := New(connString)
	if err != nil {
		panic(err)
	}

	// Setting the database connection to use in routes
	DB = db

	fmt.Println("[DATABASE]::CONNECTED")
}

// Migrate migrates all the database tables
func Migrate(tables ...interface{}) error {
	return DB.AutoMigrate(tables...)
}
