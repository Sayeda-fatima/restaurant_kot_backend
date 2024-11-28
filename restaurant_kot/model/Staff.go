package model

import "time"

type Staff struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" validate:"required"`
	PhoneNo   string    `json:"phone_no" validate:"required"`
	Email     string    `json:"email" validate:"required"`
	Role      string    `json:"role" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsDeleted bool      `json:"is_deleted"`
}

type StaffResponse struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Name    string `json:"name"`
	PhoneNo string `json:"phone_no"`
	Email   string `json:"email"`
	Role    string `json:"role"`
}
