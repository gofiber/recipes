package infrastructure

import (
	"database/sql"
	"os"
	"time"
)

// This function is used to connect to MariaDB.
func ConnectToMariaDB() (*sql.DB, error) {
	// Read DSN from environment; fall back to local default for development.
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		dsn = "root:@tcp(mariadb:3306)/fiber_dmca"
	}

	// Connect to MariaDB.
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Set up important parts as was told by the documentation.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	// Return our database instance.
	return db, nil
}
