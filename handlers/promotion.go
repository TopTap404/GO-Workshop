package handlers

import (
	"github.com/gofiber/fiber/v2"

	"workshop/services"
)

type CreatePromotionRequest struct {
	Code           string  `json:"code"`
	DiscountAmount float64 `json:"discount_amount"`
	ProductID      uint    `json:"product_id"`
}

type UpdatePromotionRequest struct {
	Code           *string  `json:"code"`
	DiscountAmount *float64 `json:"discount_amount"`
	ProductID      *uint    `json:"product_id"`
}

func ListPromotions(c *fiber.Ctx) error {
	promos, err := services.ListPromotions()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(promos)
}

func CreatePromotion(c *fiber.Ctx) error {
	var req CreatePromotionRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid json"})
	}

	promo, err := services.CreatePromotion(&services.CreatePromotionInput{
		Code:           req.Code,
		DiscountAmount: req.DiscountAmount,
		ProductID:      req.ProductID,
	})
	if err != nil {
		if err == services.ErrNotFound {
			return c.Status(404).JSON(fiber.Map{"error": "product not found"})
		}
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(promo)
}

func UpdatePromotion(c *fiber.Ctx) error {
	id, err := parseID(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}

	var req UpdatePromotionRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid json"})
	}

	promo, err := services.UpdatePromotion(id, req.Code, req.DiscountAmount, req.ProductID)
	if err != nil {
		if err == services.ErrNotFound {
			return c.Status(404).JSON(fiber.Map{"error": "not found"})
		}
		if err == services.ErrNothingToUpdate {
			return c.Status(400).JSON(fiber.Map{"error": "no fields to update"})
		}
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(promo)
}

func DeletePromotion(c *fiber.Ctx) error {
	id, err := parseID(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}

	if err := services.DeletePromotion(id); err != nil {
		if err == services.ErrNotFound {
			return c.Status(404).JSON(fiber.Map{"error": "not found"})
		}
		if err == services.ErrPromotionInUse {
			return c.Status(400).JSON(fiber.Map{"error": "cannot delete promotion (in use)"})
		}
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(204)
}
