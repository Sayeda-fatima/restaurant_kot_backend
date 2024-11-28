package model

import "time"

type Allergen struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	RecipeID     uint      `json:"recipe_id"`
	MenuItemID   uint      `json:"menu_item_id"`
	AllergenName string    `json:"allergen_name" validate:"required"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	IsDeleted    bool      `json:"is_deleted"`
}

type AllergenResponse struct {
	ID           uint   `json:"id"`
	RecipeID     uint   `json:"recipe_id"`
	MenuItemID   uint   `json:"menu_item_id"`
	AllergenName string `json:"allergen_name"`
}
