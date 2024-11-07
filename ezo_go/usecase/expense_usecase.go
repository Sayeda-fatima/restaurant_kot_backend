package usecase

import (
	"github.com/NazishAhsan/easy_busy_book_go/common"
	"github.com/NazishAhsan/easy_busy_book_go/model"
	"github.com/NazishAhsan/easy_busy_book_go/repository"
	"github.com/NazishAhsan/easy_busy_book_go/validator"
)

type (
	ExpenseUsecase interface {
		GetExpenseList(organizationID uint) ([]model.ExpenseResponse, error)
		CreateExpense(expense model.Expense) (model.ExpenseResponse, error)
		UpdateExpense(expense model.Expense, id uint) (model.ExpenseResponse, error)
		DeleteExpense(expense model.Expense, id uint) error
		ExpenseReport(organizationID uint, dateFrom string, dateTo string)([]model.ExpenseResponse, error)
	}

	expenseUsecase struct {
		er repository.ExpenseRepository
		ev validator.ExpenseValidator
	}
)

func NewExpenseUsecase(er repository.ExpenseRepository, ev validator.ExpenseValidator) ExpenseUsecase {
	return &expenseUsecase{er, ev}
}

func (eu *expenseUsecase) GetExpenseList(organizationID uint) ([]model.ExpenseResponse, error) {

	expenses := []model.Expense{}
	if err := eu.er.GetExpenseList(&expenses, organizationID); err != nil {
		common.Logger.LogError().Fields(map[string]interface{}{"error": err.Error()}).Msg("GetExpenseList")
		return nil, err
	}

	resExpenses := []model.ExpenseResponse{}
	for _, v := range expenses {
		res := model.ExpenseResponse{
			OrganizationID:  v.OrganizationID,
			ID:              v.ID,
			SupplierID:      v.SupplierID,
			SupplierName:    v.SupplierName,
			ExpenseCategory: v.ExpenseCategory,
			TotalAmount:     v.TotalAmount,
			AmountPaid:      v.AmountPaid,
			AmountDue:       v.AmountDue,
			Note:            v.Note,
			ModeOfPayment:   v.ModeOfPayment,
		}
		resExpenses = append(resExpenses, res)
	}

	return resExpenses, nil
}

func (eu *expenseUsecase) CreateExpense(expense model.Expense) (model.ExpenseResponse, error) {

	if err := eu.ev.ExpenseValidate(expense); err != nil {
		common.Logger.LogError().Fields(map[string]interface{}{"error": err.Error()}).Msg("CreateExpense")
		return model.ExpenseResponse{}, err
	}

	if err := eu.er.CreateExpense(&expense); err != nil {
		common.Logger.LogError().Fields(map[string]interface{}{"error": err.Error()}).Msg("CreateExpense")
		return model.ExpenseResponse{}, err
	}

	resExpense := model.ExpenseResponse{
		ID:              expense.ID,
		OrganizationID:  expense.OrganizationID,
		SupplierID:      expense.SupplierID,
		SupplierName:    expense.SupplierName,
		ExpenseCategory: expense.ExpenseCategory,
		TotalAmount:     expense.TotalAmount,
		AmountPaid:      expense.AmountPaid,
		AmountDue:       expense.AmountDue,
		Note:            expense.Note,
		ModeOfPayment:   expense.ModeOfPayment,
	}
	return resExpense, nil
}

func (eu *expenseUsecase) UpdateExpense(expense model.Expense, id uint) (model.ExpenseResponse, error) {

	if err := eu.ev.ExpenseValidate(expense); err != nil {
		common.Logger.LogError().Fields(map[string]interface{}{"error": err.Error()}).Msg("UpdateExpense")
		return model.ExpenseResponse{}, err
	}

	if err := eu.er.UpdateExpense(&expense, id); err != nil {
		common.Logger.LogError().Fields(map[string]interface{}{"error": err.Error()}).Msg("UpdateExpense")
		return model.ExpenseResponse{}, err
	}

	resExpense := model.ExpenseResponse{
		ID:              expense.ID,
		OrganizationID:  expense.OrganizationID,
		SupplierID:      expense.SupplierID,
		SupplierName:    expense.SupplierName,
		ExpenseCategory: expense.ExpenseCategory,
		TotalAmount:     expense.TotalAmount,
		AmountPaid:      expense.AmountPaid,
		AmountDue:       expense.AmountDue,
		Note:            expense.Note,
		ModeOfPayment:   expense.ModeOfPayment,
	}
	return resExpense, nil
}

func (eu *expenseUsecase) DeleteExpense(expense model.Expense, id uint) error {

	if err := eu.er.DeleteExpense(&expense, id); err != nil {
		common.Logger.LogError().Fields(map[string]interface{}{"error": err.Error()}).Msg("DeleteExpense")
		return err
	}
	return nil
}

func (eu *expenseUsecase) ExpenseReport(organizationID uint, dateFrom string, dateTo string)([]model.ExpenseResponse, error){

	expenses := []model.Expense{}
	if err := eu.er.ExpenseReport(&expenses, organizationID, dateFrom, dateTo); err!=nil{
		return nil, err
	}

	resExpense := []model.ExpenseResponse{}
	for _, v := range(expenses){
		res := model.ExpenseResponse{
			ID: v.ID,
			OrganizationID: v.OrganizationID,
			SupplierID: v.SupplierID,
			SupplierName: v.SupplierName,
			ExpenseCategory: v.ExpenseCategory,
			TotalAmount: v.TotalAmount,
			AmountPaid: v.AmountPaid,
			AmountDue: v.AmountDue,
			Note: v.Note,
			ModeOfPayment: v.ModeOfPayment,
		}
		resExpense = append(resExpense, res)
	}
	return resExpense, nil
}