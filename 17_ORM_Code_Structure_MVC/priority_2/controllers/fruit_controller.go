package controllers

import (
	"assignment2/database"
	"assignment2/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// Fruityvice API response structure
type FruityviceResponse struct {
	Name       string `json:"name"`
	Family     string `json:"family"`
	Order      string `json:"order"`
	Nutritions struct {
		Calories      float64 `json:"calories"`
		Fat           float64 `json:"fat"`
		Sugar         float64 `json:"sugar"`
		Carbohydrates float64 `json:"carbohydrates"`
		Protein       float64 `json:"protein"`
	} `json:"nutritions"`
}

// Get all fruits
func GetFruits(c echo.Context) error {
	var fruits []models.Fruit
	database.DB.Find(&fruits)
	return c.JSON(http.StatusOK, fruits)
}

// Get fruit by ID
func GetFruitByID(c echo.Context) error {
	id := c.Param("id")
	var fruit models.Fruit
	if err := database.DB.Where("id = ?", id).First(&fruit).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Fruit not found"})
	}
	return c.JSON(http.StatusOK, fruit)
}

func CreateFruit(c echo.Context) error {
	fruit := new(models.Fruit)
	if err := c.Bind(fruit); err != nil {
		return err
	}

	// Make an API call to Fruityvice to get nutritional data
	fruitName := fruit.Name
	apiUrl := fmt.Sprintf("https://www.fruityvice.com/api/fruit/%s", fruitName)

	resp, err := http.Get(apiUrl)
	if err != nil || resp.StatusCode != 200 {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Could not retrieve nutritional data"})
	}
	defer resp.Body.Close()

	var fruityviceResponse FruityviceResponse
	if err := json.NewDecoder(resp.Body).Decode(&fruityviceResponse); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error decoding nutritional data"})
	}

	// Assign nutritional data to the fruit
	fruit.ID = uuid.New()
	fruit.Calories = fruityviceResponse.Nutritions.Calories
	fruit.Fat = fruityviceResponse.Nutritions.Fat
	fruit.Sugar = fruityviceResponse.Nutritions.Sugar
	fruit.Carbohydrates = fruityviceResponse.Nutritions.Carbohydrates
	fruit.Protein = fruityviceResponse.Nutritions.Protein

	// Save fruit to the database
	database.DB.Create(&fruit)
	return c.JSON(http.StatusCreated, fruit)
}

// Delete fruit by ID
func DeleteFruit(c echo.Context) error {
	id := c.Param("id")
	var fruit models.Fruit
	if err := database.DB.Where("id = ?", id).First(&fruit).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Fruit not found"})
	}
	database.DB.Delete(&fruit)
	return c.NoContent(http.StatusNoContent)
}
