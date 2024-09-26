package apicart

import (
	// model "nack/models"
	homeapi "nack/api/homeapi"
	

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	Price     float64 `gorm:"not null" json:"price"`
	Quantity  int     `gorm:"not null" json:"quantity"`
	ProductID uint    `gorm:"not null" json:"productID"` // FK referencing Product.ID
	UserID    uint    `gorm:"not null" json:"userid"`
	Product   homeapi.Product    // FK referencing User.ID
}

func GetCart(db *gorm.DB, c *fiber.Ctx, userId uint) error {
	var carts []Cart

	if err := db.Preload("Product").Find(&carts, "user_id = ?", userId).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "No products in cart "})
	}

	return c.JSON(carts)
}

func AddToCart(db *gorm.DB, c *fiber.Ctx, userId uint) error {
	// สร้าง slice ของ Cart เพื่อเก็บรายการสินค้า
	var carts []Cart

	// Parsing request body
	if err := c.BodyParser(&carts); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	// เพิ่ม userID ให้กับแต่ละรายการสินค้าใน cart
	for i := range carts {
		carts[i].UserID = userId
	}

	// บันทึกข้อมูลลงในฐานข้อมูล
	if err := db.Create(&carts).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not add to cart",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Products added to cart",
		"cart":    carts,
	})
}
