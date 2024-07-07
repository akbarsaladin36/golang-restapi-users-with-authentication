package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// var jwtKey = []byte("secretkey12345")

var jwtKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

type AuthMiddlewareClaim struct {
	UserUsername string `json:"user_username"`
	UserEmail    string `json:"user_email"`
	jwt.RegisteredClaims
}

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "400",
				"message": "Header authorization tidak ada! Silakan login terlebih dahulu!",
			})
			c.Abort()
			return
		}

		authHeaderParts := strings.Split(authHeader, " ")

		if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "400",
				"message": "Terjadi kesalahan format header authorization! Silakan coba lagi!",
			})
			c.Abort()
			return
		}

		tokenString := authHeaderParts[1]

		claims := &AuthMiddlewareClaim{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrTokenExpired {
				c.JSON(http.StatusUnauthorized, gin.H{
					"status":  "400",
					"message": "Token expired!",
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "400",
				"message": "Token tidak sesuai!",
			})
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "400",
				"message": "Token tidak sesuai!",
			})
			c.Abort()
			return
		}

		c.Set("auth_user_username", claims.UserUsername)
		c.Set("auth_user_email", claims.UserEmail)

		c.Next()
	}
}

func IsAdminAccess() gin.HandlerFunc {
	return func(c *gin.Context) {
		userName, _ := c.Get("auth_user_username")

		if userName != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "400",
				"message": "URL ini hanya bisa diakses oleh admin!",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func CurrentUser(c *gin.Context) (string, string, bool) {
	userName, userNameExists := c.Get("auth_user_username")
	userEmail, userEmailExists := c.Get("auth_user_email")

	if !userNameExists || !userEmailExists {
		return "", "", false
	}

	userNameString, userNameOk := userName.(string)
	userEmailString, userEmailOk := userEmail.(string)

	if !userNameOk || !userEmailOk {
		return "", "", false
	}

	return userNameString, userEmailString, true
}
