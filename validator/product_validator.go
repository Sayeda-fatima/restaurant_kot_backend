package validator

import (
	"github.com/Sayeda-fatima/restaurant_kot_backend/model"
	"github.com/go-playground/validator"
)

type (
	ProductValidator interface {
		ProductValidate(product *model.Product) error
	}

	productValidator struct{
		validator *validator.Validate
	}
)

func NewProductValidator() ProductValidator{
	return &productValidator{
		validator: validator.New(),
	}
}

func (pr *productValidator) ProductValidate(product *model.Product) error{

	if err := pr.validator.Struct(product); err != nil{
		return err
	}
	return nil
}