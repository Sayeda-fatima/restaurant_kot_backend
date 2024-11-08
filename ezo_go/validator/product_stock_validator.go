package validator

import (
	"github.com/NazishAhsan/easy_busy_book_go/model"
	"github.com/go-playground/validator"
)

type ProductStockValidator interface {
	ProductStockValidate(productStock model.ProductStock) error
}

type productStockValidator struct {
	validator *validator.Validate
}

func NewProductStockValidator() ProductStockValidator {
	return &productStockValidator{
		validator: validator.New(),
	}
}

func (pr *productStockValidator) ProductStockValidate(productStock model.ProductStock) error {

	if err := pr.validator.Struct(&productStock); err != nil {
		return err
	}
	return nil
}
