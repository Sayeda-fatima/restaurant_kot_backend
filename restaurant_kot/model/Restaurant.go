package model

import "time"

type Restaurant struct {
	ID             uint         `json:"id" gorm:"primaryKey"`
	OrganizationID uint         `json:"organization_id" validate:"required"`
	Organization   Organization `json:"-" gorm:"foreignKey:OrganizationID;references:ID" validate:"-"`
	Name           string       `json:"name" validate:"required"`
	PhoneNo        string       `json:"phone_no" validate:"required"`
	Email          string       `json:"email" validate:"required" gorm:"unique"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
	IsDeleted      bool         `json:"is_deleted"`
}

type RestaurantResponse struct {
	ID             uint   `json:"id"`
	OrganizationID uint   `json:"organization_id"`
	Name           string `json:"name"`
	PhoneNo        string `json:"phone_no"`
	Email          string `json:"email"`
}
