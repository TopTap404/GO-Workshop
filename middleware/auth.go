package middleware

import (
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Auth(c *fiber.Ctx) error {
	auth := c.Get("Authorization")
	if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
		return c.Status(401).JSON(fiber.Map{"error": "missing token"})
	}

	tokenStr := strings.TrimPrefix(auth, "Bearer ")

	secret := []byte(os.Getenv("JWT_SECRET"))
	if len(secret) == 0 {
		secret = []byte("dev-secret")
	}

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
		// กันโดนเปลี่ยน alg
		if t.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return secret, nil
	})

	if err != nil || !token.Valid {
		return c.Status(401).JSON(fiber.Map{"error": "invalid or expired token"})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(401).JSON(fiber.Map{"error": "invalid token claims"})
	}

	// ส่งต่อให้ handler ใช้
	c.Locals("user_id", claims["sub"])
	c.Locals("email", claims["email"])

	return c.Next()
}
