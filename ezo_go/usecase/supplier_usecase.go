package usecase

import (
	"github.com/NazishAhsan/easy_busy_book_go/common"
	"github.com/NazishAhsan/easy_busy_book_go/model"
	"github.com/NazishAhsan/easy_busy_book_go/repository"
	"github.com/NazishAhsan/easy_busy_book_go/validator"
)

type (
	SupplierUsecase interface {
		GetSupplierList(organizationID uint) ([]model.SupplierResponse, error)
		CreateSupplier(supplier model.Supplier) (model.SupplierResponse, error)
		UpdateSupplier(supplier model.Supplier, id uint) (model.SupplierResponse, error)
		DeleteSupplier(supplier model.Supplier, id uint) error
		SearchSupplier(organizationID uint, term string) ([]model.SupplierResponse, error)
	}

	supplierUsecase struct {
		sr repository.SupplierRepository
		sv validator.SupplierValidator
	}
)

func NewSupplierUsecase(sr repository.SupplierRepository, sv validator.SupplierValidator) SupplierUsecase {
	return &supplierUsecase{sr, sv}
}

func (su *supplierUsecase) GetSupplierList(organizationID uint) ([]model.SupplierResponse, error) {

	suppliers := []model.Supplier{}
	if err := su.sr.GetSupplierList(&suppliers, organizationID); err != nil {
		common.Logger.LogError().Fields(map[string]interface{}{"error": err.Error()}).Msg("GetSupplierList")
		return nil, err
	}

	resSuppliers := []model.SupplierResponse{}
	for _, v := range suppliers {
		res := model.SupplierResponse{
			ID:                 v.ID,
			Name:               v.Name,
			PhoneNo:            v.PhoneNo,
			Category:           v.Category,
			BillingAddress:     v.BillingAddress,
			BillingProvince:    v.BillingProvince,
			BillingPostalCode:  v.BillingPostalCode,
			DeliveryAddress:    v.DeliveryAddress,
			DeliveryProvince:   v.DeliveryProvince,
			DeliveryPostalCode: v.DeliveryPostalCode,
			GstNumber:          v.GstNumber,
			BillingTerm:        v.BillingTerm,
			BillingType:        v.BillingType,
			DateOfBirth:        v.DateOfBirth,
			WhatsappAlert:      v.WhatsappAlert,
		}
		resSuppliers = append(resSuppliers, res)
	}

	return resSuppliers, nil
}

func (su *supplierUsecase) CreateSupplier(supplier model.Supplier) (model.SupplierResponse, error) {

	if err := su.sv.SupplierValidate(supplier); err != nil {
		common.Logger.LogError().Fields(map[string]interface{}{"error": err.Error()}).Msg("CreateSupplier")
		return model.SupplierResponse{}, err
	}

	if err := su.sr.CreateSupplier(&supplier); err != nil {
		return model.SupplierResponse{}, err
	}

	resSupplier := model.SupplierResponse{
		ID:                 supplier.ID,
		Name:               supplier.Name,
		PhoneNo:            supplier.PhoneNo,
		Category:           supplier.Category,
		BillingAddress:     supplier.BillingAddress,
		BillingProvince:    supplier.BillingProvince,
		BillingPostalCode:  supplier.BillingPostalCode,
		DeliveryAddress:    supplier.DeliveryAddress,
		DeliveryProvince:   supplier.DeliveryProvince,
		DeliveryPostalCode: supplier.DeliveryPostalCode,
		GstNumber:          supplier.GstNumber,
		BillingTerm:        supplier.BillingTerm,
		BillingType:        supplier.BillingType,
		DateOfBirth:        supplier.DateOfBirth,
		WhatsappAlert:      supplier.WhatsappAlert,
	}

	return resSupplier, nil
}

func (su *supplierUsecase) UpdateSupplier(supplier model.Supplier, id uint) (model.SupplierResponse, error) {

	if err := su.sv.SupplierValidate(supplier); err != nil {
		common.Logger.LogError().Fields(map[string]interface{}{"error": err.Error()}).Msg("UpdateSupplier")
		return model.SupplierResponse{}, err
	}

	if err := su.sr.UpdateSupplier(&supplier, id); err != nil {
		common.Logger.LogError().Fields(map[string]interface{}{"error": err.Error()}).Msg("UpdateSupplier")
		return model.SupplierResponse{}, err
	}

	resSupplier := model.SupplierResponse{
		ID:                 supplier.ID,
		Name:               supplier.Name,
		PhoneNo:            supplier.PhoneNo,
		Category:           supplier.Category,
		BillingAddress:     supplier.BillingAddress,
		BillingProvince:    supplier.BillingProvince,
		BillingPostalCode:  supplier.BillingPostalCode,
		DeliveryAddress:    supplier.DeliveryAddress,
		DeliveryProvince:   supplier.DeliveryProvince,
		DeliveryPostalCode: supplier.DeliveryPostalCode,
		GstNumber:          supplier.GstNumber,
		BillingTerm:        supplier.BillingTerm,
		BillingType:        supplier.BillingType,
		DateOfBirth:        supplier.DateOfBirth,
		WhatsappAlert:      supplier.WhatsappAlert,
	}

	return resSupplier, nil
}

func (su *supplierUsecase) DeleteSupplier(supplier model.Supplier, id uint) error {

	if err := su.sr.DeleteSupplier(&supplier, id); err != nil {
		common.Logger.LogError().Fields(map[string]interface{}{"error": err.Error()}).Msg("DeleteSupplier")
		return err
	}
	return nil
}

func (su *supplierUsecase) SearchSupplier(organizationID uint, term string) ([]model.SupplierResponse, error){

	suppliers := []model.Supplier{}
	if err := su.sr.SearchSupplier(&suppliers, organizationID, term); err!=nil{
		return nil, err
	}

	resSupplier := []model.SupplierResponse{}
	for _, v := range(suppliers){
		res := model.SupplierResponse{
			ID:                 v.ID,
			Name:               v.Name,
			PhoneNo:            v.PhoneNo,
			Category:           v.Category,
			BillingAddress:     v.BillingAddress,
			BillingProvince:    v.BillingProvince,
			BillingPostalCode:  v.BillingPostalCode,
			DeliveryAddress:    v.DeliveryAddress,
			DeliveryProvince:   v.DeliveryProvince,
			DeliveryPostalCode: v.DeliveryPostalCode,
			GstNumber:          v.GstNumber,
			BillingTerm:        v.BillingTerm,
			BillingType:        v.BillingType,
			DateOfBirth:        v.DateOfBirth,
			WhatsappAlert:      v.WhatsappAlert,
		}
		resSupplier = append(resSupplier, res)
	}

	return resSupplier, nil
}