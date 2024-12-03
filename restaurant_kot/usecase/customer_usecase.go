package usecase

import (
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/model"
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/repository"
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/validator"
)

type (
	CustomerUsecase interface {
		GetCustomerList(organizationID uint, restaurantID uint) ([]model.CustomerResponse, error)
		CreateCustomer(customer model.Customer) (model.CustomerResponse, error)
		UpdateCustomer(customer model.Customer, id uint) (model.CustomerResponse, error)
		DeleteCustomer(customer model.Customer, id uint) error
	}
	customerUsecase struct{
		cr repository.CustomerRepository
		cv validator.CustomerValidator
	}
)

func NewCustomerUsecase(cr repository.CustomerRepository, cv validator.CustomerValidator) CustomerUsecase{
	return &customerUsecase{cr,cv}
}

func (cu *customerUsecase) GetCustomerList(organizationID uint, restaurantID uint) ([]model.CustomerResponse, error){

	customers := []model.Customer{}
	if err := cu.cr.GetCustomerList(&customers, organizationID, restaurantID); err != nil{
		return nil, err
	}

	resCustomer := []model.CustomerResponse{}
	for _, v := range(customers){
		res := model.CustomerResponse{
			ID: v.ID,
			OrganizationID: v.OrganizationID,
			RestaurantID: v.RestaurantID,
			Name: v.Name,
			PhoneNo: v.PhoneNo,
			Email: v.Email,
		}
		resCustomer = append(resCustomer, res)
	}
	return resCustomer, nil
}

func (cu *customerUsecase) CreateCustomer(customer model.Customer) (model.CustomerResponse, error){

	if err := cu.cv.CustomerValidate(&customer); err != nil{
		return model.CustomerResponse{}, err
	}

	if err := cu.cr.CreateCustomer(&customer); err != nil{
		return model.CustomerResponse{}, err
	}

	resCustomer := model.CustomerResponse{
		ID: customer.ID,
		OrganizationID: customer.OrganizationID,
		RestaurantID: customer.RestaurantID,
		Name: customer.Name,
		PhoneNo: customer.PhoneNo,
		Email: customer.Email,
	}
	return resCustomer, nil
}

func (cu *customerUsecase) UpdateCustomer(customer model.Customer, id uint) (model.CustomerResponse, error){

	if err := cu.cv.CustomerValidate(&customer); err != nil{
		return model.CustomerResponse{}, err
	}

	if err := cu.cr.UpdateCustomer(&customer, id); err != nil{
		return model.CustomerResponse{}, err
	}

	resCustomer := model.CustomerResponse{
		ID: customer.ID,
		OrganizationID: customer.OrganizationID,
		RestaurantID: customer.RestaurantID,
		Name: customer.Name,
		Email: customer.Email,
		PhoneNo: customer.PhoneNo,
	}
	return resCustomer, nil
}

func (cu *customerUsecase) DeleteCustomer(customer model.Customer, id uint) error{

	if err := cu.cr.DeleteCustomer(&customer, id); err != nil{
		return err
	}
	return nil
}