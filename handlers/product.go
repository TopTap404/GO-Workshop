package handlers

import (
	"github.com/gofiber/fiber/v2"

	"workshop/services"
)

type CreateProductRequest struct {
	Name  string  `json:"name"`
	Info  string  `json:"info"`
	Price float64 `json:"price"`
}

type UpdateProductRequest struct {
	Name  string  `json:"name"`
	Info  string  `json:"info"`
	Price float64 `json:"price"`
}

func ListProducts(c *fiber.Ctx) error {
	products, err := services.ListProducts()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(products)
}

func GetProduct(c *fiber.Ctx) error {
	id, err := parseID(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}

	product, err := services.GetProduct(id)
	if err != nil {
		if err == services.ErrNotFound {
			return c.Status(404).JSON(fiber.Map{"error": "not found"})
		}
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(product)
}

func CreateProduct(c *fiber.Ctx) error {
	var req CreateProductRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid json"})
	}

	product, err := services.CreateProduct(&services.CreateProductInput{
		Name:  req.Name,
		Info:  req.Info,
		Price: req.Price,
	})
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {
	id, err := parseID(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}

	var req UpdateProductRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid json"})
	}

	product, err := services.UpdateProduct(id, &req.Name, &req.Info, &req.Price)
	if err != nil {
		if err == services.ErrNotFound {
			return c.Status(404).JSON(fiber.Map{"error": "not found"})
		}
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(product)
}

func DeleteProduct(c *fiber.Ctx) error {
	id, err := parseID(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}

	if err := services.DeleteProduct(id); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(204)
}
