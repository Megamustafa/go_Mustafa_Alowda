package middlewares

import (
	"strings"
	_ "time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func LogMiddleware(e *echo.Echo) {
	e.Use(middleware.Logger())
}

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return echo.ErrUnauthorized
		}
		tokenString = strings.Split(tokenString, " ")[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})
		if err != nil || !token.Valid {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}

func RateLimiterMiddleware(e *echo.Echo) {
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))
}
