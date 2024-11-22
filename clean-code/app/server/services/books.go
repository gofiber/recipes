package services

import (
	"context"
	"fmt"

	"app/datasources/database"
	"app/server/domain"
)

type BooksService interface {
	GetBooks(ctx context.Context) ([]domain.Book, error)
	SaveBook(ctx context.Context, newBook domain.Book) error
}

type booksService struct {
	db database.Database
}

func NewBooksService(db database.Database) BooksService {
	return &booksService{db: db}
}

func (s *booksService) GetBooks(ctx context.Context) ([]domain.Book, error) {
	dbRecords, err := s.db.LoadAllBooks(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to load books: %w", err)
	}

	books := make([]domain.Book, 0, len(dbRecords))
	for _, record := range dbRecords {
		books = append(books, domain.Book{
			Title: record.Title,
		})
	}

	return books, nil
}

func (s *booksService) SaveBook(ctx context.Context, book domain.Book) error {
	dbBook := database.Book{
		Title: book.Title,
	}

	err := s.db.CreateBook(ctx, dbBook)
	if err != nil {
		return fmt.Errorf("failed to save book: %w", err)
	}

	return nil
}
