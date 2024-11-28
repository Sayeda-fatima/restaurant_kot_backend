package model

import "time"

type Menu struct {
	ID             uint         `json:"id" gorm:"primaryKey"`
	OrganizationID uint         `json:"organization_id" gorm:"not null"`
	Organization   Organization `json:"-" gorm:"foreignKey:OrganizationID;references:ID" validate:"-"`
	RestaurantID   uint         `json:"restaurant_id" gorm:"not null"`
	Restaurant     Restaurant   `json:"-" gorm:"foreignKey:RestaurantID;references:ID" validate:"-"`
	Name           string       `json:"name" gorm:"not null;size:255" validate:"required"`
	MenuItems      []MenuItem   `json:"menu_items" gorm:"foreignKey:MenuID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
	IsDeleted      bool         `json:"is_deleted" gorm:"not null;default:0"`
}

type MenuResponse struct {
	ID             uint       `json:"id"`
	OrganizationID uint       `json:"organization_id"`
	RestaurantID   uint       `json:"restaurant_id"`
	Name           string     `json:"name"`
	MenuItems      []MenuItem `json:"menu_items"`
}
