package database

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strings"
)

var ErrUnsupportedDatabase = errors.New("unsupported database URL")

// Book represents a book in the database.
type Book struct {
	ID    int
	Title string
}

// NewBook represents a new book to be created to the database.
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

// NewDatabase creates a new Database instance.
func NewDatabase(ctx context.Context, databaseURL string) (*DB, error) {
	if databaseURL == "" {
		db := newMemoryDB()
		slog.Info("Using in-memory database implementation")
		return &DB{impl: db}, nil
	}

	if strings.HasPrefix(databaseURL, "postgres://") {
		db, err := newPostgresDB(ctx, databaseURL)
		if err != nil {
			return nil, fmt.Errorf("failed to initialize PostgreSQL database connection: %w", err)
		}
		slog.Info("Using PostgreSQL database implementation")
		return &DB{impl: db}, nil
	}

	return nil, fmt.Errorf("%w: %s", ErrUnsupportedDatabase, databaseURL)
}

type DB struct {
	impl Database
}

func (db *DB) LoadAllBooks(ctx context.Context) ([]Book, error) {
	return db.impl.LoadAllBooks(ctx)
}

func (db *DB) CreateBook(ctx context.Context, newBook NewBook) error {
	return db.impl.CreateBook(ctx, newBook)
}

func (db *DB) CloseConnections() {
	db.impl.CloseConnections()
}
