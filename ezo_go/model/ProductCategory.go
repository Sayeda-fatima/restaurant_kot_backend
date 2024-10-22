package model

import (
	"time"
)

type ProductCategory struct {
	OrganizationID uint         `json:"organization_id"`
	Organization   Organization `gorm:"foreignKey:OrganizationID;references:ID" json:"-" validate:"-"`
	ID             uint         `json:"id" gorm:"primaryKey"`
	Category       string       `json:"category" gorm:"unique" validate:"required"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
	IsDeleted      bool         `json:"is_deleted"`
}

type ProductCategoryResponse struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Category  string    `json:"category" gorm:"unique"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
