package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type JwtCustomClaims struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func CreateJWTToken(id uint, email string) *jwt.Token {
	claims := &JwtCustomClaims{
		id,
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
			Issuer:    "Compnouron",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token
}

func CreateSignedJWTToken(id uint, email string) (string, error) {
	token := CreateJWTToken(id, email)
	encodedToken, err := token.SignedString([]byte(os.Getenv("SIGNING_KEY")))
	if err != nil {
		return "", err
	}
	return encodedToken, nil
}

func GetUserDetails(c echo.Context) (uint, string) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	userID := claims.ID
	email := claims.Email

	return userID, email
}
