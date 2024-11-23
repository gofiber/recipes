package database

import (
	"context"
	"fmt"
	"log"
	"strings"
)

// Book represents a book in the database
type Book struct {
	ID    int
	Title string
}

// NewBook represents a new book to be created to the database
type NewBook struct {
	Title string
}

// Database is an interface for interacting with the database
// With using this the implementation can be changed without affecting the rest of the code.
type Database interface {
	LoadAllBooks(ctx context.Context) ([]Book, error)
	CreateBook(ctx context.Context, newBook NewBook) error
	CloseConnections()
}

// NewDatabase creates a new Database instance
func NewDatabase(ctx context.Context, databaseURL string) (Database, error) {
	if databaseURL == "" {
		log.Printf("Using in-memory database")
		return newMemoryDB(), nil
	} else if strings.HasPrefix(databaseURL, "postgres://") {
		db, err := newPostgresDB(ctx, databaseURL)
		if err != nil {
			return nil, fmt.Errorf("failed to create postgres database: %w", err)
		}
		log.Printf("Using Postgres database")
		return db, nil
	}
	return nil, fmt.Errorf("unsupported database: %s", databaseURL)

}
