package apihome

import (
	model "nack/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetCategory(db *gorm.DB, c *fiber.Ctx) error {
	var category []model.Category

	if result := db.Find(&category); result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Not Get Banner"})
	}

	return c.JSON(category)
}

type Product struct {
	gorm.Model
	Code        string   `gorm:"type:varchar(255);not null;unique" json:"code"`
	Name        string   `gorm:"type:varchar(255);not null;unique" json:"name"`
	Description string   `gorm:"type:varchar(1024);not null" json:"description"`
	Price       float64  `gorm:"not null" json:"price"`
	Rating      int      `gorm:"not null" json:"rating"`
	Image       string   `gorm:"type:text" json:"image"`
	CategoryID  uint     `gorm:"not null" json:"category_id"`                                          // FK referencing Category.ID // Corrected to use Category
	MallID      uint     `gorm:"not null" json:"mall_id"`                                              // FK referencing Mall.ID    // Corrected to use Mall
}


func GetCategoryById(db *gorm.DB, c *fiber.Ctx) error {
	id := c.Params("id")

	var products []Product
	
    // ค้นหาผลิตภัณฑ์ทั้งหมดที่มี CategoryID ตรงกับที่ระบุ
    if result := db.Where("category_id = ?", id).Find(&products); result.Error != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "No products found for this category"})
    }

    return c.JSON(products)
}