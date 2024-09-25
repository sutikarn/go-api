package main

import (
	"fmt"
	"log"
	"os"
	"time"

	apihome "nack/api/homeapi"
	apiProfile "nack/api/profileapi"
	apimeow "nack/api/signin_signup"

	"nack/loaddata"
	model "nack/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	host     = "localhost"  // or the Docker service name if running in another container
	port     = 5432         // default PostgreSQL port
	user     = "myuser"     // as defined in docker-compose.yml
	password = "mypassword" // as defined in docker-compose.yml
	dbname   = "mydatabase" // as defined in docker-compose.yml
)

var db *gorm.DB

func init() {

	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	fmt.Println([]byte(os.Getenv("SECRET_KEY")))

}

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

func setDatabase() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// New logger for detailed SQL logging
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Enable color
		},
	)

	dbc, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	db = dbc
	if err != nil {
		panic("failed to connect to database")
	}

	err = db.AutoMigrate(
		&model.User{},
		&model.Profile{},
		&model.Address{},
		&model.Category{},
		&model.Mall{},
		&model.Product{},
		&model.Order{},
		&model.Banner{},
		&model.Cart{},
		&model.Favorite{},
	)
	if err != nil {
		fmt.Println(err)
	}

	// loadDataBase(db)

	fmt.Println("Database migration complete")
}

func loadDataBase(db *gorm.DB) {
	loaddata.LoadData(db)
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
		return apiProfile.CreateProfile(db, c, 21)
	})

	app.Listen(":8080")
}

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

func main() {

	setDatabase()
	setPathApi()

}
