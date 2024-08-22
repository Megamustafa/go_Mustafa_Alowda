package controllers

import (
	"net/http"
	"post-manager/models"
	_ "post-manager/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := new(models.User)
		if err := c.Bind(user); err != nil {
			return err
		}
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		user.Password = string(hashedPassword)
		db.Create(&user)
		return c.JSON(http.StatusCreated, user)
	}
}

func LoginUser(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := new(models.User)
		if err := c.Bind(user); err != nil {
			return err
		}
		var dbUser models.User
		db.Where("email = ?", user.Email).First(&dbUser)
		if dbUser.ID == 0 {
			return echo.ErrUnauthorized
		}
		if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)); err != nil {
			return echo.ErrUnauthorized
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": dbUser.ID,
			"exp":     time.Now().Add(time.Hour * 72).Unix(),
		})
		tokenString, _ := token.SignedString([]byte("secret"))
		return c.JSON(http.StatusOK, map[string]string{
			"token": tokenString,
		})
	}
}
