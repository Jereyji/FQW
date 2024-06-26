package service

import (
	"errors"
	"time"

	todo "github.com/Jereyji/FQW.git"
	"github.com/Jereyji/FQW.git/internal/repository"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

const (
	signingkey = "8109612acafb4ae0ff34f5f1fa549577f4ca3a4a294f559498c111cc7d92973e5dde4eb64f086b49e063708705338f29b662047c09c850f5bb21da65f37036b4"
	tokenTTL = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json: "user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	enc, err := encryptedString(user.Password)
	if err != nil {
		return 0, err
	}
	user.Password = enc
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
    user, err := s.repo.GetUser(username)
    if err != nil {
        return "", err
    }

    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
    if err != nil {
        return "", err
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
        jwt.StandardClaims{
            ExpiresAt: time.Now().Add(tokenTTL).Unix(),
            IssuedAt:  time.Now().Unix(),
        },
        user.Id,
    })

    return token.SignedString([]byte(signingkey))
}


func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingkey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func encryptedString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}