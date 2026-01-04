package models

import "time"

type User struct {
	ID uint `json:"id" gorm:"primaryKey"`

	Name     string `json:"name" gorm:"not null"`
	LastName string `json:"last_name" gorm:"not null"`

	Email    string `json:"email" gorm:"size:255;uniqueIndex;not null"`
	Password string `json:"-" gorm:"not null"`

	CitizenID   string `json:"citizenid" gorm:"not null"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	AddressInfo string `json:"address_info"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
