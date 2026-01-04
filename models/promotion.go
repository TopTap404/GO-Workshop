package models

import "time"

// Promotion: ส่วนลดเป็นจำนวนเงิน (บาท) และใช้ได้กับ "สินค้า" ที่ระบุเท่านั้น
// - Code ต้อง unique
// - DiscountAmount เป็นจำนวนเงิน (บาท) ไม่ใช่เปอร์เซ็นต์
type Promotion struct {
	ID uint `json:"id" gorm:"primaryKey"`

	Code           string  `json:"code" gorm:"size:64;uniqueIndex;not null"`
	DiscountAmount float64 `json:"discount_amount" gorm:"not null"`

	// ลดให้กับสินค้าที่ระบุเท่านั้น
	ProductID uint `json:"product_id" gorm:"not null"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
