package database

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type DatabaseMock struct {
	mock.Mock
}

func (m *DatabaseMock) LoadAllBooks(ctx context.Context) ([]Book, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]Book), args.Error(1)
}

func (m *DatabaseMock) CreateBook(ctx context.Context, newBook NewBook) error {
	args := m.Called(ctx, newBook)
	return args.Error(0)
}

func (m *DatabaseMock) CloseConnections() {
}
