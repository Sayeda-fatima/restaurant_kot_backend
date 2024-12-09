package model

type ProductStock struct{
	ID 			uint		`json:"id" gorm:"primaryKey"`
	OrganizationID uint           `json:"organization_id" gorm:"not null;foreignKey"`
	Organization   Organization   `json:"-" gorm:"foreignKey:OrganizationID;references:ID" validate:"-"`
	RestaurantID   uint           `json:"restaurant_id" gorm:"not null;foreignKey"`
	Restaurant     Restaurant     `json:"-" gorm:"foreignKey:RestaurantID;references:ID" validate:"-"`
	ProductID 		uint		`json:"product_id" gorm:"not null"`
	Product			Product		`json:"-" gorm:"foreignKey:ProductID;references:ID" validate:"-"`
	Quantity		int64		`json:"quantity" gorm:"not null" validate:"required"`
	//UnitCost 		
}