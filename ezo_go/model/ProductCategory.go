package model

import (
	"time"
)

type ProductCategory struct{
	//OrganizationID	
	ID			uint		`json:"id" gorm:"primaryKey"`
	Category	string		`json:"category" gorm:"unique"`
	CreatedAt	time.Time	`json:"created_at"`
	UpdatedAt	time.Time	`json:"updated_at"`
}

type ProductCategoryResponse struct{
	ID			uint		`json:"id" gorm:"primaryKey"`
	Category	string		`json:"category" gorm:"unique"`
	CreatedAt	time.Time	`json:"created_at"`
	UpdatedAt	time.Time	`json:"updated_at"`		
}