package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User struct
type User struct {
	gorm.Model
	Id        string     `gorm:"primaryKey" json:"id"`
	Username  string     `gorm:"uniqueIndex;not null" json:"username"`
	Email     string     `gorm:"uniqueIndex;not null" json:"email"`
	Password  string     `gorm:"not null" json:"password"`
	Names     string     `json:"names"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	LastLogin *time.Time `json:"last_login"`
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// CreateUser adds a new user to the database
func (r *UserRepository) CreateUser(email, username, passwordHash string) (*User, error) {
	user := &User{
		Id:        uuid.New().String(),
		Email:     email,
		Username:  username,
		Password:  passwordHash,
		CreatedAt: time.Now(),
	}

	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserByEmail retrieves a user by their email address
func (r *UserRepository) GetUserByEmail(email string) (*User, error) {
	var user User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByID retrieves a user by their ID
func (r *UserRepository) GetUserByID(id string) (*User, error) {
	var user User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) DeleteUser(id string) (error){
	var user User
	if err := r.db.Where("id = ?", id).Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) UpdateUser(id string,updateUser User) (*User, error) {
	var user User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	user.Email = updateUser.Email
	user.Username = updateUser.Username
	user.Password = updateUser.Password
	user.Names = updateUser.Names
	user.UpdatedAt = time.Now()
	if err := r.db.Save(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}


