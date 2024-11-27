package model

import "time"

type Recipe struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	Name        string     `json:"name" validate:"required"`
	Instruction string     `json:"instruction" validate:"required"`
	CookingTime string     `json:"cooking_time" validate:"required"`
	Serving     int        `json:"serving" validate:"required"`
	Products    []Product  `gorm:"many2many:recipe_products"`
	Allergens   []Allergen `json:"allergen" gorm:"foreignKey:ID"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type RecipeResponse struct {
	ID          uint       `json:"id"`
	Name        string     `json:"name"`
	Instruction string     `json:"instruction"`
	CookingTime string     `json:"cooking_time"`
	Serving     int        `json:"serving"`
	Allergens   []Allergen `json:"allergen"`
}
