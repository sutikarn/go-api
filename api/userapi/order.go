package apiuser

import (
	// cartapi "nack/api/cartapi"
	model "nack/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Price     float64 `gorm:"not null" json:"price"`
	Quantity  int     `gorm:"not null" json:"quantity"`
	ProductID uint    `gorm:"not null" json:"productID"` // FK referencing Product.ID
	UserID    uint    `gorm:"not null" json:"userid"`
}

func CreateOrder(db *gorm.DB, c *fiber.Ctx, userID uint) error {
	var orderRequest []Order

	// Parsing request body
	if err := db.Model(&model.Cart{}).
		Joins("JOIN products ON carts.product_id = products.id").
		Where("carts.user_id = ?", userID).
		Select("carts.price, carts.quantity, carts.product_id").
		Scan(&orderRequest).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "No products in cart"})
	}

	for i := range orderRequest {
		orderRequest[i].UserID = userID
	}

	// บันทึกข้อมูลลงในฐานข้อมูล
	if err := db.Create(&orderRequest).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not add to cart",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Create Order Success",
		"data":    orderRequest,
	})
}

type Orders struct {
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

func GetOrder(db *gorm.DB, c *fiber.Ctx, userID uint) error {
	var orderResponse []Orders

	if err := db.Model(&model.Order{}).
		Joins("JOIN products ON orders.product_id = products.id").
		Where("orders.user_id = ?", userID).
		Select("orders.id,orders.price, orders.quantity, orders.product_id, orders.user_id,products.code AS product_code, products.name AS product_name, products.description AS product_description, products.price AS product_price, products.image AS product_image ").
		Scan(&orderResponse).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "No products in cart"})
	}

	return c.JSON(fiber.Map{
		"message": "Get Order Success",
		"data":    orderResponse,
	})
}
