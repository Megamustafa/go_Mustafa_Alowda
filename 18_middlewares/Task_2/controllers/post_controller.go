package controllers

import (
	"net/http"
	"post-manager/models"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type H map[string]interface{}

func GetPosts(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var posts []models.Post
		db.Find(&posts)
		return c.JSON(http.StatusOK, posts)
	}
}

func GetPost(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		var post models.Post
		db.First(&post, id)
		if post.ID == 0 {
			return c.JSON(http.StatusNotFound, H{"message": "Post not found"})
		}
		return c.JSON(http.StatusOK, post)
	}
}

func CreatePost(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		post := new(models.Post)
		if err := c.Bind(post); err != nil {
			return err
		}
		db.Create(&post)
		return c.JSON(http.StatusCreated, post)
	}
}

func UpdatePost(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		var post models.Post
		db.First(&post, id)
		if post.ID == 0 {
			return c.JSON(http.StatusNotFound, H{"message": "Post not found"})
		}
		if err := c.Bind(&post); err != nil {
			return err
		}
		db.Save(&post)
		return c.JSON(http.StatusOK, post)
	}
}

func DeletePost(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		var post models.Post
		db.First(&post, id)
		if post.ID == 0 {
			return c.JSON(http.StatusNotFound, H{"message": "Post not found"})
		}
		db.Delete(&post)
		return c.NoContent(http.StatusNoContent)
	}
}
