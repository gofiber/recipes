package city

import (
	"context"
	"time"
)

// Implementation of the repository in this service.
type cityService struct {
	cityReposiory CityRepository
}

// Create a new 'service' or 'use-case' for 'User' entity.
func NewCityService(r CityRepository) CityService {
	return &cityService{
		cityReposiory: r,
	}
}

// Implementation of 'FetchCities'.
func (s *cityService) FetchCities(ctx context.Context) (*[]CityAndUser, error) {
	return s.cityReposiory.GetCities(ctx)
}

// Implementation of 'FetchCity'.
func (s *cityService) FetchCity(ctx context.Context, cityID int) (*CityAndUser, error) {
	return s.cityReposiory.GetCity(ctx, cityID)
}

// Implementation of 'BuildCity'.
// Our business logic is to set our default variables here.
func (s *cityService) BuildCity(ctx context.Context, city *City, userID int) error {
	city.Created = time.Now().Unix()
	city.Modified = time.Now().Unix()
	city.User = userID
	return s.cityReposiory.CreateCity(ctx, city)
}

// Implementation of 'ModifyCity'.
// Same as above, our business logic is to set our default variables.
func (s *cityService) ModifyCity(ctx context.Context, cityID int, city *City, userID int) error {
	city.Modified = time.Now().Unix()
	city.User = userID
	return s.cityReposiory.UpdateCity(ctx, cityID, city)
}

// Implementation of 'DestroyCity'.
func (s *cityService) DestroyCity(ctx context.Context, cityID int) error {
	return s.cityReposiory.DeleteCity(ctx, cityID)
}
