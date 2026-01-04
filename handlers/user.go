package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"workshop/middleware"
	"workshop/services"
)

type CreateUserRequest struct {
	Name        string `json:"name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	AddressInfo string `json:"address_info"`
}

type UpdateUserRequest struct {
	Name  *string `json:"name"`
	Email *string `json:"email"`
}

func ListUsers(c *fiber.Ctx) error {
	users, err := services.ListUsers()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	id, err := parseID(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}

	user, err := services.GetUser(id)
	if err != nil {
		if err == services.ErrNotFound {
			return c.Status(404).JSON(fiber.Map{"error": "not found"})
		}
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(user)
}

func CreateUser(c *fiber.Ctx) error {
	var req CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid json"})
	}

	user, err := services.CreateUser(&services.CreateUserInput{
		Name:        req.Name,
		LastName:    req.LastName,
		Email:       req.Email,
		Password:    req.Password,
		PhoneNumber: req.PhoneNumber,
		Address:     req.Address,
		AddressInfo: req.AddressInfo,
	})
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id, err := parseID(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}

	var req UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid json"})
	}

	user, err := services.UpdateUser(id, req.Name, req.Email)
	if err != nil {
		if err == services.ErrNotFound {
			return c.Status(404).JSON(fiber.Map{"error": "not found"})
		}
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	id, err := parseID(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}

	if err := services.DeleteUser(id); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(204)
}

func GetMyProfile(c *fiber.Ctx) error {
	userID := middleware.UserIDFromContext(c)
	if userID == 0 {
		return c.Status(401).JSON(fiber.Map{"error": "invalid token"})
	}

	user, err := services.GetProfile(userID)
	if err != nil {
		if err == services.ErrNotFound {
			return c.Status(404).JSON(fiber.Map{"error": "not found"})
		}
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"id":           user.ID,
		"name":         user.Name,
		"last_name":    user.LastName,
		"email":        user.Email,
		"phone_number": user.PhoneNumber,
		"address":      user.Address,
		"address_info": user.AddressInfo,
		"created_at":   user.CreatedAt,
		"updated_at":   user.UpdatedAt,
	})
}

func parseID(raw string) (uint, error) {
	n, err := strconv.ParseUint(raw, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(n), nil
}
