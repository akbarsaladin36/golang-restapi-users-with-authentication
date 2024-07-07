package helpers

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// var jwtKey = []byte("secretkey12345")

var jwtKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

type JWTClaim struct {
	UserUsername string `json:"user_username"`
	UserEmail    string `json:"user_email"`
	jwt.RegisteredClaims
}

func GenerateJWTAuthentication(username string, email string) (tokenString string, err error) {
	claims := &JWTClaim{
		UserUsername: username,
		UserEmail:    email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}
