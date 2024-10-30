package model

import (
	"time"
)

type Product struct {
	OrganizationID                uint            `json:"organization_id"`
	Organization                  Organization    `gorm:"foreignKey:OrganizationID;references:ID" json:"-" validate:"-"`
	ID                            uint            `json:"id" gorm:"primaryKey"`
	Name                          string          `json:"name" validate:"required"`
	Image                         string          `json:"image" validate:"required"`
	SellPrice                     float64         `json:"sell_price" validate:"required"`
	MeasuringUnit                 string          `json:"measuring_unit" validate:"required"`
	CategoryID                    uint            `json:"category_id" validate:"required"`
	Category                      ProductCategory `gorm:"foreignKey:CategoryID;references:ID" json:"-" validate:"-"`
	Quantity                      int64           `json:"quantity" validate:"required"`
	Mrp                           float64         `json:"mrp" validate:"required"`
	PurchasePrice                 float64         `json:"purchase_price" validate:"required"`
	AcSalePrice                   float64         `json:"ac_sale_price" validate:"required"`
	NonAcSalePrice                float64         `json:"non_ac_sale_price" validate:"required"`
	OnlineDeliverySellPrice       float64         `json:"online_delivery_sell_price" validate:"required"`
	OnlineSellPrice               float64         `json:"online_sell_price" validate:"required"`
	Tax                           string          `json:"tax" validate:"omitempty"`
	PriceWithTax                  string          `json:"price_with_tax" validate:"omitempty"`
	Cess                          int64           `json:"cess" validate:"omitempty"`
	HsnCode                       string          `json:"hsn_code" validate:"omitempty"`
	Description                   string          `json:"description" validate:"omitempty"`
	LowStockAlert                 int64           `json:"low_stock_alert" validate:"omitempty"`
	StorageLocation               string          `json:"storage_location" validate:"omitempty"`
	BulkPurchaseUnit              string          `json:"bulk_purchase_unit" validate:"omitempty"`
	RetailSaleUnitPerBulkPurchase float64         `json:"retail_sale_unit_per_bulk_purchase" validate:"omitempty"`
	BulkPurchaseUnitPerRetailSale float64         `json:"bulk_purchase_unit_per_retail_sale" validate:"omitempty"`
	ExpiryDate                    string          `json:"expiry_date" validate:"omitempty"`
	ShowProductOnlineStore        string          `json:"show_product_online_store" validate:"omitempty"`
	CreatedAt                     time.Time       `json:"created_at"`
	UpdatedAt                     time.Time       `json:"updated_at"`
	IsDeleted                     bool            `json:"is_deleted"`
}

type ProductResponse struct {
	OrganizationID                uint    `json:"organization_id"`
	ID                            uint    `json:"id" gorm:"primaryKey"`
	Name                          string  `json:"name"`
	Image                         string  `json:"image"`
	SellPrice                     float64 `json:"sell_price"`
	MeasuringUnit                 string  `json:"measuring_unit"`
	CategoryID                    uint    `json:"category_id"`
	Quantity                      int64   `json:"quantity"`
	Mrp                           float64 `json:"mrp"`
	PurchasePrice                 float64 `json:"purchase_price"`
	AcSalePrice                   float64 `json:"ac_sale_price"`
	NonAcSalePrice                float64 `json:"non_ac_sale_price"`
	OnlineDeliverySellPrice       float64 `json:"online_delivery_sell_price"`
	OnlineSellPrice               float64 `json:"online_sell_price"`
	Tax                           string  `json:"tax"`
	PriceWithTax                  string  `json:"price_with_tax"`
	Cess                          int64   `json:"cess"`
	HsnCode                       string  `json:"hsn_code"`
	Description                   string  `json:"description"`
	LowStockAlert                 int64   `json:"low_stock_alert"`
	StorageLocation               string  `json:"storage_location"`
	BulkPurchaseUnit              string  `json:"bulk_purchase_unit"`
	RetailSaleUnitPerBulkPurchase float64 `json:"retail_sale_unit_per_bulk_purchase"`
	BulkPurchaseUnitPerRetailSale float64 `json:"bulk_purchase_unit_per_retail_sale"`
	ExpiryDate                    string  `json:"expiry_date"`
	ShowProductOnlineStore        string  `json:"show_product_online_store"`
}
