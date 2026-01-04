package services

import (
	"errors"
	"strings"

	"gorm.io/gorm"

	"workshop/models"
	"workshop/repositories"
)

type CreateProductInput struct {
	Name  string
	Info  string
	Price float64
}

func (in *CreateProductInput) Normalize() {
	in.Name = strings.TrimSpace(in.Name)
	in.Info = strings.TrimSpace(in.Info)
}

func ListProducts() ([]models.Product, error) {
	return repositories.FindAllProducts()
}

func GetProduct(id uint) (*models.Product, error) {
	product, err := repositories.FindProductByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return product, nil
}

func CreateProduct(in *CreateProductInput) (*models.Product, error) {
	in.Normalize()

	if in.Name == "" || in.Info == "" || in.Price == 0 {
		return nil, ErrInvalidInput
	}

	product := &models.Product{
		Name:  in.Name,
		Info:  in.Info,
		Price: in.Price,
	}

	if err := repositories.CreateProduct(product); err != nil {
		return nil, err
	}
	return product, nil
}

func UpdateProduct(id uint, name *string, info *string, price *float64) (*models.Product, error) {

	product, err := repositories.FindProductByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	updates := map[string]any{}

	if name != nil {
		n := strings.TrimSpace(*name)
		if n == "" {
			return nil, ErrInvalidInput
		}
		updates["name"] = n
	}

	if info != nil {
		updates["info"] = strings.TrimSpace(*info)
	}

	if price != nil {
		if *price <= 0 {
			return nil, ErrInvalidInput
		}
		updates["price"] = *price
	}

	if len(updates) == 0 {
		return nil, ErrNothingToUpdate
	}

	if err := repositories.UpdateProduct(product, updates); err != nil {
		return nil, err
	}

	// reload เพื่อให้ response ตรง DB
	if product, err := repositories.FindProductByID(id); err != nil {
		return product, err
	}

	return product, nil
}

func DeleteProduct(id uint) error {
	return repositories.DeleteProductByID(id)
}
