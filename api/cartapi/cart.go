package apicart

import (
	model "nack/models"
	// homeapi "nack/api/homeapi"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	Price              float64 `gorm:"not null" json:"price"`
	Quantity           int     `gorm:"not null" json:"quantity"`
	ProductID          uint    `gorm:"not null" json:"productID"` // FK referencing Product.ID
	UserID             uint    `gorm:"not null" json:"userid"`
	ProductCode        string  `json:"product_code"`
	ProductName        string  `json:"product_name"`
	ProductDescription string  `json:"product_description"`
	ProductRating      int     `json:"product_rating"`
	ProductImage       string  `json:"product_image"`
}

func GetCart(db *gorm.DB, c *fiber.Ctx, userId uint) error {
	var carts []model.Cart
	var cartsResponse []Cart

	if err := db.Preload("Product").Find(&carts, "user_id = ?", userId).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "No products in cart "})
	}

	for _, cart := range carts {
		cartRes := Cart{
			Price:              cart.Price,
			Quantity:           cart.Quantity,
			ProductID:          cart.ProductID,
			ProductCode:        cart.Product.Code,
			ProductName:        cart.Product.Name,
			ProductDescription: cart.Product.Description,
			ProductRating:      cart.Product.Rating,
			ProductImage:       cart.Product.Image,
		}

		// ใช้ append เพื่อเพิ่มค่าเข้าไปใน slice
		cartsResponse = append(cartsResponse, cartRes)
	}

	return c.JSON(cartsResponse)
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
		"message": "Added Products to cart Success",
	})
}
