package city

import (
	"context"
	"database/sql"
)

// Queries that we will use.
const (
	QUERY_GET_CITIES  = "SELECT c.id, c.name, c.created, c.modified, u.id, u.name, u.address, u.created, u.modified FROM cities AS c JOIN users AS u ON (c.user = u.id)"
	QUERY_GET_CITY    = "SELECT c.id, c.name, c.created, c.modified, u.id, u.name, u.address, u.created, u.modified FROM cities AS c JOIN users AS u ON (c.user = u.id) WHERE c.id = ?"
	QUERY_CREATE_CITY = "INSERT INTO cities (name, created, modified, user) VALUES (?, ?, ?, ?)"
	QUERY_UPDATE_CITY = "UPDATE cities SET name = ?, modified = ?, user = ? WHERE id = ?"
	QUERY_DELETE_CITY = "DELETE FROM cities WHERE id = ?"
)

// Represents that we will use MariaDB in order to implement the methods.
type mariaDBRepository struct {
	mariadb *sql.DB
}

// Create a new repository with MariaDB as the driver.
func NewCityRepository(mariaDBConnection *sql.DB) CityRepository {
	return &mariaDBRepository{
		mariadb: mariaDBConnection,
	}
}

// Gets all cities in the database.
func (r *mariaDBRepository) GetCities(ctx context.Context) (*[]CityAndUser, error) {
	// Initialize variables.
	var cities []CityAndUser

	// Fetches all cities.
	res, err := r.mariadb.QueryContext(ctx, QUERY_GET_CITIES)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	// Scan all of the cities from the results.
	for res.Next() {
		city := &CityAndUser{}
		err = res.Scan(
			&city.ID,
			&city.Name,
			&city.Created,
			&city.Modified,
			&city.User.ID,
			&city.User.Name,
			&city.User.Address,
			&city.User.Created,
			&city.User.Modified,
		)
		if err != nil && err == sql.ErrNoRows {
			return nil, nil
		}
		if err != nil {
			return nil, err
		}
		cities = append(cities, *city)
	}

	return &cities, nil
}

// Gets a single city in the database.
func (r *mariaDBRepository) GetCity(ctx context.Context, cityID int) (*CityAndUser, error) {
	// Initialize variable.
	city := &CityAndUser{}

	// Prepare SQL to get one city.
	stmt, err := r.mariadb.PrepareContext(ctx, QUERY_GET_CITY)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Get one city and insert it to the 'city' struct.
	// If it's empty, return null.
	err = stmt.QueryRowContext(ctx, cityID).Scan(
		&city.ID,
		&city.Name,
		&city.Created,
		&city.Modified,
		&city.User.ID,
		&city.User.Name,
		&city.User.Address,
		&city.User.Created,
		&city.User.Modified,
	)
	if err != nil && err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	// Return result.
	return city, nil
}

// Creates a single city in the database.
func (r *mariaDBRepository) CreateCity(ctx context.Context, city *City) error {
	// Prepare context to be used.
	stmt, err := r.mariadb.PrepareContext(ctx, QUERY_CREATE_CITY)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Insert one city.
	_, err = stmt.ExecContext(ctx, city.Name, city.Created, city.Modified, city.User)
	if err != nil {
		return err
	}

	// Return empty.
	return nil
}

// Updates a single city in the database.
func (r *mariaDBRepository) UpdateCity(ctx context.Context, cityID int, city *City) error {
	// Prepare context to be used.
	stmt, err := r.mariadb.PrepareContext(ctx, QUERY_UPDATE_CITY)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Update one city.
	_, err = stmt.ExecContext(ctx, city.Name, city.Modified, city.User, cityID)
	if err != nil {
		return err
	}

	// Return empty.
	return nil
}

// Deletes a single city in the database.
func (r *mariaDBRepository) DeleteCity(ctx context.Context, cityID int) error {
	// Prepare context to be used.
	stmt, err := r.mariadb.PrepareContext(ctx, QUERY_DELETE_CITY)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Delete one city.
	_, err = stmt.ExecContext(ctx, cityID)
	if err != nil {
		return err
	}

	// Return empty.
	return nil
}
