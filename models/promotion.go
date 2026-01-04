package models

import "time"

type Promotion struct {
	ID uint `json:"id" gorm:"primaryKey"`

	Code           string  `json:"code" gorm:"size:64;uniqueIndex;not null"`
	DiscountAmount float64 `json:"discount_amount" gorm:"not null"`

	ProductID uint `json:"product_id" gorm:"not null"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
