package controllers

import (
	"go-mini-api/models"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var posts []models.Post = []models.Post{}

func GetAllPosts(c echo.Context) error {
	return c.JSON(http.StatusOK, models.BaseResponse[[]models.Post]{
		Status:  true,
		Message: "all posts",
		Data:    posts,
	})
}

func CreatePost(c echo.Context) error {
	var postRequest models.PostRequest

	if err := c.Bind(&postRequest); err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse[any]{
			Status:  false,
			Message: "invalid request",
		})
	}

	post := models.Post{
		ID:      uuid.NewString(),
		Title:   postRequest.Title,
		Content: postRequest.Content,
	}

	posts = append(posts, post)

	return c.JSON(http.StatusCreated, models.BaseResponse[models.Post]{
		Status:  true,
		Message: "post created",
		Data:    post,
	})
}
