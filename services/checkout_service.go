package services

import (
	"errors"
	"workshop/models"
	"workshop/repositories"
)

var ErrCartEmpty = errors.New("cart is empty")

func Checkout(userID uint) (*models.Order, error) {
	cart, err := repositories.GetOrCreateCart(userID)
	if err != nil {
		return nil, err
	}

	cart, err = repositories.GetCart(cart.ID)
	if err != nil {
		return nil, err
	}

	if len(cart.Items) == 0 {
		return nil, ErrCartEmpty
	}

	var total float64 = 0
	var items []models.OrderItem

	for _, it := range cart.Items {
		product, err := repositories.FindProductByID(it.ProductID)
		if err != nil {
			return nil, err
		}

		lineTotal := float64(it.Quantity) * product.Price
		total += lineTotal

		items = append(items, models.OrderItem{
			ProductID: it.ProductID,
			Quantity:  int(it.Quantity),
			Price:     product.Price,
		})
	}

	order := &models.Order{
		UserID:     userID,
		Total:      total,
		Discount:   0,
		FinalTotal: total,
		Items:      items,
	}

	if err := repositories.CreateOrder(order); err != nil {
		return nil, err
	}

	return order, nil
}
