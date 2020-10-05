package types

// LoginDTO defined the /login payload
type LoginDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"password"`
}

// SignupDTO defined the /login payload
type SignupDTO struct {
	LoginDTO
	Name string `json:"name" validate:"required,min=3"`
}

// UserResponse todo
type UserResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

// AccessResponse todo
type AccessResponse struct {
	Token string `json:"token"`
}

// AuthResponse todo
type AuthResponse struct {
	User *UserResponse   `json:"user"`
	Auth *AccessResponse `json:"auth"`
}
