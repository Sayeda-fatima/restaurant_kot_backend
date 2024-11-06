package repository

import (
	"github.com/NazishAhsan/easy_busy_book_go/common"
	"github.com/NazishAhsan/easy_busy_book_go/model"
	"gorm.io/gorm"
)

type CustomerRepository interface{
	GetCustomerList (customers *[]model.Customer, organizationID uint) error
	CreateCustomer (customer *model.Customer) error
	UpdateCustomer (customer *model.Customer, id uint) error
	DeleteCustomer (customer *model.Customer, id uint) error
	SearchCustomer (customer *[]model.Customer, organizationID uint, term string) error
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

func (cr *customerRepository) UpdateCustomer (customer *model.Customer, id uint) error {

	//print("data: %s")
	common.Logger.LogInfo().Fields(map[string]interface{}{"data":customer}).Msg("test")
	result := cr.db.Model(customer).Where("id=?", id).Updates(customer)
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

func (cr *customerRepository) SearchCustomer (customers *[]model.Customer, organizationID uint, term string) error{

	if err := cr.db.Where("id LIKE ? or name LIKE ? or phone_no LIKE ?", "%"+term+"%", "%"+term+"%", "%"+term+"%").Having("organization_id=? and is_deleted=0", organizationID).Find(customers).Error; err!=nil{
		return err
	}
	return nil
}