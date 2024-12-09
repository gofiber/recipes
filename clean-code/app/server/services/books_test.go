package services

import (
	"context"
	"testing"

	"app/datasources/database"
	"app/server/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetBooks(t *testing.T) {
	mockDB := new(database.DatabaseMock)
	mockDB.On("LoadAllBooks", mock.Anything).Return([]database.Book{{Title: "Title"}}, nil)

	service := NewBooksService(mockDB)
	books, err := service.GetBooks(context.Background())
	assert.Nil(t, err)
	assert.Len(t, books, 1)
}

func TestGetBooks_Fails(t *testing.T) {
	mockDB := new(database.DatabaseMock)
	mockDB.On("LoadAllBooks", mock.Anything).Return(nil, assert.AnError)

	service := NewBooksService(mockDB)
	_, err := service.GetBooks(context.Background())
	assert.NotNil(t, err)
}

func TestSaveBook(t *testing.T) {
	mockDB := new(database.DatabaseMock)
	mockDB.On("CreateBook", mock.Anything, database.NewBook{Title: "Title"}).Return(nil)

	service := NewBooksService(mockDB)
	err := service.SaveBook(context.Background(), domain.Book{Title: "Title"})
	assert.Nil(t, err)
}

func TestSaveBook_Fails(t *testing.T) {
	mockDB := new(database.DatabaseMock)
	mockDB.On("CreateBook", mock.Anything, database.NewBook{Title: "Title"}).Return(assert.AnError)

	service := NewBooksService(mockDB)
	err := service.SaveBook(context.Background(), domain.Book{Title: "Title"})
	assert.NotNil(t, err)
}
