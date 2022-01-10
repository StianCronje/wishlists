package database

import (
	"fmt"
	"os"
	"wishlists/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	
	DB_HOST := os.Getenv("DB_HOST")
	DB_USER := os.Getenv("DB_USER")
	DB_PASS := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")

	dsn := DB_USER + ":" + DB_PASS + "@tcp(" + DB_HOST + ")/" + DB_NAME
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connect to database.")
	}

	connection.AutoMigrate(&models.User{})
	connection.AutoMigrate(&models.WishItem{})

	DB = connection
}

func Disconnect() {
	db, err := DB.DB()
	if err != nil {
		fmt.Println("Could not close database.")
	}

	db.Close()
}