package handler

import (
	"strconv"

	"api-fiber-gorm/database"
	"api-fiber-gorm/model"

	"github.com/gofiber/fiber/v3"
)

// GetAllProducts query all products
func GetAllProducts(c fiber.Ctx) error {
	db := database.DB
	var products []model.Product
	db.Find(&products)
	return c.JSON(fiber.Map{"status": "success", "message": "All products", "data": products})
}

// GetProduct query product
func GetProduct(c fiber.Ctx) error {
	// Parse the id instead of passing the raw string to GORM, which would treat a
	// non-numeric value as a raw WHERE clause (SQL injection).
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Invalid product ID", "data": nil})
	}
	db := database.DB
	var product model.Product
	db.Find(&product, id)
	if product.Title == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "No product found with ID", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Product found", "data": product})
}

// CreateProduct new product
func CreateProduct(c fiber.Ctx) error {
	db := database.DB
	product := new(model.Product)
	if err := c.Bind().Body(product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Couldn't create product", "data": err})
	}
	if err := db.Create(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Couldn't create product", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Created product", "data": product})
}

// DeleteProduct delete product
func DeleteProduct(c fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Invalid product ID", "data": nil})
	}
	db := database.DB

	var product model.Product
	db.First(&product, id)
	if product.Title == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "No product found with ID", "data": nil})
	}
	db.Delete(&product)
	return c.JSON(fiber.Map{"status": "success", "message": "Product successfully deleted", "data": nil})
}
