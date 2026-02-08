package database

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockDatabase struct {
	mock.Mock
}

func (m *MockDatabase) LoadAllBooks(ctx context.Context) ([]Book, error) {
	args := m.Called(ctx)
	if books, ok := args.Get(0).([]Book); ok {
		return books, args.Error(1)
	}
	return []Book{}, args.Error(1)
}

func (m *MockDatabase) CreateBook(ctx context.Context, newBook NewBook) error {
	args := m.Called(ctx, newBook)
	return args.Error(0)
}

func (m *MockDatabase) CloseConnections() {
}
