package model

import (
	"time"
)

type Product struct {
	OrganizationID                uint            `json:"organization_id"`
	Organization                  Organization    `gorm:"foreignKey:OrganizationID;references:ID" json:"-"`
	ID                            uint            `json:"id" gorm:"primaryKey"`
	Name                          string          `json:"name"`
	Image                         string          `json:"image"`
	SellPrice                     float64         `json:"sell_price"`
	MeasuringUnit                 string          `json:"measuring_unit"`
	CategoryID                    uint            `json:"category_id"`
	Category                      ProductCategory `gorm:"foreignKey:CategoryID;references:ID" json:"-"`
	Quantity                      int64           `json:"quantity"`
	Mrp                           float64         `json:"mrp"`
	PurchasePrice                 float64         `json:"purchase_price"`
	AcSalePrice                   float64         `json:"ac_sale_price"`
	NonAcSalePrice                float64         `json:"non_ac_sale_price"`
	OnlineDeliverySellPrice       float64         `json:"online_delivery_sell_price"`
	OnlineSellPrice               float64         `json:"online_sell_price"`
	Tax                           string          `json:"tax"`
	PriceWithTax                  string          `json:"price_with_tax"`
	Cess                          int64           `json:"cess"`
	HsnCode                       string          `json:"hsn_code"`
	Description                   string          `json:"description"`
	LowStockAlert                 int64           `json:"low_stock_alert"`
	StorageLocation               string          `json:"storage_location"`
	BulkPurchaseUnit              string          `json:"bulk_purchase_unit"`
	RetailSaleUnitPerBulkPurchase float64         `json:"retail_sale_unit_per_bulk_purchase"`
	BulkPurchaseUnitPerRetailSale float64         `json:"bulk_purchase_unit_per_retail_sale"`
	ExpiryDate                    string          `json:"expiry_date"`
	ShowProductOnlineStore        string          `json:"show_product_online_store"`
	CreatedAt                     time.Time       `json:"created_at"`
	UpdatedAt                     time.Time       `json:"updated_at"`
}

type ProductResponse struct {
	OrganizationID                uint            `json:"organization_id"`
	ID                            uint            `json:"id" gorm:"primaryKey"`
	Name                          string          `json:"name"`
	Image                         string          `json:"image"`
	SellPrice                     float64         `json:"sell_price"`
	MeasuringUnit                 string          `json:"measuring_unit"`
	CategoryID                    uint            `json:"category_id"`
	Quantity                      int64           `json:"quantity"`
	Mrp                           float64         `json:"mrp"`
	PurchasePrice                 float64         `json:"purchase_price"`
	AcSalePrice                   float64         `json:"ac_sale_price"`
	NonAcSalePrice                float64         `json:"non_ac_sale_price"`
	OnlineDeliverySellPrice       float64         `json:"online_delivery_sell_price"`
	OnlineSellPrice               float64         `json:"online_sell_price"`
	Tax                           string         `json:"tax"`
	PriceWithTax                  string          `json:"price_with_tax"`
	Cess                          int64           `json:"cess"`
	HsnCode                       string          `json:"hsn_code"`
	Description                   string          `json:"description"`
	LowStockAlert                 int64           `json:"low_stock_alert"`
	StorageLocation               string          `json:"storage_location"`
	BulkPurchaseUnit              string          `json:"bulk_purchase_unit"`
	RetailSaleUnitPerBulkPurchase float64         `json:"retail_sale_unit_per_bulk_purchase"`
	BulkPurchaseUnitPerRetailSale float64         `json:"bulk_purchase_unit_per_retail_sale"`
	ExpiryDate                    string          `json:"expiry_date"`
	ShowProductOnlineStore        string          `json:"show_product_online_store"`
}
