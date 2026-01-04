package models

type Product struct {
	ID    uint    `json:"id" gorm:"primaryKey"`
	Name  string  `json:"name" gorm:"not null"`
	Info  string  `json:"info" gorm:"not null"`
	Price float64 `json:"price" gorm:"not null"`
}
