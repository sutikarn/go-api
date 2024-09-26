package apihome

import (
	// model "nack/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetProduct(db *gorm.DB, c *fiber.Ctx) error {
	var product []Product

	if result := db.Find(&product); result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Not Get Product"})
	}

	return c.JSON(product)
}


func GetProductById(db *gorm.DB, c *fiber.Ctx) error {
	var product []Product
	id := c.Params("id")
	if result := db.Where("id = ?", id).Find(&product); result.Error != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "No products id "})
    }

	if len(product) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "No product id"})
	}

	return c.JSON(product)
}
