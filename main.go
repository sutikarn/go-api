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

	token, err := jwt.ParseWithClaims(cookie, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized) // ส่งสถานะ unauthorized
	}
	
	// เช็ค claims ว่าถูกต้อง
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return c.SendStatus(fiber.StatusUnauthorized) // ส่งสถานะ unauthorized
	}

	// Print claims for debugging
	fmt.Println("Claims:", claims)

	// Extract user ID from claims and set it in the context
	if userID, exists := claims["user_id"]; exists {
		c.Locals("userID", userID) // Store user ID for later use
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized: user_id not found in claims"})
	}

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
		userID := c.Locals("userID").(int)
        fmt.Println(userID)
		return apiProfile.GetProfile(db, c, userID)
	})

	app.Post("/authen/createprofile", func(c *fiber.Ctx) error {
		return apiProfile.CreateProfile(db, c, 21)
	})

	app.Listen(":8080")
}

func main() {

	setDatabase()
	setPathApi()

}
