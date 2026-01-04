package repositories

import "workshop/database"

import "workshop/models"

func FindAllPromotions() ([]models.Promotion, error) {
	var promotions []models.Promotion
	if err := database.DB.Order("id desc").Find(&promotions).Error; err != nil {
		return nil, err
	}
	return promotions, nil
}

func FindPromotionByID(id uint) (*models.Promotion, error) {
	var promotion models.Promotion
	if err := database.DB.First(&promotion, id).Error; err != nil {
		return nil, err
	}
	return &promotion, nil
}

func FindPromotionByCode(code string) (*models.Promotion, error) {
	var promotion models.Promotion
	if err := database.DB.Where("code = ?", code).First(&promotion).Error; err != nil {
		return nil, err
	}
	return &promotion, nil
}

func CreatePromotion(promotion *models.Promotion) error {
	return database.DB.Create(promotion).Error
}

func UpdatePromotion(promotion *models.Promotion, updates map[string]any) error {
	return database.DB.Model(promotion).Updates(updates).Error
}

func DeletePromotionByID(id uint) (int64, error) {
	res := database.DB.Delete(&models.Promotion{}, id)
	return res.RowsAffected, res.Error
}

func CountCartsUsingPromotion(promotionID uint) (int64, error) {
	var count int64
	err := database.DB.Model(&models.Cart{}).Where("promotion_id = ?", promotionID).Count(&count).Error
	return count, err
}