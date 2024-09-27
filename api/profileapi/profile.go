package apiProfile

import (
	"fmt"
	// model "nack/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	FirstName string `gorm:"type:varchar(255);not null" json:"firstname"`
	LastName  string `gorm:"type:varchar(255);not null" json:"lastname"`
	Mobile    string `gorm:"type:varchar(10)" json:"mobile"`
	Sex       string `gorm:"type:varchar(10)" json:"sex"`
	Status    string `gorm:"type:varchar(10)" json:"status"`
	Image     string `gorm:"type:text" json:"image"`
	UserID    uint   `gorm:"not null" json:"userid"` // FK referencing User.ID
}

func GetProfile(db *gorm.DB, c *fiber.Ctx, userID uint) error {
	var profile Profile
	if err := db.First(&profile, "user_id = ?", userID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Profile not found"})
	}
	return c.JSON(profile)
}

func CreateProfile(db *gorm.DB, c *fiber.Ctx, UserID uint) error {
	profile := new(Profile)

	if err := db.First(&profile, "user_id = ?", UserID).Error; err == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Profile not create"})
	}

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

	profile := new(Profile)

	//check profile userid
	db.Where("user_id = ?", userID).First(&profile)
	if err := db.Select("id, created_at, updated_at, deleted_at, first_name, last_name, mobile, sex, status, image, user_id").First(&profile, "user_id = ?", userID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Profile not found"})
	}

	// check request
	if err := c.BodyParser(profile); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	// updateprofile
	profile.UserID = userID
	if result := db.Save(&profile); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update profile"})
	}

	fmt.Println(profile)

	db.Save(&profile)
	return c.JSON(profile)
}
