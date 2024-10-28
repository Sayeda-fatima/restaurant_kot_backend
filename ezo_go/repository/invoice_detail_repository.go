package repository

import (
	"github.com/NazishAhsan/easy_busy_book_go/model"
	"gorm.io/gorm"
)

type (
	InvoiceDetailRepository interface {
		GetInvoiceDetailList(invoiceDetails *[]model.InvoiceDetail, organizationID uint) error
		CreateInvoiceDetail (invoiceDetail *[]model.InvoiceDetail) error
		UpdateInvoiceDetail (invoiceDetail *model.InvoiceDetail, id uint) error
		DeleteInvoiceDetail (invoiceDetail *model.InvoiceDetail, id uint) error
	}

	invoiceDetailRepository struct{
		db *gorm.DB
	}
)

func NewInvoiceDetailRepository (db *gorm.DB) InvoiceDetailRepository{
	return &invoiceDetailRepository{db}
}

func (ir *invoiceDetailRepository) GetInvoiceDetailList(invoiceDetail *[]model.InvoiceDetail, organizationID uint) error{

	if err := ir.db.Where("organization_id=? and is_deleted=0", organizationID).Find(invoiceDetail).Error; err!=nil{
		return err
	}
	return nil
}

func (ir *invoiceDetailRepository) CreateInvoiceDetail(invoice *[]model.InvoiceDetail) error{

	if err := ir.db.Create(invoice).Error; err!=nil{
		return err
	}
	return nil
}

func (ir *invoiceDetailRepository) UpdateInvoiceDetail(invoice *model.InvoiceDetail, id uint) error{

	result := ir.db.Model(invoice).Where("id=?", id).Updates(invoice)

	if err := result.Error; err!=nil{
		return err
	}
	return nil
}

func (ir *invoiceDetailRepository) DeleteInvoiceDetail(invoice *model.InvoiceDetail, id uint) error{

	result := ir.db.Model(invoice).Where("id=?", id).Update("is_deleted", 1)

	if err := result.Error; err!=nil{
		return err
	}
	return nil
}
