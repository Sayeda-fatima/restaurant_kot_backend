package model

import "time"

type ProductImage struct {
	ID             uint         `json:"id" gorm:"primaryKey"`
	OrganizationID uint         `json:"organization_id"`
	Organization   Organization `json:"-" gorm:"foreignKey:OrganizationID;references:ID" validate:"-"`
	ProductID      uint         `json:"product_id" validate:"required"`
	Product        Product      `json:"-" gorm:"foreignKey:ProductID;references:ID" validate:"-"`
	Url            string       `json:"url" validate:"required"`
	IsDeleted      bool         `json:"is_deleted" validate:"-"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
}

type ProductImageResponse struct {
	ID             uint   `json:"id"`
	OrganizationID uint   `json:"organization_id"`
	ProductID      uint   `json:"product_id"`
	Url            string `json:"url"`
}
