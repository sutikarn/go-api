package apicart

import (
	model "nack/models"
	// homeapi "nack/api/homeapi"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	Price    float64 `gorm:"not null" json:"price"`
	Quantity int     `gorm:"not null" json:"quantity"`
	// ProductID          uint    `gorm:"not null" json:"productID"` // FK referencing Product.ID
	UserID             uint   `gorm:"not null" json:"userid"`
	ProductCode        string `json:"product_code"`
	ProductName        string `json:"product_name"`
	ProductDescription string `json:"product_description"`
	ProductRating      int    `json:"product_rating"`
	ProductImage       string `json:"product_image"`
}

func GetCart(db *gorm.DB, c *fiber.Ctx, userId uint) error {
	var cartsResponse []Cart

	if err := db.Model(&model.Cart{}).
		Joins("JOIN products ON carts.product_id = products.id").
		Where("carts.user_id = ?", userId).
		Select("carts.id,carts.price, carts.quantity, carts.product_id, carts.user_id,products.code AS product_code, products.name AS product_name, products.description AS product_description, products.price AS product_price, products.image AS product_image ").
		Scan(&cartsResponse).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "No products in cart"})
	}

	return c.JSON(cartsResponse)
}

func AddToCart(db *gorm.DB, c *fiber.Ctx, userId uint) error {
	// สร้าง slice ของ Cart เพื่อเก็บรายการสินค้า
	var carts []model.Cart

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
		"message": "Added Products to cart Success",
	})
}
