package model

import "time"

type RecipeProduct struct{
	ID 			uint		`json:"id" gorm:"primaryKey"`
	RecipeID	uint		`json:"recipe_id" gorm:"not null" validate:"required"`
	Recipe		Recipe		`json:"-" gorm:"foreignKey:RecipeID;references:ID" validate:"-"`
	ProductID	uint		`json:"product_id" gorm:"not null" validate:"required"`
	Product		Product		`json:"-" gorm:"foreignKey:ProductID;references:ID" validate:"-"`
	Quantity	int			`json:"quantity" gorm:"not null;type:int(11)" validate:"required"`
	CreatedAt	time.Time	`json:"created_at"`
	UpdatedAt	time.Time	`json:"updated_at"`
	IsDeleted   bool      	`json:"is_deleted" gorm:"not null;default:0"`
}

type RecipeProductResponse struct{
	ID			uint		`json:"id"`
	RecipeID 	uint		`json:"recipe_id"`
	ProductID 	uint 		`json:"product_id"`
	Quantity	int			`json:"quantity"`
}
