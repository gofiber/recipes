package database

import (
	"log"
	"os"

	"gorm-mysql/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

// ConnectDb opens the MySQL connection and runs auto-migration.
// Set DB_DSN env var to override the default DSN.
// Default DSN: user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
func ConnectDb() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	// NOTE: parseTime=True is required for time.Time fields.
	//       charset=utf8mb4 is required for full UTF-8 support.
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		dsn = "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	log.Println("connected")
	db.AutoMigrate(&models.Book{})
	DBConn = db
}
