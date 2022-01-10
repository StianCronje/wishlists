package main

import (
	"net/http"
	"wishlists/database"
	"wishlists/routes"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	database.Connect()
	defer database.Disconnect()

	r := routes.Setup()

	http.ListenAndServe(":5000", r)
}