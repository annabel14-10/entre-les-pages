package usecase

import (
	"auth-service/domain"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type authUsecase struct {
	userRepo domain.UserRepository
}

func NewAuthUsecase(repo domain.UserRepository) domain.AuthUsecase {
	return &authUsecase{userRepo: repo}
}

func (a *authUsecase) Register(req *domain.AuthRequest) (*domain.AuthResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	user := domain.User{
		Username:    req.Username,
		DateOfBirth: req.Dob,
		Email:       req.Email,
		Password:    string(hashedPassword),
	}

	if err := a.userRepo.Create(&user); err != nil {
		return nil, errors.New("email already exists")
	}

	accessToken, refreshToken, err := a.generateTokens(user.ID)
	if err != nil {
		return nil, errors.New("failed to generate tokens")
	}

	return &domain.AuthResponse{
		Message:      "Registration successful",
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (a *authUsecase) Login(req *domain.AuthRequest) (*domain.AuthResponse, error) {
	user, err := a.userRepo.GetByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	accessToken, refreshToken, err := a.generateTokens(user.ID)
	if err != nil {
		return nil, errors.New("failed to generate tokens")
	}

	return &domain.AuthResponse{
		Message:      "Login successful",
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (a *authUsecase) generateTokens(userID uint) (string, string, error) {
	secret := []byte(os.Getenv("JWT_SECRET"))

	accessClaims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessString, err := accessToken.SignedString(secret)
	if err != nil {
		return "", "", err
	}

	refreshClaims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(7 * 24 * time.Hour).Unix(),
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshString, err := refreshToken.SignedString(secret)
	if err != nil {
		return "", "", err
	}

	return accessString, refreshString, nil
}

func (a *authUsecase) Refresh(refreshTokenString string) (string, error) {
	secret := []byte(os.Getenv("JWT_SECRET"))

	token, err := jwt.Parse(refreshTokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secret, nil
	})

	if err != nil || !token.Valid {
		return "", errors.New("invalid or expired refresh token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid token payload")
	}

	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		return "", errors.New("user_id not found in token")
	}
	userID := uint(userIDFloat)

	accessClaims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
	}
	newAccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)

	newAccessString, err := newAccessToken.SignedString(secret)
	if err != nil {
		return "", errors.New("failed to forge new access token")
	}

	return newAccessString, nil
}

func (a *authUsecase) GetUserByID(id uint) (*domain.User, error) {
	return a.userRepo.GetByID(id)
}