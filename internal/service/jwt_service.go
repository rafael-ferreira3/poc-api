package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenService struct{}

func NewTokenService() *TokenService {
	return &TokenService{}
}

var secretKey = []byte("secretKey")

func (t *TokenService) CreateJWTToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (t *TokenService) VerifyToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, getKeyFunc)

	if err != nil {
		return false
	}

	if !token.Valid {
		return false
	}

	return true
}

func getKeyFunc(token *jwt.Token) (interface{}, error) {
	return secretKey, nil
}
