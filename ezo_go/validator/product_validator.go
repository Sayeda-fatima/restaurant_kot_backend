package validator

import (
	"github.com/NazishAhsan/easy_busy_book_go/model"
	"github.com/go-playground/validator"
)

type ProductValidator interface {
	ProductValidate(product model.Product) error
}

type productValidator struct {
	validator *validator.Validate
}

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
