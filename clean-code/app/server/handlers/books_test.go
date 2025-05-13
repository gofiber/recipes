package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"app/server/domain"
	"app/server/services"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

const booksRoute = "/api/v1/books"

func TestGetBooks(t *testing.T) {
	mockService := new(services.MockBooksService)
	mockService.On("GetBooks", mock.Anything).Return([]domain.Book{{Title: "Title"}}, nil)

	app := fiber.New()
	app.Get(booksRoute, GetBooks(mockService))

	resp, err := app.Test(httptest.NewRequest(http.MethodGet, booksRoute, nil))
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	defer resp.Body.Close()

	body := bodyFromResponse[domain.BooksResponse](t, resp)
	assert.Len(t, body.Books, 1)
}

func TestGetBooks_ServiceFails(t *testing.T) {
	mockService := new(services.MockBooksService)
	mockService.On("GetBooks", mock.Anything).Return(nil, assert.AnError)

	app := fiber.New()
	app.Get(booksRoute, GetBooks(mockService))

	resp, err := app.Test(httptest.NewRequest(http.MethodGet, booksRoute, nil))
	require.NoError(t, err)
	assert.Equal(t, 500, resp.StatusCode)
	defer resp.Body.Close()

	body := bodyFromResponse[domain.ErrorResponse](t, resp)
	assert.Equal(t, "internal error", body.Error)
}

func TestAddBook(t *testing.T) {
	mockService := new(services.MockBooksService)
	mockService.On("SaveBook", mock.Anything, domain.Book{Title: "Title"}).Return(nil)

	app := fiber.New()
	app.Post(booksRoute, AddBook(mockService))

	resp, err := app.Test(postRequest(booksRoute, `{"title":"Title"}`))
	require.NoError(t, err)
	assert.Equal(t, 201, resp.StatusCode)
	defer resp.Body.Close()
}

func TestAddBook_InvalidRequest(t *testing.T) {
	mockService := new(services.MockBooksService)

	app := fiber.New()
	app.Post(booksRoute, AddBook(mockService))

	resp, err := app.Test(httptest.NewRequest(http.MethodPost, booksRoute, nil))
	require.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode)
	defer resp.Body.Close()

	body := bodyFromResponse[domain.ErrorResponse](t, resp)
	assert.Equal(t, "invalid request", body.Error)
}

func TestAddBook_ServiceFails(t *testing.T) {
	mockService := new(services.MockBooksService)
	mockService.On("SaveBook", mock.Anything, domain.Book{Title: "Title"}).Return(assert.AnError)

	app := fiber.New()
	app.Post(booksRoute, AddBook(mockService))

	resp, err := app.Test(postRequest(booksRoute, `{"title":"Title"}`))
	require.NoError(t, err)
	assert.Equal(t, 500, resp.StatusCode)
	defer resp.Body.Close()

	body := bodyFromResponse[domain.ErrorResponse](t, resp)
	assert.Equal(t, "internal error", body.Error)
}

func postRequest(url string, body string) *http.Request {
	req := httptest.NewRequest(http.MethodPost, url, bytes.NewBufferString(body))
	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return req
}

//nolint:ireturn
func bodyFromResponse[T any](t *testing.T, resp *http.Response) T {
	t.Helper()
	var body T
	err := json.NewDecoder(resp.Body).Decode(&body)
	require.NoError(t, err, "failed to decode response body")
	return body
}
