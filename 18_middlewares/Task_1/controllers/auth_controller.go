package controllers

import (
	"net/http"
	"package-api/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

var dbU *gorm.DB

func Register(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	// Hash password here (for simplicity not hashed)
	dbU.Create(&user)
	return c.JSON(http.StatusCreated, user)
}

func Login(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	// Find the user in DB (password comparison skipped)
	u := models.User{}
	dbU.Where("email = ?", user.Email).First(&u)
	if u.Email == "" {
		return c.JSON(http.StatusUnauthorized, "Invalid credentials")
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = u.Email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}
