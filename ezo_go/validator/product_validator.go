package validator

import (
	"github.com/NazishAhsan/easy_busy_book_go/model"
	"github.com/go-playground/validator"
)

type ProductValidator interface {
	ProductValidate(product model.Product) error
}

type (
	Product struct {
		Name                          string  `json:"name" validate:"required"`
		Image                         string  `json:"image" validate:"required"`
		SellPrice                     float64 `json:"sell_price" validate:"required"`
		MeasuringUnit                 string  `json:"measuring_unit" validate:"required"`
		CategoryID                    uint    `json:"category_id" validate:"required"`
		Quantity                      int64   `json:"quantity" validate:"required"`
		Mrp                           float64 `json:"mrp" validate:"required"`
		PurchasePrice                 float64 `json:"purchase_price" validate:"required"`
		AcSellPrice                   float64 `json:"ac_sell_price" validate:"required"`
		NonAcSellPrice                float64 `json:"non_ac_sell_price" validate:"required"`
		OnlineDeliverySellPrice       float64 `json:"online_delivery_sell_price" validate:"required"`
		OnlineSellPrice               float64 `json:"online_sell_price" validate:"required"`
		Tax                           string `json:"tax" validate:"omitempty"`
		PriceWithTax                  string  `json:"price_with_tax" validate:"omitempty"`
		Cess                          int64   `json:"cess" validate:"omitempty"`
		HsnCode                       string  `json:"hsn_code" validate:"omitempty"`
		Description                   string  `json:"description" validate:"omitempty"`
		LowStockAlert                 int64   `json:"low_stock_alert" validate:"omitempty"`
		StorageLocation               string  `json:"storage_location" validate:"omitempty"`
		BulkPurchaseUnit              string  `json:"bulk_purchase_unit" validate:"omitempty"`
		RetailSaleUnitPerBulkPurchase float64 `json:"retail_sale_unit_per_bulk_purchase" validate:"omitempty"`
		BulkPurchaseUnitPerRetailSale float64 `json:"bulk_purchase_unit_per_retail_sale" validate:"omitempty"`
		ExpiryDate                    string  `json:"expiry_date" validate:"omitempty"`
		ShowProductOnlineStore        string  `json:"show_product_online_store" validate:"omitempty"`
	}

	productValidator struct {
		validator *validator.Validate
	}
)

func NewProductValidator() ProductValidator {
	return &productValidator{
		validator: validator.New(),
	}
}

func (pr *productValidator) ProductValidate(product model.Product) error {

	if err := pr.validator.Struct(&product); err != nil {
		return err
	}
	return nil
}
