
package handlers

import (
	"workshop/services"
	"github.com/gofiber/fiber/v2"
)

func Checkout(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	order, err := services.Checkout(userID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(order)
}
