package services

import (
	"errors"
	"regexp"
	"strings"
	"unicode"
)

var (
	ErrInvalidEmail = errors.New("invalid email format")
	ErrWeakPassword = errors.New("password does not meet requirements")
)

var emailRegex = regexp.MustCompile(
	`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`,
)

func ValidateEmail(email string) error {
	email = strings.TrimSpace(email)
	if !emailRegex.MatchString(email) {
		return ErrInvalidEmail
	}
	return nil
}

func ValidatePassword(password string) error {
	password = strings.TrimSpace(password)

	if len(password) < 10 {
		return ErrWeakPassword
	}

	hasUpper := false
	hasLower := false

	for _, r := range password {
		if r > unicode.MaxASCII {
			return ErrWeakPassword
		}
		if unicode.IsSpace(r) {
			return ErrWeakPassword
		}
		if r >= 'A' && r <= 'Z' {
			hasUpper = true
		}
		if r >= 'a' && r <= 'z' {
			hasLower = true
		}
	}

	if !hasUpper || !hasLower {
		return ErrWeakPassword
	}
	return nil
}
