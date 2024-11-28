package model

import "time"

type RestaurantTable struct {
	ID             uint         `json:"id" gorm:"primaryKey"`
	OrganizationID uint         `json:"organization_id"`
	Organization   Organization `json:"-" gorm:"foreignKey:OrganizationID;references:ID"`
	RestaurantID   uint         `json:"restaurant_id"`
	Restaurant     Restaurant   `json:"-" gorm:"foreignKey:RestaurantID;references:ID"`
	Capacity       int          `json:"capacity" validate:"required"`
	Status         string       `json:"status" validate:"required"` // in set: booked, free, occupied
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
	IsDeleted		bool		`json:"is_deleted"`
}

type RestaurantTableResponse struct {
	ID 				uint 		`json:"id"`
	OrganizationID 	uint		`json:"organization_id"`
	RestaurantID	uint		`json:"restaurant_id"`
	Capacity 		int			`json:"capacity"`
	Status			string		`json:"status"`
}
