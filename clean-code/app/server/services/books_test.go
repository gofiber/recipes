package services

import (
	"context"
	"testing"

	"app/datasources/database"
	"app/server/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestGetBooks(t *testing.T) {
	mockDB := new(database.MockDatabase)
	mockDB.On("LoadAllBooks", mock.Anything).Return([]database.Book{{Title: "Title"}}, nil)

	service := NewBooksService(mockDB)
	books, err := service.GetBooks(context.Background())
	require.NoError(t, err)
	assert.Len(t, books, 1)
}

func TestGetBooks_Fails(t *testing.T) {
	mockDB := new(database.MockDatabase)
	mockDB.On("LoadAllBooks", mock.Anything).Return(nil, assert.AnError)

	service := NewBooksService(mockDB)
	_, err := service.GetBooks(context.Background())
	assert.Error(t, err)
}

func TestSaveBook(t *testing.T) {
	mockDB := new(database.MockDatabase)
	mockDB.On("CreateBook", mock.Anything, database.NewBook{Title: "Title"}).Return(nil)

	service := NewBooksService(mockDB)
	err := service.SaveBook(context.Background(), domain.Book{Title: "Title"})
	require.NoError(t, err)
}

func TestSaveBook_Fails(t *testing.T) {
	mockDB := new(database.MockDatabase)
	mockDB.On("CreateBook", mock.Anything, database.NewBook{Title: "Title"}).Return(assert.AnError)

	service := NewBooksService(mockDB)
	err := service.SaveBook(context.Background(), domain.Book{Title: "Title"})
	assert.Error(t, err)
}
