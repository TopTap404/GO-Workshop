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

	protected := api.Group("/", middleware.Auth)
	protected.Get("/users", handlers.ListUsers)
	protected.Get("/profile", handlers.GetMyProfile)
	protected.Get("/users/:id", handlers.GetUser)
	protected.Patch("/users/:id", handlers.UpdateUser)
	protected.Delete("/users/:id", handlers.DeleteUser)

	protected.Get("/products", handlers.ListProducts)
	protected.Get("/products/:id", handlers.GetProduct)
	protected.Post("/products", handlers.CreateProduct)
	protected.Patch("/products/:id", handlers.UpdateProduct)
	protected.Delete("/products/:id", handlers.DeleteProduct)

	protected.Get("/promotions", handlers.ListPromotions)
	protected.Post("/promotions", handlers.CreatePromotion)
	protected.Patch("/promotions/:id", handlers.UpdatePromotion)
	protected.Delete("/promotions/:id", handlers.DeletePromotion)

	protected.Get("/cart", handlers.GetMyCart)
	protected.Post("/cart/items", handlers.AddItemToCart)
	protected.Patch("/cart/items/:productId", handlers.UpdateMyCartItem)
	protected.Delete("/cart/items/:productId", handlers.RemoveMyCartItem)
	protected.Post("/cart/promotion", handlers.ApplyPromotionToMyCart)
	protected.Delete("/cart/promotion", handlers.ClearPromotionFromMyCart)
}
