package usecase

import (
	"github.com/NazishAhsan/easy_busy_book_go/common"
	"github.com/NazishAhsan/easy_busy_book_go/model"
	"github.com/NazishAhsan/easy_busy_book_go/repository"
	"github.com/NazishAhsan/easy_busy_book_go/validator"
)

type( 
	CustomerUsecase interface{
		GetCustomerList (organizationID uint) ([]model.CustomerResponse, error)
		CreateCustomer(customer model.Customer) (model.CustomerResponse, error)
		UpdateCustomer (customer model.Customer, id uint) (model.CustomerResponse, error)
		DeleteCustomer (customer model.Customer, id uint) error
		SearchCustomer (organizationID uint, term string) ([]model.CustomerResponse, error)
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
	 
	if err := cu.cr.GetCustomerList(&customers, organizationID); err!=nil{
		common.Logger.LogError().Msg(err.Error())
		return nil, err
	}

	resCustomers := []model.CustomerResponse{}
	for _, v:= range customers{
		res := model.CustomerResponse{
			ID: v.ID,
			Name: v.Name,
			PhoneNo: v.PhoneNo,
			Category: v.Category,
			BillingAddress: v.BillingAddress,
			BillingProvince: v.BillingProvince,
			BillingPostalCode: v.BillingPostalCode,
			DeliveryAddress: v.DeliveryAddress,
			DeliveryProvince: v.DeliveryProvince,
			DeliveryPostalCode: v.DeliveryPostalCode,
			GstNumber: v.GstNumber,
			BillingTerm: v.BillingTerm,
			BillingType: v.BillingType,
			DateOfBirth: v.DateOfBirth,
			WhatsappAlert: v.WhatsappAlert,
		}
		resCustomers = append(resCustomers, res);
	}

	return resCustomers, nil
}

func (cu *customerUsecase) CreateCustomer (customer model.Customer) (model.CustomerResponse, error) {

	if err := cu.cv.CustomerValidate(customer); err!=nil{
		common.Logger.LogError().Msg(err.Error())
		return model.CustomerResponse{}, err
	}

	if err := cu.cr.CreateCustomer(&customer); err!=nil {
		common.Logger.LogError().Msg(err.Error())
		return model.CustomerResponse{}, err
	}

	resCustomer := model.CustomerResponse{
		ID: customer.ID,
		Name: customer.Name,
		PhoneNo: customer.PhoneNo,
		Category: customer.Category,
		BillingAddress: customer.BillingAddress,
		BillingProvince: customer.BillingProvince,
		BillingPostalCode: customer.BillingPostalCode,
		DeliveryAddress: customer.DeliveryAddress,
		DeliveryProvince: customer.DeliveryProvince,
		DeliveryPostalCode: customer.DeliveryPostalCode,
		GstNumber: customer.GstNumber,
		BillingTerm: customer.BillingTerm,
		BillingType: customer.BillingType,
		DateOfBirth: customer.DateOfBirth,
		WhatsappAlert: customer.WhatsappAlert,
	}
	return resCustomer, nil
}

func (cu *customerUsecase) UpdateCustomer (customer model.Customer, id uint) (model.CustomerResponse, error) {
	
	if err := cu.cv.CustomerValidate(customer); err!=nil{
		common.Logger.LogError().Msg(err.Error())
		return model.CustomerResponse{}, err
	}

	if err := cu.cr.UpdateCustomer(&customer, id); err!=nil{
		common.Logger.LogError().Msg(err.Error())
		return model.CustomerResponse{}, err
	}

	resCustomer := model.CustomerResponse{
		ID: customer.ID,
		Name: customer.Name,
		PhoneNo: customer.PhoneNo,
		Category: customer.Category,
		BillingAddress: customer.BillingAddress,
		BillingProvince: customer.BillingProvince,
		BillingPostalCode: customer.BillingPostalCode,
		DeliveryAddress: customer.DeliveryAddress,
		DeliveryProvince: customer.DeliveryProvince,
		DeliveryPostalCode: customer.DeliveryPostalCode,
		GstNumber: customer.GstNumber,
		BillingTerm: customer.BillingTerm,
		BillingType: customer.BillingType,
		DateOfBirth: customer.DateOfBirth,
		WhatsappAlert: customer.WhatsappAlert,
	}

	return resCustomer, nil
}

func (cu *customerUsecase) DeleteCustomer (customer model.Customer, id uint) error {

	if err := cu.cr.DeleteCustomer(&customer, id); err!=nil{
		common.Logger.LogError().Msg(err.Error())
		return err
	}
	return nil

}

func (cu *customerUsecase) SearchCustomer(organizationID uint, term string) ([]model.CustomerResponse, error){

	customers := []model.Customer{}

	if err := cu.cr.SearchCustomer(&customers, organizationID, term); err!=nil{
		return nil, err
	}
	resCustomers := []model.CustomerResponse{}
	for _, v:= range customers{
		res := model.CustomerResponse{
			ID: v.ID,
			Name: v.Name,
			PhoneNo: v.PhoneNo,
			Category: v.Category,
			BillingAddress: v.BillingAddress,
			BillingProvince: v.BillingProvince,
			BillingPostalCode: v.BillingPostalCode,
			DeliveryAddress: v.DeliveryAddress,
			DeliveryProvince: v.DeliveryProvince,
			DeliveryPostalCode: v.DeliveryPostalCode,
			GstNumber: v.GstNumber,
			BillingTerm: v.BillingTerm,
			BillingType: v.BillingType,
			DateOfBirth: v.DateOfBirth,
			WhatsappAlert: v.WhatsappAlert,
		}
		resCustomers = append(resCustomers, res);
	}

	return resCustomers, nil
}