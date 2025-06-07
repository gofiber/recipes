package dal_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	tc "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"

	"local-development/testcontainers/app/dal"
	"local-development/testcontainers/config/database"
)

func TestTodos(t *testing.T) {
	pgCtr, err := postgres.Run(
		context.Background(),
		"postgres:16",
		postgres.WithDatabase("todos"),
		postgres.WithUsername("todos"),
		postgres.WithPassword("todos"),
		postgres.BasicWaitStrategies(),
	)
	tc.CleanupContainer(t, pgCtr)
	require.NoError(t, err)

	connString, err := pgCtr.ConnectionString(context.Background())
	require.NoError(t, err)

	db, err := database.New(connString)
	require.NoError(t, err)

	err = db.AutoMigrate(&dal.User{}, &dal.Todo{})
	require.NoError(t, err)

	user := &dal.User{
		Name:     "John Doe",
		Email:    "john.doe@example.com",
		Password: "password",
	}
	result := dal.CreateUser(db, user)
	require.NoError(t, result.Error)

	// Make sure that gorm.Model.ID is uint64, which could happen
	// if the machine compiling the code has multiple versions of gorm.
	uid := uint64(user.ID)

	result = result.Scan(&user)
	require.NoError(t, result.Error)

	t.Run("create", func(t *testing.T) {
		result := dal.CreateTodo(db, &dal.Todo{
			Task: "Buy groceries",
			User: &uid,
		})
		require.NoError(t, result.Error)

		todo1 := &dal.Todo{}
		result = result.Scan(todo1)
		require.NoError(t, result.Error)

		// create a second todo
		result = dal.CreateTodo(db, &dal.Todo{
			Task: "Clean the swimming pool",
			User: &uid,
		})
		require.NoError(t, result.Error)

		t.Run("find", func(t *testing.T) {
			result := dal.FindTodo(db, &dal.Todo{}, "todos.task = ?", "Buy groceries")
			require.NoError(t, result.Error)
			require.Equal(t, int64(1), result.RowsAffected)

			t.Run("todos-by-user", func(t *testing.T) {
				result := dal.FindTodosByUser(db, &[]dal.Todo{}, uid)
				require.NoError(t, result.Error)
				require.Equal(t, int64(2), result.RowsAffected)
			})

			t.Run("todo-by-user", func(t *testing.T) {
				result := dal.FindTodoByUser(db, &[]dal.Todo{}, todo1.ID, uid)
				require.NoError(t, result.Error)
				require.Equal(t, int64(1), result.RowsAffected)

				todo := &dal.Todo{}
				result = result.Scan(todo)
				require.NoError(t, result.Error)
				require.Equal(t, todo1.ID, todo.ID)
			})
		})

		t.Run("update", func(t *testing.T) {
			result := dal.FindTodo(db, &dal.Todo{}, "todos.task = ?", "Buy a new car")
			require.Error(t, result.Error)
			require.Zero(t, result.RowsAffected)

			result = dal.UpdateTodo(db, todo1.ID, &uid, &dal.Todo{Task: "Buy a new car"})
			require.NoError(t, result.Error)

			result = dal.FindTodo(db, &dal.Todo{}, "todos.task = ?", "Buy a new car")
			require.NoError(t, result.Error)
			require.Equal(t, int64(1), result.RowsAffected)
		})

		t.Run("delete", func(t *testing.T) {
			result := dal.CreateTodo(db, &dal.Todo{
				Task: "do the house cleaning",
				User: &uid,
			})
			require.NoError(t, result.Error)

			todo := &dal.Todo{}
			result = result.Scan(todo)
			require.NoError(t, result.Error)

			result = dal.DeleteTodo(db, todo.ID, &uid)
			require.NoError(t, result.Error)
			require.Equal(t, int64(1), result.RowsAffected)
		})
	})
}
