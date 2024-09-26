package apihome

import (
	model "nack/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetProduct(db *gorm.DB, c *fiber.Ctx) error {
	var product []model.Product

	if result := db.Find(&product); result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Not Get Product"})
	}

	return c.JSON(product)
}


func GetProductById(db *gorm.DB, c *fiber.Ctx) error {
	var product []model.Product

	if result := db.Find(&product); result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Not Get Product"})
	}

	return c.JSON(product)
}
