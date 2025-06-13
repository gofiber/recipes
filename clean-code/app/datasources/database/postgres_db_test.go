package database

import (
	"context"
	"testing"

	"github.com/pashagolub/pgxmock/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPostgresDB_GetBooks(t *testing.T) {
	mockPool, err := pgxmock.NewPool()
	require.NoError(t, err)

	mockPool.ExpectQuery("SELECT id, title FROM books").
		WillReturnRows(pgxmock.NewRows([]string{"id", "title"}).
			AddRow(1, "book1"))

	db := postgresDB{
		pool: mockPool,
	}
	result, err := db.LoadAllBooks(context.Background())
	require.NoError(t, err)
	assert.Len(t, result, 1)
	assertBook(t, result[0], 1, NewBook{Title: "book1"})

	require.NoError(t, mockPool.ExpectationsWereMet())
}

func TestPostgresDB_GetBooks_Fail(t *testing.T) {
	mockPool, err := pgxmock.NewPool()
	require.NoError(t, err)

	mockPool.ExpectQuery("SELECT id, title FROM books").
		WillReturnError(assert.AnError)

	db := postgresDB{
		pool: mockPool,
	}
	result, err := db.LoadAllBooks(context.Background())
	assert.Nil(t, result)
	require.ErrorContains(t, err, "failed to query books table")

	require.NoError(t, mockPool.ExpectationsWereMet())
}

func TestPostgresDB_CreateBook(t *testing.T) {
	mockPool, err := pgxmock.NewPool()
	require.NoError(t, err)

	mockPool.ExpectExec("INSERT INTO books").
		WithArgs("book1").
		WillReturnResult(pgxmock.NewResult("INSERT", 1))

	db := postgresDB{
		pool: mockPool,
	}
	err = db.CreateBook(context.Background(), NewBook{Title: "book1"})
	require.NoError(t, err)

	require.NoError(t, mockPool.ExpectationsWereMet())
}

func TestPostgresDB_CreateBook_Fail(t *testing.T) {
	mockPool, err := pgxmock.NewPool()
	require.NoError(t, err)

	mockPool.ExpectExec("INSERT INTO books").
		WithArgs("book1").
		WillReturnError(assert.AnError)

	db := postgresDB{
		pool: mockPool,
	}
	err = db.CreateBook(context.Background(), NewBook{Title: "book1"})
	require.ErrorContains(t, err, "failed to insert book")

	require.NoError(t, mockPool.ExpectationsWereMet())
}
