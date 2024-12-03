package model

import "time"

type Customer struct {
	OrganizationID uint         `json:"organization_id" gorm:"not null"`
	Organization   Organization `json:"-" gorm:"foreignKey:OrganizationID;references:ID" validate:"-"`
	RestaurantID   uint         `json:"restaurant_id" gorm:"not null"`
	Restaurant     Restaurant   `json:"-" gorm:"foreignKey:RestaurantID;references:ID" validate:"-"`
	ID             uint         `json:"id" gorm:"primaryKey"`
	Name           string       `json:"name" gorm:"not null;size:255" validate:"required"`
	PhoneNo        string       `json:"phone_no" gorm:"not null;size:255" validate:"required"`
	Email          string       `json:"email" gorm:"not null;size:255" validate:"required"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
	IsDeleted      bool         `json:"is_deleted" gorm:"not null;default:0"`
}

type CustomerResponse struct {
	OrganizationID uint   `json:"organization_id"`
	RestaurantID   uint   `json:"restaurant_id"`
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	PhoneNo        string `json:"phone_no"`
	Email          string `json:"email"`
}
