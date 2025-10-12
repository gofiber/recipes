package router

import (
	"auth-jwt-gorm/database"
	"auth-jwt-gorm/handlers"
	"auth-jwt-gorm/middleware"
	"auth-jwt-gorm/models"
	"auth-jwt-gorm/services"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New())
	api.Get("/", handlers.Hello)

	// Auth
	userRepo := models.NewUserRepository(database.DB)
	refreshTokenRepo := models.NewRefreshTokenRepository(database.DB)
	authService := services.NewAuthService(userRepo, refreshTokenRepo, os.Getenv("JWT_SECRET"), 15*time.Minute)
	authHandler := handlers.NewAuthHandler(authService)

	auth := api.Group("/auth")
	auth.Post("/login", authHandler.Login)
	auth.Post("/register", authHandler.Register)
	auth.Post("/logout", authHandler.Logout)
	auth.Post("/refresh-token", authHandler.RefreshToken)

	// User
	userHandler := handlers.NewUserHandler(userRepo)
	user := api.Group("/users")
	user.Get("/:id", middleware.Protected(), userHandler.GetUser)
	user.Post("/", middleware.Protected(), userHandler.CreateUser)
	user.Patch("/:id", middleware.Protected(), userHandler.UpdateUser)
	user.Delete("/:id", middleware.Protected(), userHandler.DeleteUser)

	// Product
	productRepo := models.NewProductRepository(database.DB)
	productHandler := handlers.NewProductHandler(productRepo)

	product := api.Group("/products")
	product.Get("/", productHandler.GetAllProducts)
	product.Get("/:id", productHandler.GetProduct)
	product.Post("/", middleware.Protected(), productHandler.CreateProduct)
	product.Delete("/:id", middleware.Protected(), productHandler.DeleteProduct)
}
