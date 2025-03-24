package service

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

func CreateToken(userId int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	secret := []byte(os.Getenv("JWT_SECRET_KEY"))
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", fmt.Errorf("error creating token: %w", err)
	}

	return tokenString, nil
}
