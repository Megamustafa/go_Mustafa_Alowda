package middleware

import (
	_ "net/http"
	_ "strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

var jwtKey = []byte("secret")

func GenerateToken(userID string) (string, error) {
	claims := &jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func JWTMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: jwtKey,
	})
}
