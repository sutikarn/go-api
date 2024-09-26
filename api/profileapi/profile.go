package apiProfile

import (
	"fmt"
	model "nack/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetProfile(db *gorm.DB, c *fiber.Ctx, UserID uint) error {
	var profile model.Profile
	result := db.Preload("User").Where("user_id = ?", UserID).First(&profile)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Not Get Profile"})
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

func UpdateProfile(db *gorm.DB, c *fiber.Ctx, userID uint) error {

	profile := new(model.Profile)

	db.Where("user_id = ?", userID).First(&profile)
	if err := c.BodyParser(profile); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	fmt.Println(profile.Mobile)
	// อัปเดตโปรไฟล์
	profile.UserID = userID
	if result := db.Save(&profile); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update profile"})
	}

    db.Save(&profile)
	return c.JSON(profile)
}
