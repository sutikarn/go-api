package apimeow

import (
	"nack/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Signup(db *gorm.DB, c *fiber.Ctx) error {
	// สร้าง UUID ใหม่
	newUUID := uuid.New()

	// แปลง UUID เป็น string
	uuidString := newUUID.String()
	user := new(model.User)

	// Parse request body into user
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	// Check for existing email
	var existingUser model.User
	if err := db.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		// Email already exists
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Email already exists"})
	}

	// Encrypt the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not hash password"})
	}
	user.Password = string(hashedPassword)
	user.Code = uuidString

	// Create user
	if result := db.Create(user); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create user"})
	}

	// Return created user
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Register Success"})
}
