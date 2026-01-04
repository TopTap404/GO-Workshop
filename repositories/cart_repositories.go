package repositories

import "workshop/database"

import "workshop/models"

func GetOrCreateCartByUserID(userID uint) (*models.Cart, error) {
	var cart models.Cart
	err := database.DB.Where("user_id = ?", userID).First(&cart).Error
	if err == nil {
		return &cart, nil
	}
	// create
	cart = models.Cart{UserID: userID}
	if err := database.DB.Create(&cart).Error; err != nil {
		return nil, err
	}
	return &cart, nil
}

func FindCartItems(cartID uint) ([]models.CartItem, error) {
	var items []models.CartItem
	if err := database.DB.Where("cart_id = ?", cartID).Order("id asc").Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func FindCartItemByCartAndProduct(cartID, productID uint) (*models.CartItem, error) {
	var item models.CartItem
	if err := database.DB.Where("cart_id = ? AND product_id = ?", cartID, productID).First(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func CreateCartItem(item *models.CartItem) error {
	return database.DB.Create(item).Error
}

func UpdateCartItem(item *models.CartItem, updates map[string]any) error {
	return database.DB.Model(item).Updates(updates).Error
}

func DeleteCartItemByCartAndProduct(cartID, productID uint) (int64, error) {
	res := database.DB.Where("cart_id = ? AND product_id = ?", cartID, productID).Delete(&models.CartItem{})
	return res.RowsAffected, res.Error
}

func SetCartPromotion(cart *models.Cart, promotionID *uint) error {
	return database.DB.Model(cart).Updates(map[string]any{"promotion_id": promotionID}).Error
}