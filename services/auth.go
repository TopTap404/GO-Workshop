package services

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateAccessToken(userID uint, email string) (token string, expiresIn int, err error) {
	secret := []byte(os.Getenv("JWT_SECRET"))
	if len(secret) == 0 {
		secret = []byte("dev-secret")
	}

	ttl := time.Hour
	now := time.Now()

	claims := jwt.MapClaims{
		"sub":   fmt.Sprintf("%d", userID),
		"email": email,
		"iat":   now.Unix(),
		"exp":   now.Add(ttl).Unix(),
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := t.SignedString(secret)
	if err != nil {
		return "", 0, err
	}

	return signed, int(ttl.Seconds()), nil
}
