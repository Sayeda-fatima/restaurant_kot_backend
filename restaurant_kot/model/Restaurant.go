package model

import "time"

type Restaurant struct {
	ID             uint         `json:"id" gorm:"primaryKey"`
	OrganizationID uint         `json:"organization_id" gorm:"not null" validate:"required"`
	Organization   Organization `json:"-" gorm:"foreignKey:OrganizationID;references:ID" validate:"-"`
	Name           string       `json:"name" gorm:"not null;size:255" validate:"required"`
	PhoneNo        string       `json:"phone_no" gorm:"not null;size:255" validate:"required"`
	Email          string       `json:"email" validate:"required" gorm:"not null;unique"`
	Address        string       `json:"address" validate:"omitempty"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
	IsDeleted      bool         `json:"is_deleted" gorm:"not null;default:0"`
}

type RestaurantResponse struct {
	ID             uint   `json:"id"`
	OrganizationID uint   `json:"organization_id"`
	Name           string `json:"name"`
	PhoneNo        string `json:"phone_no"`
	Email          string `json:"email"`
	Address        string `json:"address"`
}
