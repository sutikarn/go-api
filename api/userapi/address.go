package apiuser

import (
	model "nack/models"
	// homeapi "nack/api/homeapi"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAddress(db *gorm.DB, c *fiber.Ctx, userID uint) error {
	var addressResponse []model.Address

	if result := db.Where("user_id = ?", userID).Find(&addressResponse); result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "No Address "})
	}

	return c.JSON(fiber.Map{
		"message": "Get Address Success",
		"data":    addressResponse,
	})
}

func CreateAddress(db *gorm.DB, c *fiber.Ctx, userID uint) error {

	address := new(model.Address)

	if err := c.BodyParser(address); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	address.UserID = userID

	if result := db.Create(address); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create Address"})
	}

	return c.JSON(fiber.Map{
		"message": "Create Address Success",
	})
}

func UpdateAddress(db *gorm.DB, c *fiber.Ctx, userID uint) error {

	address := new(model.Address)

	//check profile userid
	db.Where("user_id = ?", userID).First(&address)
	if err := db.Where("user_id = ?", userID).First(&address).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Profile not found"})
	}

	// check request
	if err := c.BodyParser(address); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	// updateprofile
	address.UserID = userID
	if result := db.Save(&address); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update profile"})
	}

	// db.Save(&address)

	return c.JSON(fiber.Map{
		"message": "Update Address Success",
	})
}

func DeleteAddress(db *gorm.DB, c *fiber.Ctx, userId uint) error {

	var address model.Address
	id := c.Params("id")

	if result := db.Where("id = ?", id).Find(&address); result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user is not match"})
	}

	if address.UserID != userId {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	if result := db.Delete(&address); result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Not Delete Address"})
	}

	return c.JSON(fiber.Map{
		"message": "Delete Address Success",
	})
}
