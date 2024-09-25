package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	host     = "localhost"  // or the Docker service name if running in another container
	port     = 5432         // default PostgreSQL port
	user     = "myuser"     // as defined in docker-compose.yml
	password = "mypassword" // as defined in docker-compose.yml
	dbname   = "mydatabase" // as defined in docker-compose.yml
)

func init() {

	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	fmt.Println([]byte(os.Getenv("SECRET_KEY")))

}

func main() {

	setDatabase()
	setPathApi()

}
