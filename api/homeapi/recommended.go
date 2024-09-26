package apihome


import (
	// model "nack/models"
	"math/rand"
	"time"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetProductRecommend(db *gorm.DB, c *fiber.Ctx) error {
	var products []Product


	if result := db.Find(&products); result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Not Get Product"})
	}
	

	// สุ่มดึงค่า 10 อัน
	rand.NewSource(time.Now().UnixNano()) // ตั้ง seed สำหรับการสุ่ม
	rand.Shuffle(len(products), func(i, j int) {
		products[i], products[j] = products[j], products[i]
	})
    
	var selectedProducts []Product
	if len(products) >= 10 {
		selectedProducts = products[:10]
	} else {
		selectedProducts = products[:]
	} 

	// ดึง 10 อันแรก


	return c.JSON(selectedProducts)
}