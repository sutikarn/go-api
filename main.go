package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"  // or the Docker service name if running in another container
	port     = 5432         // default PostgreSQL port
	user     = "myuser"     // as defined in docker-compose.yml
	password = "mypassword" // as defined in docker-compose.yml
	dbname   = "mydatabase" // as defined in docker-compose.yml
)

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var db *sql.DB

func main() {
	// Connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open a connection
	sdb, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	db = sdb

	defer db.Close()
	// Check the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected!")

	app := fiber.New()

	app.Get("/product/:id", getProductHandler)
  app.Get("/products/", getProductsHandler)
	app.Post("/product/", createProductHandler)
  app.Put("/product/:id", updateProductHandler)
  app.Delete("product/:id",deleteProductHandler)

	app.Listen(":8080")

}

func getProductHandler(c *fiber.Ctx) error {
	produceId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	product, err := getProduct(produceId)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(product)
}

func getProductsHandler(c *fiber.Ctx) error {
	product, err := getProducts()
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(product)
}

func createProductHandler(c *fiber.Ctx) error {
	p := new(Product)
	if err := c.BodyParser(p); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err := createProduct(p)

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(p)
}

func updateProductHandler(c *fiber.Ctx) error {
  id := c.Params("id")
  p := new(Product)
  if err := c.BodyParser(p); err != nil {
    return err
  }

  // Update product in the database
  _, err := db.Exec("UPDATE products SET name = $1, price = $2 WHERE id = $3", p.Name, p.Price, id)
  if err != nil {
    return err
  }

  p.ID, _ = strconv.Atoi(id)
  return c.JSON(p)
}

func deleteProductHandler(c *fiber.Ctx) error {
  id := c.Params("id")

  // Delete product from database
  _, err := db.Exec("DELETE FROM products WHERE id = $1", id)
  if err != nil {
    return err
  }

  return c.SendStatus(fiber.StatusNoContent)
}
