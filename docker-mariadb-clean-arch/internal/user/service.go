package user

import (
	"context"
	"time"
)

// Implementation of the repository in this service.
type userService struct {
	userRepository UserRepository
}

// Create a new 'service' or 'use-case' for 'User' entity.
func NewUserService(r UserRepository) UserService {
	return &userService{
		userRepository: r,
	}
}

// Implementation of 'GetUsers'.
func (s *userService) GetUsers(ctx context.Context) (*[]User, error) {
	return s.userRepository.GetUsers(ctx)
}

// Implementation of 'GetUser'.
func (s *userService) GetUser(ctx context.Context, userID int) (*User, error) {
	return s.userRepository.GetUser(ctx, userID)
}

// Implementation of 'CreateUser'.
func (s *userService) CreateUser(ctx context.Context, user *User) error {
	// Set default value of 'Created' and 'Modified'.
	user.Created = time.Now().Unix()
	user.Modified = time.Now().Unix()

	// Pass to the repository layer.
	return s.userRepository.CreateUser(ctx, user)
}

// Implementation of 'UpdateUser'.
func (s *userService) UpdateUser(ctx context.Context, userID int, user *User) error {
	// Set value for 'Modified' attribute.
	user.Modified = time.Now().Unix()

	// Pass to the repository layer.
	return s.userRepository.UpdateUser(ctx, userID, user)
}

// Implementation of 'DeleteUser'.
func (s *userService) DeleteUser(ctx context.Context, userID int) error {
	return s.userRepository.DeleteUser(ctx, userID)
}
