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
	app.Post("/api/v1/sign-up", func(c *fiber.Ctx) error {
		return apimeow.Signup(db, c)
	})

	app.Post("/api/v1/sign-in", func(c *fiber.Ctx) error {
		return apimeow.Signin(db, c)
	})

	app.Get("/home", func(c *fiber.Ctx) error {
		return apihome.GetHome(db, c)
	})

	///goauthen

	app.Use("/api/v1/profile", authRequired)

	app.Get("/api/v1/profile", func(c *fiber.Ctx) error {

		userid, err := getUserId(c)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
		}

		return apiProfile.GetProfile(db, c, userid)
	})

	app.Post("/api/v1/profile", func(c *fiber.Ctx) error {
		userid, err := getUserId(c)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
		}
		return apiProfile.CreateProfile(db, c, userid)
	})
	app.Patch("/api/v1/profile", func(c *fiber.Ctx) error {
		userid, err := getUserId(c)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
		}
		return apiProfile.UpdateProfile(db, c, userid)
	})

	app.Use("/api/v1/product/cart", authRequired)

	app.Post("/api/v1/product/cart", func(c *fiber.Ctx) error {
		userid, err := getUserId(c)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
		}
		return apicart.AddToCart(db, c, userid)
	})

	app.Get("/api/v1/product/cart", func(c *fiber.Ctx) error {
		userid, err := getUserId(c)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
		}
		return apicart.GetCart(db, c, userid)
	})

	app.Use("/api/v1/product/favorite", authRequired)

	app.Get("/api/v1/product/favorite", func(c *fiber.Ctx) error {
		userid, err := getUserId(c)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
		}
		return apifav.GetFavorite(db, c, userid)
	})

	app.Post("/api/v1/product/favorite", func(c *fiber.Ctx) error {
		userid, err := getUserId(c)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
		}
		return apifav.AddFavorite(db, c, userid)
	})

	app.Delete("/api/v1/product/favorite/:id", func(c *fiber.Ctx) error {
		userid, err := getUserId(c)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
		}
		return apifav.DeleteFavorite(db, c, userid)
	})

	app.Use("/api/v1/address", authRequired)

	app.Get("/api/v1/address", func(c *fiber.Ctx) error {
		userid, err := getUserId(c)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
		}
		return apiuser.GetAddress(db, c, userid)
	})

	app.Post("/api/v1/address", func(c *fiber.Ctx) error {
		userid, err := getUserId(c)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
		}
		return apiuser.CreateAddress(db, c, userid)
	})

	app.Patch("/api/v1/address", func(c *fiber.Ctx) error {
		userid, err := getUserId(c)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
		}
		return apiuser.UpdateAddress(db, c, userid)
	})

	app.Delete("/api/v1/address/:id", func(c *fiber.Ctx) error {
		userid, err := getUserId(c)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
		}
		return apiuser.DeleteAddress(db, c, userid)
	})

	app.Use("/api/v1/order", authRequired)

	app.Post("/api/v1/order", func(c *fiber.Ctx) error {
		userid, err := getUserId(c)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
		}
		return apiuser.CreateOrder(db, c, userid)
	})

	app.Get("/api/v1/order", func(c *fiber.Ctx) error {
		userid, err := getUserId(c)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
		}
		return apiuser.GetOrder(db, c, userid)
	})

	app.Get("/api/v1/product/banner", func(c *fiber.Ctx) error {
		return apihome.GetBanner(db, c)
	})

	app.Get("/api/v1/product/category", func(c *fiber.Ctx) error {
		return apihome.GetCategory(db, c)
	})

	app.Get("/api/v1/product/shoppingmall", func(c *fiber.Ctx) error {
		return apihome.GetShoppingMall(db, c)
	})

	app.Get("/api/v1/product", func(c *fiber.Ctx) error {
		return apihome.GetProduct(db, c)
	})

	app.Get("/api/v1/product/:id", func(c *fiber.Ctx) error {
		return apihome.GetProductById(db, c)
	})

	app.Get("/api/v1/product/category/:id", func(c *fiber.Ctx) error {
		return apihome.GetCategoryById(db, c)
	})

	app.Get("/api/v1/product/shoppingmall/:id", func(c *fiber.Ctx) error {
		return apihome.GetShoppingMallById(db, c)
	})

	app.Get("/api/v1/products/recommend", func(c *fiber.Ctx) error {
		return apihome.GetProductRecommend(db, c)
	})

	app.Listen(":8080")
}
