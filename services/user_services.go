package services

import (
	"errors"
	"strings"

	"gorm.io/gorm"

	"workshop/models"
	"workshop/repositories"

	"golang.org/x/crypto/bcrypt"
)

type CreateUserInput struct {
	Name        string
	LastName    string
	Email       string
	Password    string
	CitizenID   string
	PhoneNumber string
	Address     string
	AddressInfo string
}

func (in *CreateUserInput) Normalize() {
	in.Name = strings.TrimSpace(in.Name)
	in.LastName = strings.TrimSpace(in.LastName)
	in.Email = strings.ToLower(strings.TrimSpace(in.Email))
	in.Password = strings.TrimSpace(in.Password)
	in.CitizenID = strings.TrimSpace(in.CitizenID)
	in.PhoneNumber = strings.TrimSpace(in.PhoneNumber)
	in.Address = strings.TrimSpace(in.Address)
	in.AddressInfo = strings.TrimSpace(in.AddressInfo)
}

var (
	ErrNotFound        = errors.New("not found")
	ErrInvalidInput    = errors.New("invalid input")
	ErrNothingToUpdate = errors.New("no update found")
)

func ListUsers() ([]models.User, error) {
	return repositories.FindAllUsers()
}

func GetUser(id uint) (*models.User, error) {
	user, err := repositories.FindUserByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return user, nil
}

func CreateUser(in *CreateUserInput) (*models.User, error) {
	in.Normalize()

	if in.Name == "" || in.LastName == "" || in.Email == "" || in.Password == "" || in.CitizenID == "" || in.PhoneNumber == "" || in.Address == "" || in.AddressInfo == "" {
		return nil, ErrInvalidInput
	}

	if err := ValidateEmail(in.Email); err != nil {
		return nil, err
	}
	if err := ValidatePassword(in.Password); err != nil {
		return nil, err
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:        in.Name,
		LastName:    in.LastName,
		Email:       in.Email,
		Password:    string(hashed),
		CitizenID:   in.CitizenID,
		PhoneNumber: in.PhoneNumber,
		Address:     in.Address,
		AddressInfo: in.AddressInfo,
	}

	if err := repositories.CreateUser(user); err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUser(id uint, name, email *string) (*models.User, error) {
	user, err := repositories.FindUserByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	updates := map[string]any{}

	if name != nil {
		n := strings.TrimSpace(*name)
		if n == "" {
			return nil, ErrInvalidInput
		}
		updates["name"] = n
	}
	if email != nil {
		e := strings.TrimSpace(*email)
		if e == "" {
			return nil, ErrInvalidInput
		}
		updates["email"] = e
	}

	if len(updates) == 0 {
		return nil, ErrInvalidInput
	}

	if err := repositories.UpdateUser(user, updates); err != nil {
		return nil, err
	}

	return user, nil
}

func DeleteUser(id uint) error {
	return repositories.DeleteUserByID(id)
}

func GetProfile(userID uint) (*models.User, error) {
	user, err := repositories.FindUserByID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return user, nil
}
