package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"nack/loaddata"
	model "nack/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

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
