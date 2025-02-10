package handlers

import (
	"email-verification/application"

	"github.com/gofiber/fiber/v2"
)

type VerificationHandler struct {
	verificationService *application.VerificationService
}

func NewVerificationHandler(service *application.VerificationService) *VerificationHandler {
	return &VerificationHandler{verificationService: service}
}

func (h *VerificationHandler) SendVerification(c *fiber.Ctx) error {
	email := c.Params("email")
	if err := h.verificationService.SendVerification(email); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{"message": "Verification code sent"})
}

func (h *VerificationHandler) CheckVerification(c *fiber.Ctx) error {
	email := c.Params("email")
	code := c.Params("code")

	if err := h.verificationService.VerifyCode(email, code); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{"message": "Code verified successfully"})
}
