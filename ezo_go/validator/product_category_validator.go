package validator

import (
	"github.com/NazishAhsan/easy_busy_book_go/model"
	"github.com/go-playground/validator"
)

type ProductCategoryValidator interface {
	ProductCategoryValidate(productCategory model.ProductCategory) error
}

type productCategoryValidator struct {
		validator *validator.Validate
	}

func NewProductCategoryValidator() ProductCategoryValidator {
	return &productCategoryValidator{
		validator: validator.New(),
	}
}

func (pr *productCategoryValidator) ProductCategoryValidate(productCategory model.ProductCategory) error {

	if err := pr.validator.Struct(&productCategory); err != nil {
		return err
	}
	return nil
}
