package repositories

import (
	"workshop/database"
	"workshop/models"
)

func FindAllProducts() ([]models.Product, error) {
	var products []models.Product
	err := database.DB.Order("id desc").Find(&products).Error
	return products, err
}

func FindProductByID(id uint) (*models.Product, error) {
	var product models.Product
	err := database.DB.First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func CreateProduct(product *models.Product) error {
	return database.DB.Create(product).Error
}

func UpdateProduct(product *models.Product, updates map[string]any) error {
	return database.DB.Model(product).Updates(updates).Error
}

func DeleteProductByID(id uint) error {
	return database.DB.Delete(&models.Product{}, id).Error
}
