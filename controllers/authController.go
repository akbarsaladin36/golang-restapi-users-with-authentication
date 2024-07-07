package controllers

import (
	"fmt"
	"golang-rest-api-authentication/database"
	"golang-rest-api-authentication/helpers"
	"golang-rest-api-authentication/inputs"
	"golang-rest-api-authentication/models"
	"golang-rest-api-authentication/responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func HelloRoutes(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "testing routes succesfull",
	})
}

func RegisterUser(c *gin.Context) {
	var registerInput inputs.RegisterUserInput

	err := c.ShouldBindJSON(&registerInput)

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

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(registerInput.UserPassword), 14)

	newRegister := models.User{
		UserUsername:    registerInput.UserUsername,
		UserEmail:       registerInput.UserEmail,
		UserPassword:    string(hashedPassword),
		UserCreatedDate: time.Now(),
	}

	registerResponseRslt := registerResponse(newRegister)

	database.DB.Create(&newRegister)

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Data user baru telah berhasil ditambahkan!",
		"data":    registerResponseRslt,
	})
}

func LoginUser(c *gin.Context) {
	var loginInput inputs.LoginUserInput

	var user models.User

	err := c.ShouldBindJSON(&loginInput)

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

	checkUser := database.DB.Where("user_username = ?", loginInput.UserUsername).First(&user).Error

	if checkUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "Username ini belum terdaftar! Silakan daftar username anda terlebih dahulu!",
		})

		return
	}

	compareUserPassword := bcrypt.CompareHashAndPassword([]byte(user.UserPassword), []byte(loginInput.UserPassword))

	if compareUserPassword != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "Password yang dimasukkan masih salah! Silakan coba lagi!",
		})

		return
	}

	tokenString, err := helpers.GenerateJWTAuthentication(user.UserUsername, user.UserEmail)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "Proses generate token bermasalah! Silakan coba lagi!",
		})

		return
	}

	loginResponseRslt := loginResponse(user, tokenString)

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Sukses Login!",
		"data":    loginResponseRslt,
	})

}

func registerResponse(rgsRsps models.User) responses.RegisterResponse {
	return responses.RegisterResponse{
		UserUsername: rgsRsps.UserUsername,
		UserEmail:    rgsRsps.UserEmail,
		UserPassword: rgsRsps.UserPassword,
	}
}

func loginResponse(lgnRsps models.User, tokenString string) responses.LoginResponse {
	return responses.LoginResponse{
		UserUsername: lgnRsps.UserUsername,
		UserEmail:    lgnRsps.UserEmail,
		UserToken:    tokenString,
	}
}
