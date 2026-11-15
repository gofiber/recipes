package database

import (
	"fmt"

	"auth-jwt-gorm/config"
	"auth-jwt-gorm/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB gorm connector
var DB *gorm.DB

// ConnectDB connect to db
func ConnectDB() {
	// DB_SSLMODE is documented in the README: disable, require, verify-full, ...
	sslMode := config.Config("DB_SSLMODE")
	if sslMode == "" {
		sslMode = "disable"
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Config("DB_HOST"),
		config.Config("DB_PORT"),
		config.Config("DB_USER"),
		config.Config("DB_PASSWORD"),
		config.Config("DB_NAME"),
		sslMode)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if err = DB.AutoMigrate(&models.Product{}, &models.User{}, &models.RefreshToken{}); err != nil {
		panic("failed to migrate database: " + err.Error())
	}
}
