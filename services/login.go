package services

import (
	"errors"
	"strings"
	"workshop/models"
	"workshop/repositories"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var ErrInvalidLogin = errors.New("Invalid Email or Password")

func Login(email, password string) (*models.User, error) {
	email = strings.TrimSpace((strings.ToLower(email)))
	password = strings.TrimSpace(password)

	if email == "" || password == "" {
		return nil, ErrInvalidInput
	}

	user, err := repositories.FindUserByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrInvalidLogin
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, ErrInvalidLogin
	}
	return user, nil
}
