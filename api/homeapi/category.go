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


func GetCategoryById(db *gorm.DB, c *fiber.Ctx) error {
	id := c.Params("id")

	var products []model.Product

    // ค้นหาผลิตภัณฑ์ทั้งหมดที่มี CategoryID ตรงกับที่ระบุ
    if result := db.Where("category_id = ?", id).Find(&products); result.Error != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "No products found for this category"})
    }

    return c.JSON(products)
}