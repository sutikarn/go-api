package apihome

import (
	model "nack/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetBanner(db *gorm.DB, c *fiber.Ctx) error {
	var banners []model.Banner

	if result := db.Find(&banners); result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Not Get Banner"})
	}

	return c.JSON(banners)
}

func GetHome(db *gorm.DB, c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNoContent)
}

