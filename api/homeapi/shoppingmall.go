package apihome


import (
	model "nack/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetShoppingMall(db *gorm.DB, c *fiber.Ctx) error {
	var mall []model.Mall

	if result := db.Find(&mall); result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Not Get Banner"})
	}

	return c.JSON(mall)
}


func GetShoppingMallById(db *gorm.DB, c *fiber.Ctx) error {

	id := c.Params("id")

	var products []Product
	
    // ค้นหาผลิตภัณฑ์ทั้งหมดที่มี mall_ID ตรงกับที่ระบุ
    if result := db.Where("mall_id = ?", id).Find(&products); result.Error != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "No products found for this category"})
    }

    return c.JSON(products)
}