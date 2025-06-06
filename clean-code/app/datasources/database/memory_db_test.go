package database

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMemoryDB_LoadBooks(t *testing.T) {
	db := newMemoryDB()
	books, err := db.LoadAllBooks(context.Background())
	require.NoError(t, err)
	assert.Empty(t, books)
}

func TestMemoryDB_SaveBook(t *testing.T) {
	db := newMemoryDB()
	newBook := NewBook{Title: "Title"}
	err := db.CreateBook(context.Background(), newBook)
	require.NoError(t, err)

	books, err := db.LoadAllBooks(context.Background())
	require.NoError(t, err)
	assert.Len(t, books, 1)
	assertBook(t, books[0], 0, newBook)
}

func TestMemoryDB_SaveBookMultiple(t *testing.T) {
	db := newMemoryDB()
	newBook1 := NewBook{Title: "Title1"}
	err := db.CreateBook(context.Background(), newBook1)
	require.NoError(t, err)

	newBook2 := NewBook{Title: "Title2"}
	err = db.CreateBook(context.Background(), newBook2)
	require.NoError(t, err)

	books, err := db.LoadAllBooks(context.Background())
	require.NoError(t, err)
	assert.Len(t, books, 2)
	assertBook(t, books[0], 0, newBook1)
	assertBook(t, books[1], 1, newBook2)
}
