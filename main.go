package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"workshop/database"
	"workshop/middleware"
	"workshop/routes"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	app := fiber.New(fiber.Config{
		AppName: "Workshop Fiber + GORM (No gRPC)",
	})

	// Safety middleware
	app.Use(recover.New())

	// Custom logger middleware
	app.Use(middleware.Logger())

	// Connect DB + migrate
	if err := database.ConnectAndMigrate(); err != nil {
		log.Fatal(err)
	}

	// Serve minimal frontend (optional)
	app.Static("/", "./public")

	// API routes
	routes.Setup(app)

	log.Println("Server listening on http://localhost:3000")
	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}
