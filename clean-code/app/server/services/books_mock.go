package services

import (
	"context"

	"app/server/domain"

	"github.com/stretchr/testify/mock"
)

type MockBooksService struct {
	mock.Mock
}

func (m *MockBooksService) GetBooks(ctx context.Context) ([]domain.Book, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]domain.Book), args.Error(1)
}

func (m *MockBooksService) SaveBook(ctx context.Context, newBook domain.Book) error {
	args := m.Called(ctx, newBook)
	return args.Error(0)
}
