package model

import "time"

type MenuItem struct {
	OrganizationID uint       `json:"organization_id"`
	RestaurantID   uint       `json:"restaurant_id"`
	ID             uint       `json:"id" gorm:"primaryKey"`
	MenuID         uint       `json:"menu_id"`
	ItemName       string     `json:"item_name" validate:"required"`
	Description    string     `json:"description" validate:"omitempty"`
	Price          int        `json:"price" validate:"required"`
	RecipeID       uint       `json:"recipe_id"`
	Recipe         Recipe     `json:"-" gorm:"foreignKey:RecipeID;references:ID" validate:"-"`
	Allergens      []Allergen `json:"allergens" gorm:"foreignKey:ID"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}

type MenuItemResponse struct {
	OrganizationID uint       `json:"organization_id"`
	RestaurantID   uint       `json:"restaurant_id"`
	ID             uint       `json:"id"`
	MenuID         uint       `json:"menu_id"`
	ItemName       string     `json:"item_name"`
	Description    string     `json:"description"`
	Price          int        `json:"price"`
	RecipeID       uint       `json:"recipe_id"`
	Allergens      []Allergen `json:"allergens"`
}
