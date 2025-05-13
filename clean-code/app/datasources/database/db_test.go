package database

import (
	"context"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDatabase_MemoryDB(t *testing.T) {
	ctx := context.Background()
	db, err := NewDatabase(ctx, "")
	require.NoError(t, err)
	assert.Equal(t, "*database.DB", reflect.TypeOf(db).String())
	assert.Equal(t, "*database.memoryDB", reflect.TypeOf(db.impl).String())
}

func TestNewDatabase_PostgresDB(t *testing.T) {
	ctx := context.Background()
	db, err := NewDatabase(ctx, "postgres://localhost:5432")
	require.NoError(t, err)
	assert.Equal(t, "*database.DB", reflect.TypeOf(db).String())
	assert.Equal(t, "*database.postgresDB", reflect.TypeOf(db.impl).String())
}

func TestNewDatabase_InvalidDatabaseConfiguration(t *testing.T) {
	ctx := context.Background()
	_, err := NewDatabase(ctx, "invalid")
	assert.ErrorContains(t, err, "unsupported database")
}

func assertBook(t *testing.T, book Book, expectedID int, expected NewBook) {
	t.Helper()
	assert.Equal(t, expectedID, book.ID)
	assert.Equal(t, expected.Title, book.Title)
}
