package validator

import (
	"github.com/NazishAhsan/easy_busy_book_go/model"
	"github.com/go-playground/validator"
)

type (
	InvoiceValidator interface {
		InvoiceValidate(invoice model.Invoice) error
	}

	invoiceValidator struct {
		validator *validator.Validate
	}
)

func NewInvoiceValidator() InvoiceValidator {
	return &invoiceValidator{
		validator: validator.New(),
	}
}

func (ir *invoiceValidator) InvoiceValidate(invoice model.Invoice) error {

	if err := ir.validator.Struct(&invoice); err != nil {
		return err
	}
	return nil
}
