package repositories

import (
	"workshop/database"
	"workshop/models"
)

func GetOrCreateCart(userID uint) (*models.Cart, error) {
	var cart models.Cart
	if err := database.DB.Where("user_id = ?", userID).First(&cart).Error; err == nil {
		return &cart, nil
	}
	cart = models.Cart{UserID: userID}
	if err := database.DB.Create(&cart).Error; err != nil {
		return nil, err
	}
	return &cart, nil
}

func AddItem(cartID, productID uint, qty int) error {
	var item models.CartItem
	if err := database.DB.Where("cart_id=? AND product_id=?", cartID, productID).First(&item).Error; err == nil {
		item.Quantity += uint(qty)
		return database.DB.Save(&item).Error
	}
	item = models.CartItem{CartID: cartID, ProductID: productID, Quantity: uint(qty)}
	return database.DB.Create(&item).Error
}

func GetCart(cartID uint) (*models.Cart, error) {
	var cart models.Cart
	if err := database.DB.Preload("Items").First(&cart, cartID).Error; err != nil {
		return nil, err
	}
	return &cart, nil
}
