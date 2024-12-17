package model

import "time"

type MenuItem struct {
	OrganizationID uint           `json:"organization_id" gorm:"not null;foreignKey"`
	Organization   Organization   `json:"-" gorm:"foreignKey:OrganizationID;references:ID" validate:"-"`
	RestaurantID   uint           `json:"restaurant_id" gorm:"not null;foreignKey"`
	Restaurant     Restaurant     `json:"-" gorm:"foreignKey:RestaurantID;references:ID" validate:"-"`
	ID             uint           `json:"id" gorm:"primaryKey"`
	MenuID         uint           `json:"menu_id" gorm:"not null"`
	ItemName       string         `json:"item_name" gorm:"not null;size:255" validate:"required"`
	Description    string         `json:"description" gorm:"not null" validate:"omitempty"`
	Currency       string         `json:"currency" gorm:"not null;size:10" validate:"required"`
	Price          int            `json:"price" gorm:"not null;type:int(11)" validate:"required"`
	Serving        int            `json:"serving" gorm:"not null;type:int(11)" validate:"required"`
	IsAvailable    bool           `json:"is_available" gorm:"not null;default:1" validate:"required"`
	RecipeID       uint           `json:"recipe_id" gorm:"not null"`
	Recipe         Recipe         `json:"recipe" gorm:"foreignKey:RecipeID;references:ID" validate:"-"`
	MenuAllergens  []MenuAllergen `json:"menu_allergens" gorm:"foreignKey:MenuItemID"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	IsDeleted      bool           `json:"is_deleted" gorm:"not null;default:0"`
}

type MenuItemResponse struct {
	OrganizationID uint           `json:"organization_id"`
	RestaurantID   uint           `json:"restaurant_id"`
	ID             uint           `json:"id"`
	MenuID         uint           `json:"menu_id"`
	ItemName       string         `json:"item_name"`
	Description    string         `json:"description"`
	Currency       string         `json:"currency"`
	Price          string         `json:"price"`
	RecipeID       uint           `json:"recipe_id"`
	Serving        int            `json:"serving"`
	IsAvailable    bool           `json:"is_available"`
	MenuAllergens  []MenuAllergen `json:"menu_allergens"`
}
