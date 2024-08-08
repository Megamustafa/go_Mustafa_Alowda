package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Food struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
}

var foods []Food

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Seed some data
	foods = append(foods, Food{ID: uuid.New().String(), Name: "Nasi Goreng", Price: 12000, Description: "delicious"})
	foods = append(foods, Food{ID: uuid.New().String(), Name: "Ayam Goreng", Price: 22000, Description: "delicious"})

	// Routes
	e.GET("/api/v1/foods", getFoods)
	e.GET("/api/v1/foods/:id", getFood)
	e.POST("/api/v1/foods", createFood)
	e.PUT("/api/v1/foods/:id", updateFood)
	e.DELETE("/api/v1/foods/:id", deleteFood)

	// Start server
	e.Logger.Fatal(e.Start(":8000"))
}

func getFoods(c echo.Context) error {
	return c.JSON(http.StatusOK, foods)
}

func getFood(c echo.Context) error {
	id := c.Param("id")
	for _, item := range foods {
		if item.ID == id {
			return c.JSON(http.StatusOK, item)
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"message": "Food not found"})
}

func createFood(c echo.Context) error {
	var food Food
	if err := c.Bind(&food); err != nil {
		return err
	}
	food.ID = uuid.New().String()
	foods = append(foods, food)
	return c.JSON(http.StatusCreated, food)
}

func updateFood(c echo.Context) error {
	id := c.Param("id")
	var updatedFood Food
	if err := c.Bind(&updatedFood); err != nil {
		return err
	}
	for index, item := range foods {
		if item.ID == id {
			updatedFood.ID = id
			foods[index] = updatedFood
			return c.JSON(http.StatusOK, updatedFood)
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"message": "Food not found"})
}

func deleteFood(c echo.Context) error {
	id := c.Param("id")
	for index, item := range foods {
		if item.ID == id {
			foods = append(foods[:index], foods[index+1:]...)
			return c.JSON(http.StatusOK, map[string]string{"message": "Food deleted"})
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"message": "Food not found"})
}
