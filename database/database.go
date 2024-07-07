package database

import (
	"fmt"
	"golang-rest-api-authentication/models"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	host := os.Getenv("DB_HOSTNAME")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	databaseName := os.Getenv("DB_NAME")

	// dsn := "root:root@tcp(127.0.0.1:3306)/db_rest_api_users?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, databaseName)

	// fmt.Println(dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Connection to database is failed")
	}

	fmt.Println("Database is succesfully connected!")

	db.AutoMigrate(&models.User{})

	DB = db
}
