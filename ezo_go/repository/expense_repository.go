package repository

import (
	"fmt"

	"github.com/NazishAhsan/easy_busy_book_go/model"
	"gorm.io/gorm"
)

type (
	ExpenseRepository interface {
		GetExpenseList(expense *[]model.Expense, organizationID uint) error
		CreateExpense(expense *model.Expense) error
		UpdateExpense(expense *model.Expense, id uint) error
		DeleteExpense(expense *model.Expense, id uint) error
		ExpenseReport(expense *[]model.Expense, organizationID uint, dateFrom string, dateTo string) error
	}

	expenseRepository struct {
		db *gorm.DB
	}
)

func NewExpenseRepository(db *gorm.DB) ExpenseRepository {
	return &expenseRepository{db}
}

func (er *expenseRepository) GetExpenseList(expense *[]model.Expense, organizationID uint) error {

	if err := er.db.Where("organization_id=? and is_deleted=0", organizationID).Find(expense).Error; err != nil {
		return err
	}

	return nil
}

func (er *expenseRepository) CreateExpense(expense *model.Expense) error {

	if err := er.db.Create(expense).Error; err != nil {
		return err
	}
	return nil
}

func (er *expenseRepository) UpdateExpense(expense *model.Expense, id uint) error {

	result := er.db.Model(expense).Where("id=?", id).Updates(expense)

	if err := result.Error; err != nil {
		return err
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("record does not exist")
	}
	
	return nil
}

func (er *expenseRepository) DeleteExpense(expense *model.Expense, id uint) error {

	result := er.db.Model(expense).Where("id=?", id).Update("is_deleted", 1)

	if err := result.Error; err != nil {
		return err
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("record does not exist")
	}
	
	return nil
}

func (er *expenseRepository) ExpenseReport(expense *[]model.Expense, organizationID uint, dateFrom string, dateTo string) error{

	if err := er.db.Where("organization_id=? and is_deleted=0 and date(created_at) between ? and ?", organizationID, dateFrom, dateTo).Find(expense).Error; err!=nil{
		return err
	}
	return nil
}