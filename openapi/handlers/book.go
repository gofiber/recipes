package handlers

import (
	"context"
	"math"

	"openapi/database"
	"openapi/models"

	"github.com/danielgtaylor/huma/v2"
	"gorm.io/gorm"
)

// idExists reports whether id can be the id of a book. Ids are signed 64-bit integers in the
// database, so a larger one cannot exist and cannot be bound as a query argument either.
func idExists(id uint) bool {
	return uint64(id) <= math.MaxInt64
}

type GetAllBooksResponse struct {
	Body []models.Book `json:"body"`
}

type GetBookByIDResponse struct {
	Body models.Book `json:"body"`
}

type RegisterBookResponse struct {
	Body struct {
		ID uint `json:"id"`
	}
}

type RegisterBookDto struct {
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
}

// Response schema for PATCH & DELETE
type EmptyResponse struct{}

func GetAllBooks(ctx context.Context, _ *struct{}) (*GetAllBooksResponse, error) {
	db := database.DBConn

	var books []models.Book
	if err := db.Find(&books).Error; err != nil {
		return nil, err
	}
	return &GetAllBooksResponse{Body: books}, nil
}

// ID is a uint rather than a string on purpose: Huma rejects non-integers with a 422
// before the handler runs, and GORM would otherwise treat a non-numeric string passed to
// First() as a raw WHERE clause (SQL injection).
func GetBookByID(ctx context.Context, params *struct {
	ID uint `path:"id"`
}) (*GetBookByIDResponse, error) {
	if !idExists(params.ID) {
		return nil, huma.Error404NotFound("book not found")
	}
	db := database.DBConn

	book := new(models.Book)
	if err := db.First(&book, params.ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, huma.Error404NotFound("book not found")
		}
		return nil, err
	}
	return &GetBookByIDResponse{Body: *book}, nil
}

func RegisterBook(ctx context.Context, input *struct{ Body RegisterBookDto }) (*RegisterBookResponse, error) {
	db := database.DBConn

	book := models.Book{
		Title:     input.Body.Title,
		Author:    input.Body.Author,
		Publisher: input.Body.Publisher,
	}
	if err := db.Create(&book).Error; err != nil {
		return nil, err
	}

	return &RegisterBookResponse{Body: struct {
		ID uint `json:"id"`
	}{ID: book.ID}}, nil
}

func UpdateBook(ctx context.Context, input *struct {
	ID   uint `path:"id"`
	Body RegisterBookDto
}) (*EmptyResponse, error) {
	if !idExists(input.ID) {
		return nil, huma.Error404NotFound("book not found")
	}
	db := database.DBConn

	book := new(models.Book)
	if err := db.First(&book, input.ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, huma.Error404NotFound("book not found")
		}
		return nil, err
	}

	book.Title = input.Body.Title
	book.Author = input.Body.Author
	book.Publisher = input.Body.Publisher
	if err := db.Save(&book).Error; err != nil {
		return nil, err
	}

	return &EmptyResponse{}, nil
}

func DeleteBook(ctx context.Context, params *struct {
	ID uint `path:"id"`
}) (*EmptyResponse, error) {
	if !idExists(params.ID) {
		return nil, huma.Error404NotFound("book not found")
	}
	db := database.DBConn

	book := new(models.Book)
	if err := db.First(&book, params.ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, huma.Error404NotFound("book not found")
		}
		return nil, err
	}

	if err := db.Delete(&book).Error; err != nil {
		return nil, err
	}

	return &EmptyResponse{}, nil
}
