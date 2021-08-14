package city

import (
	"context"
	"docker-mariadb-clean-arch/internal/user"
)

// Represents 'cities' object.
type City struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Created  int64  `json:"created"`
	Modified int64  `json:"modified"`
	User     int    `json:"user"`
}

// Represents our 1-to-many relationship, User to Cities.
// In other words, a single user can have many cities.
// This is used for the presentation layer only.
type CityAndUser struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Created  int64     `json:"created"`
	Modified int64     `json:"modified"`
	User     user.User `json:"user"`
}

// Our repository will implement these methods.
type CityRepository interface {
	GetCities(ctx context.Context) (*[]CityAndUser, error)
	GetCity(ctx context.Context, cityID int) (*CityAndUser, error)
	CreateCity(ctx context.Context, city *City) error
	UpdateCity(ctx context.Context, cityID int, city *City) error
	DeleteCity(ctx context.Context, cityID int) error
}

// Our use-case or service will implement these methods.
// Method names does not matter - I intentionally changed the names here so they are different from the repository.
// Usually, 'services' should be as close as possible to a business use-case / problem.
// Some methods will also take 'userID' as an argument - because we need it to fill the 'User' value in the struct.
type CityService interface {
	FetchCities(ctx context.Context) (*[]CityAndUser, error)
	FetchCity(ctx context.Context, cityID int) (*CityAndUser, error)
	BuildCity(ctx context.Context, city *City, userID int) error
	ModifyCity(ctx context.Context, cityID int, city *City, userID int) error
	DestroyCity(ctx context.Context, cityID int) error
}
