package models

import "time"

type Order struct {
	ID         uint `gorm:"primaryKey"`
	UserID     uint
	Total      float64
	Discount   int
	FinalTotal float64
	CreatedAt  time.Time
	Items      []OrderItem
}

type OrderItem struct {
	ID        uint `gorm:"primaryKey"`
	OrderID   uint
	ProductID uint
	Quantity  int
	Price     float64
}
