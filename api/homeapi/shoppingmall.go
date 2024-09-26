package apihome


import (
	model "nack/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetShoppingMall(db *gorm.DB, c *fiber.Ctx) error {
	var mall []model.Mall

	if result := db.Find(&mall); result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Not Get Banner"})
	}

	return c.JSON(mall)
}


func GetShoppingMallById(db *gorm.DB, c *fiber.Ctx) error {
	var mall []model.Mall

	if result := db.Find(&mall); result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Not Get Banner"})
	}

	return c.JSON(mall)
}