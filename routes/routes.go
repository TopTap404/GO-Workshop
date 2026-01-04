package routes

import (
	"github.com/gofiber/fiber/v2"

	"workshop/handlers"
	"workshop/middleware"
)

func Setup(app *fiber.App) {

	api := app.Group("/api")

	api.Post("/users", handlers.CreateUser)
	api.Post("/login", handlers.Login)

	protected := app.Group("/", middleware.Auth)
	protected.Get("/users", handlers.ListUsers)
	protected.Get("/users/:id", handlers.GetUser)
	protected.Get("/products", handlers.ListProducts)
	protected.Get("/products/:id", handlers.GetProduct)
	protected.Post("/products", handlers.CreateProduct)
	protected.Patch("/products/:id", handlers.UpdateProduct)
	protected.Patch("/users/:id", handlers.UpdateUser)
	protected.Delete("products/:id", handlers.DeleteProduct)
	protected.Delete("/users/:id", handlers.DeleteUser)
}
