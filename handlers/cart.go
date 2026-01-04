package handlers

import (
	"github.com/gofiber/fiber/v2"

	"workshop/middleware"
	"workshop/services"
)

type AddToCartRequest struct {
	ProductID uint `json:"product_id"`
	Quantity  uint `json:"quantity"`
}

type UpdateCartItemRequest struct {
	Quantity uint `json:"quantity"`
}

type ApplyPromotionRequest struct {
	Code string `json:"code"`
}

func GetMyCart(c *fiber.Ctx) error {
	userID := middleware.UserIDFromContext(c)
	if userID == 0 {
		return c.Status(401).JSON(fiber.Map{"error": "unauthorized"})
	}
	cart, err := services.GetCart(userID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(cart)
}

func AddItemToCart(c *fiber.Ctx) error {
	userID := middleware.UserIDFromContext(c)
	if userID == 0 {
		return c.Status(401).JSON(fiber.Map{"error": "unauthorized"})
	}
	var req AddToCartRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid json"})
	}
	cart, err := services.AddToCart(userID, req.ProductID, req.Quantity)
	if err != nil {
		switch err {
		case services.ErrNotFound:
			return c.Status(404).JSON(fiber.Map{"error": "product not found"})
		case services.ErrInvalidInput:
			return c.Status(400).JSON(fiber.Map{"error": "invalid input"})
		default:
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
	}
	return c.JSON(cart)
}

func UpdateMyCartItem(c *fiber.Ctx) error {
	userID := middleware.UserIDFromContext(c)
	if userID == 0 {
		return c.Status(401).JSON(fiber.Map{"error": "unauthorized"})
	}
	productID, err := parseID(c.Params("productId"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid product id"})
	}
	var req UpdateCartItemRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid json"})
	}
	cart, err := services.UpdateCartItemQty(userID, productID, req.Quantity)
	if err != nil {
		if err == services.ErrNotFound {
			return c.Status(404).JSON(fiber.Map{"error": "not found"})
		}
		if err == services.ErrInvalidInput {
			return c.Status(400).JSON(fiber.Map{"error": "invalid input"})
		}
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(cart)
}

func RemoveMyCartItem(c *fiber.Ctx) error {
	userID := middleware.UserIDFromContext(c)
	if userID == 0 {
		return c.Status(401).JSON(fiber.Map{"error": "unauthorized"})
	}
	productID, err := parseID(c.Params("productId"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid product id"})
	}
	cart, err := services.RemoveFromCart(userID, productID)
	if err != nil {
		if err == services.ErrNotFound {
			return c.Status(404).JSON(fiber.Map{"error": "not found"})
		}
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(cart)
}

func ApplyPromotionToMyCart(c *fiber.Ctx) error {
	userID := middleware.UserIDFromContext(c)
	if userID == 0 {
		return c.Status(401).JSON(fiber.Map{"error": "unauthorized"})
	}
	var req ApplyPromotionRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid json"})
	}
	cart, err := services.ApplyPromotionToCart(userID, req.Code)
	if err != nil {
		if err == services.ErrNotFound {
			return c.Status(404).JSON(fiber.Map{"error": "promotion not found"})
		}
		if err == services.ErrInvalidInput {
			return c.Status(400).JSON(fiber.Map{"error": "invalid input"})
		}
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(cart)
}

func ClearPromotionFromMyCart(c *fiber.Ctx) error {
	userID := middleware.UserIDFromContext(c)
	if userID == 0 {
		return c.Status(401).JSON(fiber.Map{"error": "unauthorized"})
	}
	cart, err := services.ClearPromotionFromCart(userID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(cart)
}
