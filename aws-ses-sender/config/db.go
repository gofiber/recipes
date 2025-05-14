package config

import (
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	dbInstance *gorm.DB
	dbOnce     sync.Once
)

// GetDB Database instance
func GetDB() *gorm.DB {
	dbOnce.Do(func() {
		host := GetEnv("DB_HOST", "localhost")
		port := GetEnv("DB_PORT", "5432")
		user := GetEnv("DB_USER", "postgres")
		password := GetEnv("DB_PASSWORD", "postgres")
		dbname := GetEnv("DB_NAME", "postgres")
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
			host, user, password, dbname, port)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Error),
		})
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalf("Failed to get database instance: %v", err)
		}
		sqlDB.SetMaxOpenConns(10)
		dbInstance = db
	})
	return dbInstance
}

// CloseDB closes the database connection
func CloseDB() error {
	if dbInstance == nil {
		return nil
	}
	sqlDB, err := dbInstance.DB()
	if err != nil {
		return fmt.Errorf("failed to get database instance: %v", err)
	}
	return sqlDB.Close()
}
