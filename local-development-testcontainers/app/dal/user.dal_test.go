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

func TestUsers(t *testing.T) {
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

	err = db.AutoMigrate(&dal.User{})
	require.NoError(t, err)

	t.Run("create", func(t *testing.T) {
		user := &dal.User{
			Name:     "John Doe",
			Email:    "john.doe@example.com",
			Password: "password",
		}
		result := dal.CreateUser(db, user)
		require.NoError(t, result.Error)

		result = result.Scan(&user)
		require.NoError(t, result.Error)

		uID := user.ID

		t.Run("find", func(t *testing.T) {
			result := dal.FindUser(db, &dal.User{}, "id = ?", uID)
			require.NoError(t, result.Error)
			require.Equal(t, int64(1), result.RowsAffected)

			user := &dal.User{}
			result = result.Scan(user)
			require.NoError(t, result.Error)
		})

		t.Run("find-by-email", func(t *testing.T) {
			result := dal.FindUserByEmail(db, &dal.User{}, user.Email)
			require.NoError(t, result.Error)
			require.Equal(t, int64(1), result.RowsAffected)

			user := &dal.User{}
			result = result.Scan(user)
			require.NoError(t, result.Error)
			require.Equal(t, user.ID, uID)
		})
	})
}
