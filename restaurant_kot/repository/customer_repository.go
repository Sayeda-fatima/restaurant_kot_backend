package repository

import (
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/model"
	"gorm.io/gorm"
)

type (
	CustomerRepository interface {
		GetCustomerList(customers *[]model.Customer, organizationID uint, restaurantID uint) error
		CreateCustomer(customer *model.Customer) error
		UpdateCustomer(customer *model.Customer, id uint) error
		DeleteCustomer(customer *model.Customer, id uint) error
	}

	customerRepository struct{
		db *gorm.DB
	}
)

func NewCustomerRepository(db *gorm.DB) CustomerRepository{
	return &customerRepository{db}
}

func (cr *customerRepository) GetCustomerList(customers *[]model.Customer, organizationID uint, restaurantID uint) error{

	if err := cr.db.Where("organization_id=? and restaurant_id=? and is_deleted=0", organizationID, restaurantID).Find(customers).Error; err != nil{
		return err
	}
	return nil
}

func (cr *customerRepository) CreateCustomer(customer *model.Customer) error{

	if err := cr.db.Create(customer).Error; err != nil{
		return err
	}
	return nil
}

func (cr *customerRepository) UpdateCustomer(customer *model.Customer, id uint) error{

	result := cr.db.Model(customer).Where("id=?", id).Updates(customer)

	if err := result.Error; err != nil{
		return err
	}
	return nil
}

func (cr *customerRepository) DeleteCustomer(customer *model.Customer, id uint) error{

	result := cr.db.Model(customer).Where("id=?", id).Update("is_deleted",1)

	if err := result.Error; err != nil{
		return err
	}
	return nil
}