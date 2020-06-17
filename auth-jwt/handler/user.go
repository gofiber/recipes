package handler

import (
	"api-fiber-gorm/database"
	"api-fiber-gorm/model"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func validToken(t *jwt.Token, id string) bool {
	n, err := strconv.Atoi(id)
	if err != nil {
		return false
	}

	claims := t.Claims.(jwt.MapClaims)
	uid := int(claims["user_id"].(float64))

	if uid != n {
		return false
	}

	return true
}

func validUser(id string, p string) bool {
	db := database.DB
	var user model.User
	db.First(&user, id)
	if user.Username == "" {
		return false
	}
	if !CheckPasswordHash(p, user.Password) {
		return false
	}
	return true
}

// GetUser get a user
func GetUser(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DB
	var user model.User
	db.Find(&user, id)
	if user.Username == "" {
		c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user found with ID", "data": nil})
		return
	}
	c.JSON(fiber.Map{"status": "success", "message": "Product found", "data": user})
}

// CreateUser new user
func CreateUser(c *fiber.Ctx) {
	type NewUser struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	db := database.DB
	user := new(model.User)
	if err := c.BodyParser(user); err != nil {
		c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
		return
	}

	hash, err := hashPassword(user.Password)
	if err != nil {
		c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't hash password", "data": err})
		return
	}

	user.Password = hash
	if err := db.Create(&user).Error; err != nil {
		c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create user", "data": err})
		return
	}

	newUser := NewUser{
		Email:    user.Email,
		Username: user.Username,
	}

	c.JSON(fiber.Map{"status": "success", "message": "Created user", "data": newUser})
}

// UpdateUser update user
func UpdateUser(c *fiber.Ctx) {
	type UpdateUserInput struct {
		Names string `json:"names"`
	}
	var uui UpdateUserInput
	if err := c.BodyParser(&uui); err != nil {
		c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
		return
	}
	id := c.Params("id")
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, id) {
		c.Status(500).JSON(fiber.Map{"status": "error", "message": "Invalid token id", "data": nil})
		return
	}

	db := database.DB
	var user model.User

	db.First(&user, id)
	user.Names = uui.Names
	if err := db.Save(&user); err != nil {
		c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't update user", "data": err})
		return
	}

	c.JSON(fiber.Map{"status": "success", "message": "User successfully updated", "data": user})
}

// DeleteUser delete user
func DeleteUser(c *fiber.Ctx) {
	type PasswordInput struct {
		Password string `json:"password"`
	}
	var pi PasswordInput
	if err := c.BodyParser(&pi); err != nil {
		c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
		return
	}
	id := c.Params("id")
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, id) {
		c.Status(500).JSON(fiber.Map{"status": "error", "message": "Invalid token id", "data": nil})
		return
	}

	if !validUser(id, pi.Password) {
		c.Status(500).JSON(fiber.Map{"status": "error", "message": "Not valid user", "data": nil})
		return
	}

	db := database.DB
	var user model.User

	db.First(&user, id)

	db.Delete(&user)
	c.JSON(fiber.Map{"status": "success", "message": "User successfully deleted", "data": nil})
}
