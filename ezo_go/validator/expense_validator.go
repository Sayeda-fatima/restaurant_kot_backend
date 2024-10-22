package validator

import (
	"github.com/NazishAhsan/easy_busy_book_go/model"
	"github.com/go-playground/validator"
)

type (
	ExpenseValidator interface {
		ExpenseValidate(expense model.Expense) error
	}

	expenseValidator struct {
		validator *validator.Validate
	}
)

func NewExpenseValidator() ExpenseValidator {
	return &expenseValidator{
		validator: validator.New(),
	}
}

func (er *expenseValidator) ExpenseValidate(expense model.Expense) error {

	if err := er.validator.Struct(&expense); err != nil {
		return err
	}
	return nil
}
