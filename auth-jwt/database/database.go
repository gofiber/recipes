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
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Config("DB_HOST"),
		config.Config("DB_PORT"),
		config.Config("DB_USER"),
		config.Config("DB_PASSWORD"),
		config.Config("DB_NAME"))

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&models.Product{}, &models.User{}, &models.RefreshToken{})
}
