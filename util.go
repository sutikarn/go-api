package main

import ("github.com/gofiber/fiber/v2")

func getUserId(c *fiber.Ctx) (uint, error) {
	userID := c.Locals("userID")
	// ตรวจสอบ type ของ userID เพื่อให้แน่ใจว่าเป็น float64 (ค่าที่ JWT อาจให้มาเป็น float64)
	userIDFloat, ok := userID.(float64)
	if !ok {
		return 0, c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	// แปลงเป็น int
	userIDInt := uint(userIDFloat)
	return userIDInt, nil
}