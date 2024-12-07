package database

import (
	"context"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDatabase_MemoryDB(t *testing.T) {
	ctx := context.Background()
	db, err := NewDatabase(ctx, "")
	assert.Nil(t, err)
	assert.Equal(t, "*database.memoryDB", reflect.TypeOf(db).String())
}

func TestNewDatabase_PostgresDB(t *testing.T) {
	ctx := context.Background()
	db, err := NewDatabase(ctx, "postgres://localhost:5432")
	assert.Nil(t, err)
	assert.Equal(t, "*database.postgresDB", reflect.TypeOf(db).String())
}

func TestNewDatabase_InvalidDatabaseConfiguration(t *testing.T) {
	ctx := context.Background()
	_, err := NewDatabase(ctx, "invalid")
	assert.ErrorContains(t, err, "unsupported database")
}

func assertBook(t *testing.T, book Book, expectedID int, expected NewBook) {
	assert.Equal(t, expectedID, book.ID)
	assert.Equal(t, expected.Title, book.Title)
}
