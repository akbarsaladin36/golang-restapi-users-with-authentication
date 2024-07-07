package controllers

import (
	"fmt"
	"golang-rest-api-authentication/database"
	"golang-rest-api-authentication/inputs"
	"golang-rest-api-authentication/models"
	"golang-rest-api-authentication/responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func GetAllUsers(c *gin.Context) {
	var users []models.User

	err := database.DB.Find(&users).Error

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "200",
			"message": "Data untuk keseluruhan user tidak ada! Silakan buat user baru!",
		})
		return
	}

	var usersRsps []responses.UserResponse

	for _, user := range users {

		userRsps := getUserResponse(user)

		usersRsps = append(usersRsps, userRsps)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Berhasil menampilkan data semua users!",
		"data":    usersRsps,
	})

}

func GetUser(c *gin.Context) {
	username := c.Param("username")

	var user models.User

	err := database.DB.Where("user_username = ?", username).First(&user).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "Data untuk username ini tidak ada!",
		})

		return
	}

	userRsps := getUserResponse(user)

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Berhasil menampilkan data untuk username ini!",
		"data":    userRsps,
	})
}

func CreateUser(c *gin.Context) {
	var newUserInput inputs.NewUserInput

	err := c.ShouldBindJSON(&newUserInput)

	if err != nil {
		errorMessages := []string{}

		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": errorMessages,
		})

		return
	}

	var user models.User

	checkUser := database.DB.Where("user_username = ?", newUserInput.UserUsername).First(&user).Error

	if checkUser == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "Data username ini sudah terdaftar dalam aplikasi ini! Silakan buat username baru yang berbeda!",
		})

		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(newUserInput.UserPassword), 14)

	newUser := models.User{
		UserUsername:    newUserInput.UserUsername,
		UserPassword:    string(hashedPassword),
		UserEmail:       newUserInput.UserEmail,
		UserFirstName:   newUserInput.UserFirstName,
		UserLastName:    newUserInput.UserLastName,
		UserAddress:     newUserInput.UserAddress,
		UserPhoneNumber: newUserInput.UserPhoneNumber,
		UserCreatedDate: time.Now(),
	}

	createUser := database.DB.Create(&newUser).Error

	if createUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "Gagal membuat user baru!",
		})

		return
	}

	newUserRsps := getUserResponse(newUser)

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Berhasil menampilkan data untuk username ini!",
		"data":    newUserRsps,
	})

}

func DeleteUser(c *gin.Context) {
	username := c.Param("username")

	var user models.User

	checkUser := database.DB.Where("user_username = ?", username).First(&user).Error

	if checkUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "Data username ini tidak terdaftar!",
		})

		return
	}

	database.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Data username ini berhasil dihapus!",
	})

}

func getUserResponse(userRsps models.User) responses.UserResponse {
	return responses.UserResponse{
		UserUsername:    userRsps.UserUsername,
		UserEmail:       userRsps.UserEmail,
		UserFirstName:   userRsps.UserFirstName,
		UserLastName:    userRsps.UserLastName,
		UserAddress:     userRsps.UserAddress,
		UserPhoneNumber: userRsps.UserPhoneNumber,
	}
}
