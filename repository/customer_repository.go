package repository

import (
	"fmt"

	"github.com/Sayeda-fatima/restaurant_kot_backend/model"
	"gorm.io/gorm"
)

type (
	CustomerRepository interface {
		GetCustomerList(customers *[]model.Customer, organizationID uint, restaurantID uint) error
		CreateCustomer(customer *model.Customer) error
		UpdateCustomer(customer *model.Customer, id uint, organizationID uint, restaurantID uint) error
		DeleteCustomer(customer *model.Customer, id uint, organizationID uint, restaurantID uint) error
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

func (cr *customerRepository) UpdateCustomer(customer *model.Customer, id uint, organizationID uint, restaurantID uint) error{

	result := cr.db.Model(customer).Where("id=? and organization_id=? and restaurant_id=?", id, organizationID, restaurantID).Updates(customer)

	if err := result.Error; err != nil{
		return err
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	
	return nil
}

func (cr *customerRepository) DeleteCustomer(customer *model.Customer, id uint, organizationID uint, restaurantID uint) error{

	result := cr.db.Model(customer).Where("id=? and organization_id=? and restaurant_id=?", id, organizationID, restaurantID).Update("is_deleted",1)

	if err := result.Error; err != nil{
		return err
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	
	return nil
}