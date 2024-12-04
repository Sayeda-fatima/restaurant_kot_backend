package model

import "time"

type Recipe struct {
	ID             uint         `json:"id" gorm:"primaryKey"`
	OrganizationID uint         `json:"organization_id" gorm:"not null"`
	Organization   Organization `json:"-" gorm:"foreignKey:OrganizationID;references:ID" validate:"-"`
	RestaurantID   uint         `json:"restaurant_id" gorm:"not null"`
	Restaurant     Restaurant   `json:"-" gorm:"foreignKey:RestaurantID;references:ID" validate:"-"`
	Name           string       `json:"name" gorm:"not null;size:255" validate:"required"`
	Instruction    string       `json:"instruction" gorm:"not null" validate:"required"`
	CookingTime    string       `json:"cooking_time" gorm:"not null;size:255" validate:"required"`
	Serving        int          `json:"serving" gorm:"not null;type:int(11)" validate:"required"`
	Products       []Product    `gorm:"many2many:recipe_products"`
	Allergens      []Allergen   `json:"allergen" gorm:"foreignKey:RecipeID"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
	IsDeleted      bool         `json:"is_deleted" gorm:"not null;default:0"`
}

type RecipeResponse struct {
	ID          uint       `json:"id"`
	Name        string     `json:"name"`
	Instruction string     `json:"instruction"`
	CookingTime string     `json:"cooking_time"`
	Serving     int        `json:"serving"`
	Allergens   []Allergen `json:"allergen"`
}
