package services

import (
	"errors"
	"strings"

	"gorm.io/gorm"

	"workshop/models"
	"workshop/repositories"
)

type CartItemView struct {
	Product models.Product `json:"product"`
	Quantity uint          `json:"quantity"`
	LineTotal float64      `json:"line_total"`
}

type CartView struct {
	Items []CartItemView `json:"items"`
	Subtotal float64 `json:"subtotal"`
	Discount float64 `json:"discount"`
	Total float64 `json:"total"`
	AppliedPromotion *models.Promotion `json:"promotion,omitempty"`
}

func GetCart(userID uint) (*CartView, error) {
	cart, err := repositories.GetOrCreateCartByUserID(userID)
	if err != nil { return nil, err }
	items, err := repositories.FindCartItems(cart.ID)
	if err != nil { return nil, err }

	view := &CartView{}
	sub := 0.0
	for _, it := range items {
		p, err := repositories.FindProductByID(it.ProductID)
		if err != nil {
			// ถ้าสินค้าถูกลบไปแล้ว ให้ข้าม (หรือจะ return error ก็ได้)
			continue
		}
		line := p.Price * float64(it.Quantity)
		sub += line
		view.Items = append(view.Items, CartItemView{Product: *p, Quantity: it.Quantity, LineTotal: line})
	}
	view.Subtotal = sub

	// Promotion
	var promo *models.Promotion
	if cart.PromotionID != nil {
		p, err := repositories.FindPromotionByID(*cart.PromotionID)
		if err == nil {
			promo = p
		}
	}
	view.AppliedPromotion = promo

	discount := 0.0
	if promo != nil {
		// ส่วนลดใช้ได้เฉพาะสินค้าที่ระบุ
		for _, it := range view.Items {
			if it.Product.ID == promo.ProductID {
				eligible := it.LineTotal
				discount = promo.DiscountAmount
				if discount > eligible {
					discount = eligible
				}
				break
			}
		}
	}
	view.Discount = discount
	view.Total = sub - discount
	if view.Total < 0 { view.Total = 0 }
	return view, nil
}

func AddToCart(userID, productID uint, quantity uint) (*CartView, error) {
	if productID == 0 || quantity == 0 {
		return nil, ErrInvalidInput
	}
	// product ต้องมีจริง
	if _, err := repositories.FindProductByID(productID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	cart, err := repositories.GetOrCreateCartByUserID(userID)
	if err != nil { return nil, err }

	item, err := repositories.FindCartItemByCartAndProduct(cart.ID, productID)
	if err == nil {
		updates := map[string]any{"quantity": item.Quantity + quantity}
		if err := repositories.UpdateCartItem(item, updates); err != nil { return nil, err }
		return GetCart(userID)
	}

	newItem := &models.CartItem{CartID: cart.ID, ProductID: productID, Quantity: quantity}
	if err := repositories.CreateCartItem(newItem); err != nil {
		return nil, err
	}
	return GetCart(userID)
}

func UpdateCartItemQty(userID, productID uint, quantity uint) (*CartView, error) {
	if productID == 0 {
		return nil, ErrInvalidInput
	}
	cart, err := repositories.GetOrCreateCartByUserID(userID)
	if err != nil { return nil, err }

	if quantity == 0 {
		_, err := repositories.DeleteCartItemByCartAndProduct(cart.ID, productID)
		if err != nil { return nil, err }
		return GetCart(userID)
	}

	item, err := repositories.FindCartItemByCartAndProduct(cart.ID, productID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	if err := repositories.UpdateCartItem(item, map[string]any{"quantity": quantity}); err != nil {
		return nil, err
	}
	return GetCart(userID)
}

func RemoveFromCart(userID, productID uint) (*CartView, error) {
	cart, err := repositories.GetOrCreateCartByUserID(userID)
	if err != nil { return nil, err }
	affected, err := repositories.DeleteCartItemByCartAndProduct(cart.ID, productID)
	if err != nil { return nil, err }
	if affected == 0 {
		return nil, ErrNotFound
	}
	return GetCart(userID)
}

func ApplyPromotionToCart(userID uint, code string) (*CartView, error) {
	code = strings.TrimSpace(strings.ToUpper(code))
	if code == "" {
		return nil, ErrInvalidInput
	}
	cart, err := repositories.GetOrCreateCartByUserID(userID)
	if err != nil { return nil, err }

	promo, err := repositories.FindPromotionByCode(code)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	pid := promo.ID
	if err := repositories.SetCartPromotion(cart, &pid); err != nil {
		return nil, err
	}
	return GetCart(userID)
}

func ClearPromotionFromCart(userID uint) (*CartView, error) {
	cart, err := repositories.GetOrCreateCartByUserID(userID)
	if err != nil { return nil, err }
	if err := repositories.SetCartPromotion(cart, nil); err != nil {
		return nil, err
	}
	return GetCart(userID)
}
