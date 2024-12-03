package model

import "time"

type RestaurantTable struct {
	ID             uint         `json:"id" gorm:"primaryKey"`
	TableNo        int          `json:"table_no" gorm:"not null;type:int(11)"`
	OrganizationID uint         `json:"organization_id" gorm:"not null"`
	Organization   Organization `json:"-" gorm:"foreignKey:OrganizationID;references:ID" validate:"-"`
	RestaurantID   uint         `json:"restaurant_id" gorm:"not null"`
	Restaurant     Restaurant   `json:"-" gorm:"foreignKey:RestaurantID;references:ID" validate:"-"`
	Capacity       int          `json:"capacity" gorm:"not null;type:int(11)" validate:"required"`
	Status         string       `json:"status" gorm:"not null;type:enum('booked','free','occupied')" validate:"required"` // in set: booked, free, occupied
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
	IsDeleted      bool         `json:"is_deleted" gorm:"not null;default:0"`
}

type RestaurantTableResponse struct {
	ID             uint   `json:"id"`
	OrganizationID uint   `json:"organization_id"`
	RestaurantID   uint   `json:"restaurant_id"`
	TableNo        int    `json:"table_no"`
	Capacity       int    `json:"capacity"`
	Status         string `json:"status"`
}
