package services

import (
	"cpf-cnpj-api/internal/domain/models"
	"cpf-cnpj-api/internal/domain/repositories"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepository repositories.UserRepository
	secretKey      string
}

func NewAuthService(userRepository repositories.UserRepository, secretKey string) *AuthService {
	return &AuthService{userRepository: userRepository, secretKey: secretKey}
}

func (s *AuthService) Register(username string, password string, email string) (*models.User, error) {
	existingUser, _ := s.userRepository.FindByUsername(username)
	if existingUser != nil {
		return nil, fmt.Errorf("user already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error on encrypt password: %w", err)
	}

	user := &models.User{
		Username: username,
		Password: string(hash),
		Email:    email,
	}

	if err := s.userRepository.Save(user); err != nil {
		return nil, fmt.Errorf("error on store user: %w", err)
	}

	return user, nil
}

func (s *AuthService) Login(username, password string) (string, error) {
	user, err := s.userRepository.FindByUsername(username)
	if err != nil {
		return "", fmt.Errorf("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", fmt.Errorf("invalid password")
	}

	token, err := s.generateToken(user.ID)
	if err != nil {
		return "", fmt.Errorf("error on generate token: %w", err)
	}

	return token, nil
}

func (s *AuthService) generateToken(userID uint) (string, error) {
	claims := jwt.RegisteredClaims{
		Subject:   fmt.Sprintf("%d", userID),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(s.secretKey))
}

func (s *AuthService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(s.secretKey), nil
	})
}
