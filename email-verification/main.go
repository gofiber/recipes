package main

import (
	"email-verification/api/handlers"
	"email-verification/application"
	"email-verification/config"
	"email-verification/infrastructure/code"
	"email-verification/infrastructure/email"
	"email-verification/infrastructure/repository"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Dependencies
	repo := repository.NewMemoryRepository()
	emailService := email.NewSMTPService(config.GetConfig())
	codeGen := code.NewCodeGenerator()

	// Services
	verificationService := application.NewVerificationService(repo, emailService, codeGen, config.GetConfig())

	// Handlers
	verificationHandler := handlers.NewVerificationHandler(verificationService)

	// Router
	app := fiber.New()
	app.Post("/verify/send/:email", verificationHandler.SendVerification)
	app.Post("/verify/check/:email/:code", verificationHandler.CheckVerification)

	app.Listen(":3000")
}
