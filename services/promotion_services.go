package services

import (
	"errors"
	"strings"

	"gorm.io/gorm"

	"workshop/models"
	"workshop/repositories"
)

var (
	ErrPromotionInUse = errors.New("promotion is in use")
)

type CreatePromotionInput struct {
	Code           string
	DiscountAmount float64
	ProductID      uint
}

func (in *CreatePromotionInput) Normalize() {
	in.Code = strings.TrimSpace(strings.ToUpper(in.Code))
}

func ListPromotions() ([]models.Promotion, error) {
	return repositories.FindAllPromotions()
}

func CreatePromotion(in *CreatePromotionInput) (*models.Promotion, error) {
	in.Normalize()
	if in.Code == "" || in.DiscountAmount <= 0 || in.ProductID == 0 {
		return nil, ErrInvalidInput
	}

	if _, err := repositories.FindProductByID(in.ProductID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	promo := &models.Promotion{
		Code:           in.Code,
		DiscountAmount: in.DiscountAmount,
		ProductID:      in.ProductID,
	}
	if err := repositories.CreatePromotion(promo); err != nil {
		return nil, err
	}
	return promo, nil
}

func UpdatePromotion(id uint, code *string, discountAmount *float64, productID *uint) (*models.Promotion, error) {
	promo, err := repositories.FindPromotionByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	updates := map[string]any{}
	if code != nil {
		c := strings.TrimSpace(strings.ToUpper(*code))
		if c == "" {
			return nil, ErrInvalidInput
		}
		updates["code"] = c
	}
	if discountAmount != nil {
		if *discountAmount <= 0 {
			return nil, ErrInvalidInput
		}
		updates["discount_amount"] = *discountAmount
	}
	if productID != nil {
		if *productID == 0 {
			return nil, ErrInvalidInput
		}

		if _, err := repositories.FindProductByID(*productID); err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, ErrNotFound
			}
			return nil, err
		}
		updates["product_id"] = *productID
	}

	if len(updates) == 0 {
		return nil, ErrNothingToUpdate
	}
	if err := repositories.UpdatePromotion(promo, updates); err != nil {
		return nil, err
	}
	return promo, nil
}

func DeletePromotion(id uint) error {
	promo, err := repositories.FindPromotionByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrNotFound
		}
		return err
	}
	count, err := repositories.CountCartsUsingPromotion(promo.ID)
	if err != nil {
		return err
	}
	if count > 0 {
		return ErrPromotionInUse
	}
	_, err = repositories.DeletePromotionByID(promo.ID)
	return err
}
