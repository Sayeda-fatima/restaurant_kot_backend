package repository

import (
	"github.com/NazishAhsan/easy_busy_book_go/model"
	"gorm.io/gorm"
)

type CustomerRepository interface{
	GetCustomerList (customers *[]model.Customer, organizationID uint) error
	CreateCustomer (customer *model.Customer) error
	UpdateCustomer (customer *model.Customer, id uint, customerUpdate *model.CustomerResponse) error
	DeleteCustomer (customer *model.Customer, id uint) error
}

type customerRepository struct{
	db *gorm.DB
}

func NewCustomerRepository (db *gorm.DB) CustomerRepository {

	return &customerRepository{db}
}

func (cr *customerRepository) GetCustomerList(customers *[]model.Customer, organizationID uint) error {

	if err := cr.db.Where("organization_id=? and is_deleted=0", organizationID).Find(customers).Error; err!=nil{
		return err
	}
	return nil
}

func (cr *customerRepository) CreateCustomer(customer *model.Customer) error {

	if err := cr.db.Create(customer).Error; err!=nil{
		return err
	}
	return nil
}

func (cr *customerRepository) UpdateCustomer (customer *model.Customer, id uint, customerUpdate *model.CustomerResponse) error {

	result := cr.db.Model(customer).Where("id=?", id).Updates(customerUpdate)
	if err := result.Error; err!=nil{
		return err
	}
	return nil
}

func (cr *customerRepository) DeleteCustomer (customer *model.Customer, id uint) error {

	result := cr.db.Model(customer).Where("id=?", id).Update("is_deleted", 1)

	if err := result.Error; err!=nil{
		return err
	}
	return nil
}