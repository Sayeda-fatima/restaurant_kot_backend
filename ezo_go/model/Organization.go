package model

import (
	"time"
)

type Organization struct{
	ID 			uint		`json:"id" gorm:"primaryKey"`
	Name 		string		`json:"name"`
	AccessGiven	int64		`json:"access_given"`
	CreatedAt	time.Time	`json:"created_at"`
	UpdatedAt	time.Time	`json:"updated_at"`
	IsDeleted	bool		`json:"is_deleted"`
}

type OrganizationResponse struct{
	ID 			uint		`json:"id" gorm:"primaryKey"`
	Name 		string		`json:"name"`
	AccessGiven	int64		`json:"access_given"`
}