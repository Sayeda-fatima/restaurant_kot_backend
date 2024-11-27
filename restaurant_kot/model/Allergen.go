package model

import "time"

type Allergen struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	AllergenName string    `json:"allergen_name" validate:"required"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type AllergenResponse struct {
	ID           uint   `json:"id"`
	AllergenName string `json:"allergen_name"`
}
