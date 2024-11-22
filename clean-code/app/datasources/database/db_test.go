package database

import (
	"context"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDatabase_MemoryDB(t *testing.T) {
	ctx := context.Background()
	db := NewDatabase(ctx, "")
	assert.Equal(t, "*database.memoryDB", reflect.TypeOf(db).String())
}

func TestNewDatabase_PostgresDB(t *testing.T) {
	ctx := context.Background()
	db := NewDatabase(ctx, "postgres://localhost:5432")
	assert.Equal(t, "*database.postgresDB", reflect.TypeOf(db).String())
}

func TestNewDatabase_InvalidDatabaseConfiguration(t *testing.T) {
	ctx := context.Background()
	defer func() {
		assert.NotNil(t, recover())
	}()
	_ = NewDatabase(ctx, "invalid")
}

func assertBook(t *testing.T, book Book, expectedID int, expected Book) {
	assert.Equal(t, expectedID, book.ID)
	assert.Equal(t, expected.Title, book.Title)
}
