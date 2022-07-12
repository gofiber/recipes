package handlers

import (
	"robot-monitoreo/databases"
	"robot-monitoreo/models"

	"github.com/gofiber/fiber/v2"
)

func GetDogs(c *fiber.Ctx) error {
	var dogs []models.Dog

	databases.Database.Find(&dogs)
	return c.Status(fiber.StatusOK).JSON(dogs)
}

func GetDog(c *fiber.Ctx) error {
	id := c.Params("id")
	var dog models.Dog

	result := databases.Database.Find(&dog, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(&dog)
}

func AddDog(c *fiber.Ctx) error {
	dog := new(models.Dog)

	if err := c.BodyParser(dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	databases.Database.Create(&dog)
	return c.Status(201).JSON(dog)
}

func UpdateDog(c *fiber.Ctx) error {
	dog := new(models.Dog)
	id := c.Params("id")

	if err := c.BodyParser(dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	databases.Database.Where("id = ?", id).Updates(&dog)
	return c.Status(200).JSON(dog)
}

func RemoveDog(c *fiber.Ctx) error {
	id := c.Params("id")
	var dog models.Dog

	result := databases.Database.Delete(&dog, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}
