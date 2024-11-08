package validator

import (
	"github.com/NazishAhsan/easy_busy_book_go/model"
	"github.com/go-playground/validator"
)

type ProductImageValidator interface {
	ProductImageValidate(productImage model.ProductImage) error
}

type productImageValidator struct{
	validator *validator.Validate
}

func NewProductImageValidator () ProductImageValidator{
	return &productImageValidator{
		validator: validator.New(),
	}
}

func (pr *productImageValidator) ProductImageValidate(productImage model.ProductImage) error{

	if err := pr.validator.Struct(productImage); err!=nil{
		return err
	}
	return nil
}