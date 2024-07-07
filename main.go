package main

import (
	"fmt"
	"golang-rest-api-authentication/database"
	"golang-rest-api-authentication/routes"

	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	database.ConnectDB()
	routes.ConnectRoutes()
}

func loadEnv() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Gagal mengaktifkan file .env")
	}
}
