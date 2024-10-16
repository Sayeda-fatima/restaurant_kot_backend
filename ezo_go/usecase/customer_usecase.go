package usecase

import (
	"github.com/NazishAhsan/easy_busy_book_go/model"
	"github.com/NazishAhsan/easy_busy_book_go/repository"
	"github.com/NazishAhsan/easy_busy_book_go/validator"
)

type( 
	CustomerUsecase interface{
		GetCustomerList (organizationID uint) ([]model.CustomerResponse, error)
		CreateCustomer(customer model.Customer) (model.CustomerResponse, error)
		UpdateCustomer (customer model.Customer, id uint) (model.CustomerResponse, error)
		DeleteCustomer (customer model.Customer, id uint) (model.CustomerResponse, error)
	}

	customerUsecase struct{
		cr repository.CustomerRepository
		cv validator.CustomerValidator
	}
)

func NewCustomerUsecase(cr repository.CustomerRepository, cv validator.CustomerValidator) CustomerUsecase{
	return &customerUsecase{cr,cv}
}

func (cu *customerUsecase) GetCustomerList (organizationID uint) ([]model.CustomerResponse, error){

	customers := []model.Customer{}
	 
	if err := cu.cr.GetCustomerList(&customers, organizationID).Error; err!=nil{
		return []model.CustomerResponse{}, err
	}

	for _, v range customers{
		res := model.CustomerResponse{
			CustomerName: v.CustomerName,
			CustomerPhoneNo: v.CustomerPhoneNo,
		}

		resCustomer := append(resCustomer, res)
	}

	return resCustomer, nil
}

func (cu *customerUsecase) CreateCustomer (customer model.Customer) (model.CustomerResponse, error) {

	if err := cu.cv.CustomerValidate(customer); err!=nil{
		return model.CustomerResponse{}, err
	}
}

func (cu *customerUsecase) UpdateCustomer (customer model.Customer, id uint) (model.CustomerResponse, error) {

	
}

func (cu *customerUsecase) DeleteCustomer (customer model.Customer, id uint) (model.CustomerResponse, error){

}