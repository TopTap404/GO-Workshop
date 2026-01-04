package models

import "time"

// Cart: 1 user มี 1 cart และสามารถผูก Promotion ได้ 1 ตัว
type Cart struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	UserID      uint       `json:"user_id" gorm:"uniqueIndex;not null"`
	Items       []CartItem `gorm:"foreignKey:CartID"`
	PromotionID *uint      `json:"promotion_id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type CartItem struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CartID    uint      `json:"cart_id" gorm:"index;not null"`
	ProductID uint      `json:"product_id" gorm:"index;not null"`
	Quantity  uint      `json:"quantity" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
