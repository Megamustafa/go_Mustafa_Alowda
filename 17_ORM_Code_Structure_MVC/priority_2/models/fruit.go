package models

import "github.com/google/uuid"

// Fruit represents the fruit data structure
type Fruit struct {
	ID            uuid.UUID `json:"id" gorm:"type:uuid;primary_key;"`
	Name          string    `json:"name"`
	Price         int       `json:"price"`
	Calories      float64   `json:"calories"`
	Fat           float64   `json:"fat"`
	Sugar         float64   `json:"sugar"`
	Carbohydrates float64   `json:"carbohydrates"`
	Protein       float64   `json:"protein"`
}

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
