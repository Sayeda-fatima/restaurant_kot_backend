package validator

import (
	"github.com/NazishAhsan/easy_busy_book_go/model"
	"github.com/go-playground/validator"
)

type (
	InvoiceDetailValidator interface {
		InvoiceDetailValidate(invoiceDetail model.InvoiceDetail) error
	}

	invoiceDetailValidator struct {
		validator *validator.Validate
	}
)

func NewInvoiceDetailValidator() InvoiceDetailValidator {
	return &invoiceDetailValidator{
		validator: validator.New(),
	}
}

func (ir *invoiceDetailValidator) InvoiceDetailValidate(invoiceDetail model.InvoiceDetail) error {

	if err := ir.validator.Struct(&invoiceDetail); err != nil {
		return err
	}
	return nil
}
