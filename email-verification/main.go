package main

import (
	"email-verification/api/handlers"
	"email-verification/application"
	"email-verification/config"
	"email-verification/infrastructure/code"
	"email-verification/infrastructure/email"
	"email-verification/infrastructure/repository"
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/limiter"
)

func main() {
	// Dependencies
	cfg := config.GetConfig()
	repo := repository.NewMemoryRepository()
	emailService := email.NewSMTPService(cfg)
	codeGen := code.NewCodeGenerator()

	// Services
	verificationService := application.NewVerificationService(repo, emailService, codeGen, cfg)

	// Handlers
	verificationHandler := handlers.NewVerificationHandler(verificationService)

	// Router
	app := fiber.New()
	app.Post("/verify/send/:email",
		limiter.New(limiter.Config{Max: 5, Expiration: 1 * time.Minute}),
		verificationHandler.SendVerification,
	)
	app.Post("/verify/check/:email/:code", verificationHandler.CheckVerification)

	log.Fatal(app.Listen(":3000"))
}
