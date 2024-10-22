package repository

import (
	"github.com/NazishAhsan/easy_busy_book_go/model"
	"gorm.io/gorm"
)

type(
	SupplierRepository interface{
		GetSupplierList (suppliers *[]model.Supplier, organizationID uint) error
		CreateSupplier (supplier *model.Supplier) error
		UpdateSupplier (supplier *model.Supplier, id uint) error
		DeleteSupplier (supplier *model.Supplier, id uint) error
	}

	supplierRepository struct{
		db *gorm.DB
	}
)

func NewSupplierRepository(db *gorm.DB) SupplierRepository{
	return &supplierRepository{db}
}

func (sr *supplierRepository) GetSupplierList (suppliers *[]model.Supplier, organizationID uint) error{

	if err := sr.db.Where("organization_id=? and is_deleted=0", organizationID).Find(suppliers).Error; err!=nil{
		return err
	}

	return nil
}

func (sr *supplierRepository) CreateSupplier(supplier *model.Supplier) error{

	if err := sr.db.Create(supplier).Error; err!=nil{
		return err
	}
	return nil
}

func (sr *supplierRepository) UpdateSupplier(supplier *model.Supplier, id uint) error{

	result := sr.db.Model(supplier).Where("id=?", id).Updates(supplier)

	if err := result.Error; err!=nil{
		return err
	}

	return nil
}

func (sr *supplierRepository) DeleteSupplier(supplier *model.Supplier, id uint) error{

	result := sr.db.Model(supplier).Where("id=?", id).Update("is_deleted", 1)

	if err := result.Error; err!=nil{
		return err
	}
	return nil
}