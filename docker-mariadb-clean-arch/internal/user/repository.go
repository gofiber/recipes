package user

import (
	"context"
	"database/sql"
)

// Queries that we will use.
const (
	QUERY_GET_USERS   = "SELECT * FROM users"
	QUERY_GET_USER    = "SELECT * FROM users WHERE id = ?"
	QUERY_CREATE_USER = "INSERT INTO users (name, address, created, modified) VALUES (?, ?, ?, ?)"
	QUERY_UPDATE_USER = "UPDATE users SET name = ?, address = ?, modified = ? WHERE id = ?"
	QUERY_DELETE_USER = "DELETE FROM users WHERE id = ?"
)

// Represents that we will use MariaDB in order to implement the methods.
type mariaDBRepository struct {
	mariadb *sql.DB
}

// Create a new repository with MariaDB as the driver.
func NewUserRepository(mariaDBConnection *sql.DB) UserRepository {
	return &mariaDBRepository{
		mariadb: mariaDBConnection,
	}
}

// Gets all users in the database.
func (r *mariaDBRepository) GetUsers(ctx context.Context) (*[]User, error) {
	// Initialize variables.
	var users []User

	// Get all users.
	res, err := r.mariadb.QueryContext(ctx, QUERY_GET_USERS)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	// Scan all of the results to the 'users' array.
	// If it's empty, return null.
	for res.Next() {
		user := &User{}
		err = res.Scan(&user.ID, &user.Name, &user.Address, &user.Created, &user.Modified)
		if err != nil && err == sql.ErrNoRows {
			return nil, nil
		}
		if err != nil {
			return nil, err
		}
		users = append(users, *user)
	}

	// Return all of our users.
	return &users, nil
}

// Gets a single user in the database.
func (r *mariaDBRepository) GetUser(ctx context.Context, userID int) (*User, error) {
	// Initialize variable.
	user := &User{}

	// Prepare SQL to get one user.
	stmt, err := r.mariadb.PrepareContext(ctx, QUERY_GET_USER)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Get one user and insert it to the 'user' struct.
	// If it's empty, return null.
	err = stmt.QueryRowContext(ctx, userID).Scan(&user.ID, &user.Name, &user.Address, &user.Created, &user.Modified)
	if err != nil && err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	// Return result.
	return user, nil
}

// Creates a single user in the database.
func (r *mariaDBRepository) CreateUser(ctx context.Context, user *User) error {
	// Prepare context to be used.
	stmt, err := r.mariadb.PrepareContext(ctx, QUERY_CREATE_USER)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Insert one user.
	_, err = stmt.ExecContext(ctx, user.Name, user.Address, user.Created, user.Modified)
	if err != nil {
		return err
	}

	// Return empty.
	return nil
}

// Updates a single user in the database.
func (r *mariaDBRepository) UpdateUser(ctx context.Context, userID int, user *User) error {
	// Prepare context to be used.
	stmt, err := r.mariadb.PrepareContext(ctx, QUERY_UPDATE_USER)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Update one user.
	_, err = stmt.ExecContext(ctx, user.Name, user.Address, user.Modified, userID)
	if err != nil {
		return err
	}

	// Return empty.
	return nil
}

// Deletes a single user in the database.
func (r *mariaDBRepository) DeleteUser(ctx context.Context, userID int) error {
	// Prepare context to be used.
	stmt, err := r.mariadb.PrepareContext(ctx, QUERY_DELETE_USER)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Delete one user.
	_, err = stmt.ExecContext(ctx, userID)
	if err != nil {
		return err
	}

	// Return empty.
	return nil
}
