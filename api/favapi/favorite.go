package apiFav

import (
	model "nack/models"
	// homeapi "nack/api/homeapi"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Favorite struct {
	gorm.Model
	Price    float64 `gorm:"not null" json:"price"`
	Quantity int     `gorm:"not null" json:"quantity"`
	ProductID          uint    `gorm:"not null" json:"productID"` // FK referencing Product.ID
	FavoriteDate       string `json:"favorite_date"`
	UserID             uint   `gorm:"not null" json:"userid"`
	ProductCode        string `json:"product_code"`
	ProductName        string `json:"product_name"`
	ProductDescription string `json:"product_description"`
	ProductRating      int    `json:"product_rating"`
	ProductImage       string `json:"product_image"`
}

func GetFavorite(db *gorm.DB, c *fiber.Ctx, userId uint) error {
	var favResponse []Favorite

	if err := db.Model(&model.Favorite{}).
		Joins("JOIN products ON favorites.product_id = products.id").
		Where("favorites.user_id = ?", userId).
		Select("favorites.id,favorites.price, favorites.quantity, favorites.product_id, favorites.user_id,favorites.created_at AS favorite_date,products.code AS product_code, products.name AS product_name, products.description AS product_description, products.price AS product_price, products.image AS product_image ").
		Scan(&favResponse).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "No Favorite"})
	}

	return c.JSON(fiber.Map{
		"message": "Get Favorite Success",
		"data":favResponse,
	})
}

func AddFavorite(db *gorm.DB, c *fiber.Ctx, userId uint) error {

	// สร้าง slice ของ Cart เพื่อเก็บรายการสินค้า
	var favorite model.Favorite

	// Parsing request body
	if err := c.BodyParser(&favorite); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	favorite.UserID = userId

	// บันทึกข้อมูลลงในฐานข้อมูล
	if err := db.Create(&favorite).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not add to cart",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Added Favorite Success",
	})
}

func DeleteFavorite(db *gorm.DB, c *fiber.Ctx, userId uint) error {
    var favorite model.Favorite
	id := c.Params("id") 

	if result := db.Where("id = ?", id).Find(&favorite); result.Error != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user is not match"})
    }

	if favorite.UserID != userId {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	if result := db.Delete(&favorite); result.Error != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Not Delete Fav"})
    }
	

	return c.JSON(fiber.Map{
		"message": "Delete Favorite Success",
	})
}
