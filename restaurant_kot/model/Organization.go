package model

import (
	"time"
)

type Organization struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Image       string    `json:"required" validate:"required"`
	Name        string    `json:"name" gorm:"not null;size:255" validate:"required"`
	Address     string    `json:"address" gorm:"not null;size:255" validate:"required"`
	Email       string    `json:"email" gorm:"not null;unique" validate:"required"`
	PhoneNo     string    `json:"phone_no" gorm:"not null;size:255" validate:"required"`
	AccessGiven int       `json:"access_given" gorm:"not null;type:int(11)" validate:"required"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	IsDeleted   bool      `json:"is_deleted" gorm:"not null;default:0"`
}

type OrganizationResponse struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNo     string `json:"phone_no"`
	AccessGiven int    `json:"access_given"`
}
