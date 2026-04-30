package handler

import (
	"app/database"
	"app/model"

	"github.com/gofiber/fiber/v3"
)

// GetAllProducts query all products
func GetAllProducts(c fiber.Ctx) error {
	db := database.DB
	var products []model.Product
	if result := db.Find(&products); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Couldn't retrieve products", "data": result.Error.Error()})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "All products", "data": products})
}

// GetProduct query product
func GetProduct(c fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var product model.Product
	if result := db.Find(&product, id); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Couldn't retrieve product", "data": result.Error.Error()})
	}
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
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Couldn't create product", "data": err.Error()})
	}
	if result := db.Create(&product); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Couldn't create product", "data": result.Error.Error()})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Created product", "data": product})
}

// DeleteProduct delete product
func DeleteProduct(c fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	var product model.Product
	db.First(&product, id)
	if product.Title == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "No product found with ID", "data": nil})
	}
	if result := db.Delete(&product); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Couldn't delete product", "data": result.Error.Error()})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Product successfully deleted", "data": nil})
}
