package handlers

import (
	"auth-jwt-gorm/models"

	"github.com/gofiber/fiber/v3"
)

// AuthHandler contains HTTP handlers for authentication
type ProductHandler struct {
	productRepo *models.ProductRepository
}

// NewAuthHandler creates a new auth handler
func NewProductHandler(productRepo *models.ProductRepository) *ProductHandler {
	return &ProductHandler{
		productRepo: productRepo,
	}
}

// GetAllProducts query all products
func (ph *ProductHandler) GetAllProducts(c fiber.Ctx) error {
	products, err := ph.productRepo.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to retrieve products",
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All products",
		"data":    products,
	})
}

// GetProduct query product
func (ph *ProductHandler) GetProduct(c fiber.Ctx) error {
	product, err := ph.productRepo.GetById(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "No product found with ID",
			"data":    nil,
		})
	}
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Product found",
		"data":    product,
	})
}

// CreateProduct new product
func (ph *ProductHandler) CreateProduct(c fiber.Ctx) error {
	var product models.Product
	if err := c.Bind().Body(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Review your input",
			"data":    err,
		})
	}

	if err := ph.productRepo.Create(&product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Couldn't create product",
			"data":    err,
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Created product",
		"data":    product,
	})
}

// DeleteProduct delete product
func (ph *ProductHandler) DeleteProduct(c fiber.Ctx) error {
	product, err := ph.productRepo.GetById(c.Params("id"))
	if err != nil || product == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "No product found with ID",
			"data":    nil,
		})
	}

	if err = ph.productRepo.Delete(c.Params("id")); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Delete product failed!",
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Product successfully deleted",
		"data":    nil,
	})
}
