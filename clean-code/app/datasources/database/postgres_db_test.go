package database

import (
	"context"
	"testing"

	"github.com/pashagolub/pgxmock/v4"
	"github.com/stretchr/testify/assert"
)

func TestPostgresDB_GetBooks(t *testing.T) {
	mockPool, err := pgxmock.NewPool()
	assert.Nil(t, err)

	mockPool.ExpectQuery("SELECT id, title FROM books").
		WillReturnRows(pgxmock.NewRows([]string{"id", "title"}).
			AddRow(1, "book1"))

	db := postgresDB{
		pool: mockPool,
	}
	result, err := db.LoadAllBooks(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, 1, len(result))
	assertBook(t, result[0], 1, NewBook{Title: "book1"})

	assert.Nil(t, mockPool.ExpectationsWereMet())
}

func TestPostgresDB_GetBooks_Fail(t *testing.T) {
	mockPool, err := pgxmock.NewPool()
	assert.Nil(t, err)

	mockPool.ExpectQuery("SELECT id, title FROM books").
		WillReturnError(assert.AnError)

	db := postgresDB{
		pool: mockPool,
	}
	result, err := db.LoadAllBooks(context.Background())
	assert.Nil(t, result)
	assert.ErrorContains(t, err, "failed to query books table")

	assert.Nil(t, mockPool.ExpectationsWereMet())
}

func TestPostgresDB_CreateBook(t *testing.T) {
	mockPool, err := pgxmock.NewPool()
	assert.Nil(t, err)

	mockPool.ExpectExec("INSERT INTO books").
		WithArgs("book1").
		WillReturnResult(pgxmock.NewResult("INSERT", 1))

	db := postgresDB{
		pool: mockPool,
	}
	err = db.CreateBook(context.Background(), NewBook{Title: "book1"})
	assert.Nil(t, err)

	assert.Nil(t, mockPool.ExpectationsWereMet())
}

func TestPostgresDB_CreateBook_Fail(t *testing.T) {
	mockPool, err := pgxmock.NewPool()
	assert.Nil(t, err)

	mockPool.ExpectExec("INSERT INTO books").
		WithArgs("book1").
		WillReturnError(assert.AnError)

	db := postgresDB{
		pool: mockPool,
	}
	err = db.CreateBook(context.Background(), NewBook{Title: "book1"})
	assert.ErrorContains(t, err, "failed to insert book")

	assert.Nil(t, mockPool.ExpectationsWereMet())
}
