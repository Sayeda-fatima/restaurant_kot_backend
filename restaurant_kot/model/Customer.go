package model

import "time"

type Customer struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" validate:"required"`
	PhoneNo   string    `json:"phone_no" validate:"required"`
	Email     string    `json:"email" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsDeleted bool      `json:"is_deleted"`
}

type CustomerResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	PhoneNo string `json:"phone_no"`
	Email   string `json:"email"`
}
