package main

import (
	"fmt"
	apihome "nack/api/homeapi"
	apiProfile "nack/api/profileapi"
	apimeow "nack/api/signin_signup"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)


func authRequired(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil || !token.Valid {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	claim := token.Claims.(jwt.MapClaims)

	// ตรวจสอบว่ามี user_id ใน claims หรือไม่
	if userID, exists := claim["user_id"]; exists {
		c.Locals("userID", userID) // บันทึก userID ลงใน context
	} else {
		return c.SendStatus(fiber.StatusUnauthorized) // ส่งสถานะ Unauthorized
	}

	fmt.Println(claim)

	return c.Next()
}

func setPathApi (){
	app := fiber.New()

	///set path/authen check auth
	app.Use("/authen", authRequired)

	///noauthen
	app.Post("/sign-up", func(c *fiber.Ctx) error {
		return apimeow.Signup(db, c)
	})

	app.Post("/sign-in", func(c *fiber.Ctx) error {
		return apimeow.Signin(db, c)
	})

	app.Get("/home", func(c *fiber.Ctx) error {
		return apihome.GetHome(db, c)
	})

	///goauthen

	app.Get("/authen/banner", func(c *fiber.Ctx) error {
		return apihome.GetBanner(db, c)
	})

	app.Get("/authen/profile", func(c *fiber.Ctx) error {

		userid, err := getUserId(c)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
		}

		return apiProfile.GetProfile(db, c, userid)
	})

	app.Post("/authen/createprofile", func(c *fiber.Ctx) error {
		userid, err := getUserId(c)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
		}
		return apiProfile.CreateProfile(db, c, userid)
	})

	app.Patch("/authen/updateprofile", func(c *fiber.Ctx) error {
		userid, err := getUserId(c)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
		}
		return apiProfile.UpdateProfile(db, c, userid)
	})

	app.Listen(":8080")
}