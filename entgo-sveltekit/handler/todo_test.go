package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	"app/entity"
	"app/entity/enttest"
	"app/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	_ "github.com/mattn/go-sqlite3"
)

// Tests the GetAllTodos handler.
func TestGetAllTodos(t *testing.T) {
	_client := setupDatabase(t)
	_handler := setupTodoHandler(_client)
	defer _client.Close()

	app := fiber.New()

	app.Get("/todos", _handler.GetAllTodos)

	r := httptest.NewRequest("GET", "/todos", nil)

	resp, err := app.Test(r)
	if err != nil {
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	var data []*entity.Todo
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		assert.Fail(t, err.Error())
	}

	fmt.Println("All Todos:", data)
	fmt.Println("Status Code:", resp.StatusCode)
}

// Test the GetTodoByID handler.
func TestGetTodoByID(t *testing.T) {
	_client := setupDatabase(t)
	_handler := setupTodoHandler(_client)
	defer _client.Close()

	app := fiber.New()

	app.Get("/todos/:id", _handler.GetTodoByID)

	r := httptest.NewRequest("GET", "/todos/"+sampleID.String(), nil)

	resp, err := app.Test(r)
	if err != nil {
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	var data entity.Todo
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		assert.Fail(t, err.Error())
	}
	assert.Equal(t, sampleID, data.ID)
	assert.Equal(t, "Sample todo data for unit test.", data.Content)
	assert.Equal(t, true, data.Completed)

	fmt.Println("Todo:", data.String())
	fmt.Println("Status Code:", resp.StatusCode)
}

// Tests the CreateTodo handler.
func TestCreateTodo(t *testing.T) {
	_client := setupDatabase(t)
	_handler := setupTodoHandler(_client)
	defer _client.Close()

	app := fiber.New()

	app.Post("/todos", _handler.CreateTodo)

	body := []byte(`{"content":"Example content"}`)
	r := httptest.NewRequest("POST", "/todos", bytes.NewReader(body))
	r.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(r)
	if err != nil {
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)

	var data entity.Todo
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, "Example content", data.Content)
	assert.Equal(t, false, data.Completed)

	fmt.Println("Todo ID:", data.ID)
	fmt.Println("Status Code:", resp.StatusCode)
}

// Tests the UpdateTodoByID handler.
func TestUpdateTodoByID(t *testing.T) {
	_client := setupDatabase(t)
	_handler := setupTodoHandler(_client)
	defer _client.Close()

	app := fiber.New()

	app.Put("/todos/:id", _handler.UpdateTodoByID)

	body := []byte(`{"content":"Updated content"}`)
	r := httptest.NewRequest("PUT", "/todos/"+sampleID.String(), bytes.NewReader(body))
	r.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(r)
	if err != nil {
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	var data entity.Todo
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		assert.Fail(t, err.Error())
	}
	assert.Equal(t, "Updated content", data.Content)

	fmt.Println("Updated content:", data.Content)
	fmt.Println("Status Code:", resp.StatusCode)
}

// Tests the DeleteTodoByID handler.
func TestDeleteTodoByID(t *testing.T) {
	_client := setupDatabase(t)
	_handler := setupTodoHandler(_client)
	defer _client.Close()

	app := fiber.New()

	app.Delete("/todos/:id", _handler.DeleteTodoByID)

	r := httptest.NewRequest("DELETE", fmt.Sprintf("/todos/%s", sampleID), nil)

	resp, err := app.Test(r)
	if err != nil {
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	var data fiber.Map
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		assert.Fail(t, err.Error())
	}

	fmt.Println("Message:", data["message"])
	fmt.Println("Status Code:", resp.StatusCode)
}

var sampleID = uuid.MustParse("69f48f6b-ed0d-4839-a0eb-15b218b01ca7")

// Sets up a database connection and creates the necessary schema and sample data for
// testing purposes.
func setupDatabase(t *testing.T) (client *entity.Client) {
	client = enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		panic(err)
	}

	if err := setSampleData(ctx, client); err != nil {
		panic(err)
	}
	return client
}

// A new instance of the TodoHandler struct with the given client.
func setupTodoHandler(client *entity.Client) *handler.TodoHandler {
	return handler.NewTodoHandler(client)
}

// Sets sample data for a todo item in a client.
func setSampleData(ctx context.Context, client *entity.Client) error {
	_, err := client.Todo.Create().
		SetID(sampleID).
		SetContent("Sample todo data for unit test.").
		SetCompleted(true).
		Save(ctx)
	return err
}
