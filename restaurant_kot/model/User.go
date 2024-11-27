package model

import (
	"time"
)

type User struct {
	ID             uint         `json:"id" gorm:"primaryKey"`
	OrganizationID uint         `json:"organization_id" gorm:"foreignKey"`
	Organization   Organization `gorm:"foreignKey:OrganizationID;references:ID" validate:"-"`
	Name           string       `json:"name"`
	Email          string       `json:"email" gorm:"unique" validate:"required,email"`
	Password       string       `json:"password" validate:"required,min=8"`
	AccessType     string       `json:"access_type"`
	ApiToken       string       `json:"api_token"`
	RefreshToken   string       `json:"refresh_token"`
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
