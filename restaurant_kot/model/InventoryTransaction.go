package model

import "time"

type InventoryTransaction struct {
	ID              uint         `json:"id" gorm:"primaryKey"`
	OrganizationID  uint         `json:"organization_id" gorm:"not null"`
	Organization    Organization `json:"-" gorm:"foreignKey:OrganizationID;references:ID" validate:"-"`
	RestaurantID    uint         `json:"restaurant_id" gorm:"not null"`
	Restaurant      Restaurant   `json:"-" gorm:"foreignKey:RestaurantID;references:ID" validate:"-"`
	ProductID       uint         `json:"product_id" gorm:"not null"`
	Product         Product      `json:"product" gorm:"foreignKey:ProductID;references:ID"`
	StockBeforeUpdate	float64		`json:"stock_before_update" gorm:"not null"`
	StockUnitPriceBeforeUpdate	float64	`json:"stock_unit_price_before_update" gorm:"not null"`
	TransactionType string       `json:"transaction_type" gorm:"not null;type:enum('initial_stock','purchase','sale','adjustment','waste')"` // enum: initial_stock, purchase, sale, adjustment, waste
	Quantity        float64      `json:"quantity" gorm:"not null"`                                                           // Positive for additions, negative for reductions
	UnitCost        float64      `json:"unit_cost" gorm:"not null"`                                                          // Cost per unit (for purchases)
	TotalCost       float64      `json:"total_cost"`                                                                         // Total cost for the transaction (auto-calculated)
	Reason          string       `json:"reason" gorm:"type:varchar(255)"`                                                    // Reason for adjustment/waste e.g., stock correction, spoilage
	RecordedAt      time.Time    `json:"recorded_at" gorm:"not null"`                                                        // Timestamp of transaction
	CreatedAt       time.Time    `json:"created_at"`
	UpdatedAt       time.Time    `json:"updated_at"`
}

type CurrentInventory struct{
	ID				uint			`json:"id" gorm:"primaryKey"`
	OrganizationID  uint         	`json:"organization_id" gorm:"not null"`
	Organization    Organization 	`json:"-" gorm:"foreignKey:OrganizationID;references:ID" validate:"-"`
	RestaurantID    uint         	`json:"restaurant_id" gorm:"not null"`
	Restaurant      Restaurant   	`json:"-" gorm:"foreignKey:RestaurantID;references:ID" validate:"-"`
	InventoryValue	float64			`json:"inventory_value" gorm:"not null"`
	CreatedAt		time.Time		`json:"created_at"`
	UpdatedAt		time.Time		`json:"updated_at"`
}