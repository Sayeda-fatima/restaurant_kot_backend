package repository

import (
	"github.com/NazishAhsan/easy_busy_book_go/model"
	"gorm.io/gorm"
)

type (
	InvoiceRepository interface {
		GetInvoiceList(invoices *[]model.Invoice, organizationID uint) error
		CreateInvoice (invoice *model.Invoice) error
		UpdateInvoice (invoice *model.Invoice, id uint) error
		DeleteInvoice (invoice *model.Invoice, id uint) error
	}

	invoiceRepository struct{
		db *gorm.DB
	}
)

func NewInvoiceRepository (db *gorm.DB) InvoiceRepository{
	return &invoiceRepository{db}
}

func (ir *invoiceRepository) GetInvoiceList(invoices *[]model.Invoice, organizationID uint) error{

	if err := ir.db.Where("organization_id=? and is_deleted=0", organizationID).Find(invoices).Error; err!=nil{
		return err
	}
	return nil
}

func (ir *invoiceRepository) CreateInvoice(invoice *model.Invoice) error{

	if err := ir.db.Create(invoice).Error; err!=nil{
		return err
	}
	return nil
}

func (ir *invoiceRepository) UpdateInvoice(invoice *model.Invoice, id uint) error{

	result := ir.db.Model(invoice).Where("id=?", id).Updates(invoice)

	if err := result.Error; err!=nil{
		return err
	}
	return nil
}

func (ir *invoiceRepository) DeleteInvoice(invoice *model.Invoice, id uint) error{

	result := ir.db.Model(invoice).Where("id=?", id).Update("is_deleted", 1)

	if err := result.Error; err!=nil{
		return err
	}
	return nil
}