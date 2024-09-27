package main

import (
	"fmt"
	apicart "nack/api/cartapi"
	apifav "nack/api/favapi"
	apihome "nack/api/homeapi"
	apiProfile "nack/api/profileapi"
	apimeow "nack/api/signin_signup"
	apiuser "nack/api/userapi"
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

func setPathApi() {
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

	app.Get("/banner", func(c *fiber.Ctx) error {
		return apihome.GetBanner(db, c)
	})

	app.Get("/category", func(c *fiber.Ctx) error {
		return apihome.GetCategory(db, c)
	})

	app.Get("/shoppingmall", func(c *fiber.Ctx) error {
		return apihome.GetShoppingMall(db, c)
	})

	app.Get("/product", func(c *fiber.Ctx) error {
		return apihome.GetProduct(db, c)
	})

	app.Get("/product/:id", func(c *fiber.Ctx) error {
		return apihome.GetProductById(db, c)
	})

	app.Get("/category/:id", func(c *fiber.Ctx) error {
		return apihome.GetCategoryById(db, c)
	})

	app.Get("/shoppingmall/:id", func(c *fiber.Ctx) error {
		return apihome.GetShoppingMallById(db, c)
	})

	app.Get("/productRecommend", func(c *fiber.Ctx) error {
		return apihome.GetProductRecommend(db, c)
	})

	///goauthen

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

	app.Post("/authen/addtocart", func(c *fiber.Ctx) error {
		userid, err := getUserId(c)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
		}
		return apicart.AddToCart(db, c, userid)
	})

	app.Get("/authen/getcart", func(c *fiber.Ctx) error {
		userid, err := getUserId(c)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
		}
		return apicart.GetCart(db, c, userid)
	})

	app.Get("/authen/favorite", func(c *fiber.Ctx) error {
		userid, err := getUserId(c)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
		}
		return apifav.GetFavorite(db, c, userid)
	})

	app.Post("/authen/favorite", func(c *fiber.Ctx) error {
		userid, err := getUserId(c)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
		}
		return apifav.AddFavorite(db, c, userid)
	})

	app.Delete("/authen/favorite/:id", func(c *fiber.Ctx) error {
		userid, err := getUserId(c)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
		}
		return apifav.DeleteFavorite(db, c, userid)
	})

	app.Get("/authen/address", func(c *fiber.Ctx) error {
		userid, err := getUserId(c)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
		}
		return apiuser.GetAddress(db, c, userid)
	})

	app.Post("/authen/address", func(c *fiber.Ctx) error {
		userid, err := getUserId(c)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
		}
		return apiuser.CreateAddress(db, c, userid)
	})

	app.Patch("/authen/address", func(c *fiber.Ctx) error {
		userid, err := getUserId(c)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
		}
		return apiuser.UpdateAddress(db, c, userid)
	})

	app.Delete("/authen/address/:id", func(c *fiber.Ctx) error {
		userid, err := getUserId(c)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
		}
		return apiuser.DeleteAddress(db, c, userid)
	})



	app.Listen(":8080")
}
