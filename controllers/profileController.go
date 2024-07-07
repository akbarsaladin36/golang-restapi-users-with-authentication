package controllers

import (
	"fmt"
	"golang-rest-api-authentication/database"
	"golang-rest-api-authentication/inputs"
	"golang-rest-api-authentication/middleware"
	"golang-rest-api-authentication/models"
	"golang-rest-api-authentication/responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetMyProfile(c *gin.Context) {
	var user models.User

	currentUserName, _, _ := middleware.CurrentUser(c)

	err := database.DB.Where("user_username = ?", currentUserName).First(&user).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "Data untuk profil ini tidak ada!",
		})
		return
	}

	profileRsps := getProfileResponse(user)

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Berhasil menampilkan data untuk username ini!",
		"data":    profileRsps,
	})
}

func UpdateMyProfile(c *gin.Context) {
	var user models.User

	currentUserName, _, _ := middleware.CurrentUser(c)

	checkUser := database.DB.Where("user_username = ?", currentUserName).First(&user).Error

	if checkUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "Data untuk profil ini tidak ada!",
		})
		return
	}

	var userInput inputs.ProfileUserInput

	err := c.ShouldBindJSON(&userInput)

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

	user.UserUsername = userInput.UserUsername
	user.UserEmail = userInput.UserEmail
	user.UserFirstName = userInput.UserFirstName
	user.UserLastName = userInput.UserLastName
	user.UserAddress = userInput.UserAddress
	user.UserPhoneNumber = userInput.UserPhoneNumber
	user.UserUpdatedDate = time.Now()

	updateUser := database.DB.Save(&user).Error

	if updateUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "Proses update data profil gagal!",
		})
		return
	}

	updateProfileRsps := getProfileResponse(user)

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Berhasil melakukan update data untuk profil ini!",
		"data":    updateProfileRsps,
	})

}

func getProfileResponse(userRsps models.User) responses.UserResponse {
	return responses.UserResponse{
		UserUsername:    userRsps.UserUsername,
		UserEmail:       userRsps.UserEmail,
		UserFirstName:   userRsps.UserFirstName,
		UserLastName:    userRsps.UserLastName,
		UserAddress:     userRsps.UserAddress,
		UserPhoneNumber: userRsps.UserPhoneNumber,
	}
}
