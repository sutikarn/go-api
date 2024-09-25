package apiProfile

import (
	"fmt"
	model "nack/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetProfile(db *gorm.DB, c *fiber.Ctx, UserID int) error {
	cookie := c.Cookies("cookie_name")
   fmt.Printf("cookie: %s\n", cookie)
	var profile model.Profile
	result := db.Preload("User").Where("user_id = ?", UserID).First(&profile)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Not Get Banner"})
	}
	return c.JSON(profile)
}

func CreateProfile(db *gorm.DB, c *fiber.Ctx, UserID uint) error {
	profile := new(model.Profile)
	if err := c.BodyParser(profile); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}
	profile.UserID = UserID

	if result := db.Create(profile); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create Profile"})
	}
	return c.JSON(profile)
}
