package database

import (
	"context"
	"fmt"
	"log/slog"
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

// Database defines the interface for interacting with the book database.
// Using this interface allows changing the implementation without affecting the rest of the code.
type Database interface {
	// LoadAllBooks retrieves all books from the database.
	LoadAllBooks(ctx context.Context) ([]Book, error)

	// CreateBook adds a new book to the database.
	CreateBook(ctx context.Context, newBook NewBook) error

	// CloseConnections closes all open connections to the database.
	CloseConnections()
}

// NewDatabase creates a new Database instance
func NewDatabase(ctx context.Context, databaseURL string) (Database, error) {
	if databaseURL == "" {
		slog.Info("Using in-memory database")
		return newMemoryDB(), nil
	} else if strings.HasPrefix(databaseURL, "postgres://") {
		db, err := newPostgresDB(ctx, databaseURL)
		if err != nil {
			return nil, fmt.Errorf("failed to create postgres database: %w", err)
		}
		slog.Info("Using Postgres database")
		return db, nil
	}
	return nil, fmt.Errorf("unsupported database: %s", databaseURL)

}
