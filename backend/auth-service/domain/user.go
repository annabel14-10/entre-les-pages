package domain

import "errors"

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	DateOfBirth string `json:"dob"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"-"`
}

type AuthRequest struct {
	Username string `json:"username"`
	Dob      string `json:"dob"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Message      string `json:"message"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type UserRepository interface {
	GetByEmail(email string) (*User, error)
	Create(user *User) error
	GetByID(id uint) (*User, error)
}

type AuthUsecase interface {
	Login(req *AuthRequest) (*AuthResponse, error)
	Register(req *AuthRequest) (*AuthResponse, error)
	Refresh(refreshToken string) (string, error)
	GetUserByID(id uint) (*User, error)
}

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrInvalidPassword   = errors.New("invalid password")
	ErrUserAlreadyExists = errors.New("user already exists")
)