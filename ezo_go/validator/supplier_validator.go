package validator

import (
	"github.com/NazishAhsan/easy_busy_book_go/model"
	"github.com/go-playground/validator"
)

type SupplierValidator interface {
	SupplierValidate(supplier model.Supplier) error
}

type supplierValidator struct {
	validator *validator.Validate
}

func NewSupplierValidator() SupplierValidator {
	return &supplierValidator{
		validator: validator.New(),
	}
}

func (sr *supplierValidator) SupplierValidate(supplier model.Supplier) error {

	if err := sr.validator.Struct(supplier); err != nil {
		return err
	}
	return nil
}
