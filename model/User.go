package model

import (
	"time"
)

type User struct {
	ID             uint         `json:"id" gorm:"primaryKey"`
	OrganizationID uint         `json:"organization_id" gorm:"not null"`
	Organization   Organization `gorm:"foreignKey:OrganizationID;references:ID" validate:"-"`
	RestaurantID   uint         `json:"restaurant_id" gorm:"not null"`
	Restaurant     Restaurant   `json:"-" gorm:"foreignKey:RestaurantID;references:ID" validate:"-"`
	Name           string       `json:"name" gorm:"not null;size:255"`
	Email          string       `json:"email" gorm:"not null;unique" validate:"required,email"`
	Password       string       `json:"password" gorm:"not null;size:255" validate:"required,min=8"`
	AccessType     string       `json:"access_type" gorm:"not null;type:enum('admin','manager','staff')"`
	ApiToken       string       `json:"api_token" gorm:"size:255"`
	RefreshToken   string       `json:"refresh_token" gorm:"size:255"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
}

type UserResponse struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Email string `json:"email" gorm:"unique"`
}

type PasswordResetToken struct {
	Email     string    `json:"email" validate:"required"`
	Token     string    `json:"token" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
}
