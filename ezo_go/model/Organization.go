package model

import (
	"time"
)

type Organization struct{
	ID 			uint		`json:"id" gorm:"primaryKey"`
	Image		string		`json:"required" validate:"required"`
	Name 		string		`json:"name" validate:"required"`
	Address 	string		`json:"address" validate:"required"`
	PhoneNo		string		`json:"phone_no" validate:"required"`
	AccessGiven	int64		`json:"access_given" validate:"required"`
	CreatedAt	time.Time	`json:"created_at"`
	UpdatedAt	time.Time	`json:"updated_at"`
	IsDeleted	bool		`json:"is_deleted"`
}

type OrganizationResponse struct{
	ID 			uint		`json:"id" gorm:"primaryKey"`
	Name 		string		`json:"name"`
	Address		string		`json:"address"`
	PhoneNo		string		`json:"phone_no"`
	AccessGiven	int64		`json:"access_given"`
}