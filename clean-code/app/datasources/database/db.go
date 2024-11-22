package database

import (
	"context"
	"log"
	"strings"
)

type Book struct {
	ID    int
	Title string
}

type Database interface {
	LoadAllBooks(ctx context.Context) ([]Book, error)
	CreateBook(ctx context.Context, newBook Book) error
	CloseConnections()
}

func NewDatabase(ctx context.Context, databaseURL string) Database {
	if databaseURL == "" {
		log.Printf("Using in-memory database")
		return newMemoryDB()
	} else if strings.HasPrefix(databaseURL, "postgres://") {
		db, err := newPostgresDB(ctx, databaseURL)
		if err != nil {
			log.Panicf("failed to create postgres database: %v", err)
		}
		log.Printf("Using Postgres database")
		return db
	}
	log.Panicf("unsupported database: %s", databaseURL)
	return nil

}
