package handlers

import (
	"workshop/services"

	"github.com/gofiber/fiber/v2"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid json"})
	}

	user, err := services.Login(req.Email, req.Password)
	if err != nil {
		if err == services.ErrInvalidLogin {
			return c.Status(401).JSON(fiber.Map{"error": "invalid email or password"})
		}
		return c.Status(500).JSON(fiber.Map{"error": "internal error"})
	}

	token, expiresIn, err := services.GenerateAccessToken(user.ID, user.Email)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to generate token"})
	}

	return c.JSON(fiber.Map{
		"access_token": token,
		"token_type":   "Bearer",
		"expires_in":   expiresIn,
		"user": fiber.Map{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}
