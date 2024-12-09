package database

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemoryDB_LoadBooks(t *testing.T) {
	db := newMemoryDB()
	books, err := db.LoadAllBooks(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, 0, len(books))
}

func TestMemoryDB_SaveBook(t *testing.T) {
	db := newMemoryDB()
	newBook := NewBook{Title: "Title"}
	err := db.CreateBook(context.Background(), newBook)
	assert.Nil(t, err)

	books, err := db.LoadAllBooks(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, 1, len(books))
	assertBook(t, books[0], 0, newBook)
}

func TestMemoryDB_SaveBookMultiple(t *testing.T) {
	db := newMemoryDB()
	newBook1 := NewBook{Title: "Title1"}
	err := db.CreateBook(context.Background(), newBook1)
	assert.Nil(t, err)

	newBook2 := NewBook{Title: "Title2"}
	err = db.CreateBook(context.Background(), newBook2)
	assert.Nil(t, err)

	books, err := db.LoadAllBooks(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, 2, len(books))
	assertBook(t, books[0], 0, newBook1)
	assertBook(t, books[1], 1, newBook2)
}
