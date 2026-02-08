package services

import (
	"errors"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"

	"local-development/testcontainers/app/dal"
	"local-development/testcontainers/app/types"
	"local-development/testcontainers/config/database"
	"local-development/testcontainers/utils"
	"local-development/testcontainers/utils/jwt"
	"local-development/testcontainers/utils/password"
)

// Login service logs in a user
func Login(ctx fiber.Ctx) error {
	b := new(types.LoginDTO)

	if err := utils.ParseBodyAndValidate(ctx, b); err != nil {
		return err
	}

	u := &types.UserResponse{}

	err := dal.FindUserByEmail(database.DB, u, b.Email).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid email or password")
	}

	if err := password.Verify(u.Password, b.Password); err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid email or password")
	}

	t := jwt.Generate(&jwt.TokenPayload{
		ID: u.ID,
	})

	return ctx.JSON(&types.AuthResponse{
		User: u,
		Auth: &types.AccessResponse{
			Token: t,
		},
	})
}

// Signup service creates a user
func Signup(ctx fiber.Ctx) error {
	b := new(types.SignupDTO)

	if err := utils.ParseBodyAndValidate(ctx, b); err != nil {
		return err
	}

	err := dal.FindUserByEmail(database.DB, &struct{ ID string }{}, b.Email).Error

	// If email already exists, return
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.NewError(fiber.StatusConflict, "Email already exists")
	}

	user := &dal.User{
		Name:     b.Name,
		Password: password.Generate(b.Password),
		Email:    b.Email,
	}

	// Create a user, if error return
	if err := dal.CreateUser(database.DB, user); err.Error != nil {
		return fiber.NewError(fiber.StatusConflict, err.Error.Error())
	}

	// Make sure that gorm.Model.ID is uint64, which could happen
	// if the machine compiling the code has multiple versions of gorm.
	uid := uint64(user.ID)

	// generate access token
	t := jwt.Generate(&jwt.TokenPayload{
		ID: uid,
	})

	return ctx.JSON(&types.AuthResponse{
		User: &types.UserResponse{
			ID:    uid,
			Name:  user.Name,
			Email: user.Email,
		},
		Auth: &types.AccessResponse{
			Token: t,
		},
	})
}
