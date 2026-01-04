package middleware

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Logger prints basic request info and latency.
func Logger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		err := c.Next()

		status := c.Response().StatusCode()
		log.Printf("%s %s -> %d (%s)", c.Method(), c.Path(), status, time.Since(start))
		return err
	}
}
