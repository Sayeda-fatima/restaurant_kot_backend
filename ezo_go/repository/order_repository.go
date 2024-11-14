package repository

import (
	"github.com/NazishAhsan/easy_busy_book_go/model"
	"gorm.io/gorm"
)

type (
	OrderRepository interface {
		GetOrderList(orders *[]model.Order, organizationID uint) error
		CreateOrder(order *model.Order) error
		UpdateOrder(order *model.Order, id uint) error
		DeleteOrder(order *model.Order, id uint) error
		InvoiceReportCustomer (order *[]model.Order, organizationID uint, customerID uint, dateFrom string, dateTo string) error
		GetInvoice(order *model.Order, id uint) error
	}

	orderRepository struct {
		db *gorm.DB
	}
)

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db}
}

func (or *orderRepository) GetOrderList(orders *[]model.Order, organizationID uint) error {

	if err := or.db.Where("organization_id=? and is_deleted=0", organizationID).Find(orders).Error; err != nil {
		return err
	}
	return nil
}

func (or *orderRepository) CreateOrder(order *model.Order) error {

	if err := or.db.Create(order).Error; err != nil {
		return err
	}
	return nil
}

func (or *orderRepository) UpdateOrder(order *model.Order, id uint) error {

	result := or.db.Model(order).Where("id=?", id).Updates(order)

	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func (or *orderRepository) DeleteOrder(order *model.Order, id uint) error {

	result := or.db.Model(order).Where("id=?", id).Update("is_deleted", 1)

	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func (or *orderRepository) InvoiceReportCustomer(order *[]model.Order, organizationID uint, customerID uint, dateFrom string, dateTo string) error{

	if err := or.db.Where("organization_id=? and customer_id=? and date(created_at) between ? and ?", organizationID, customerID, dateFrom, dateTo).Find(order).Error; err!=nil{
		return err
	}
	return nil
}

func (or *orderRepository) GetInvoice(order *model.Order, id uint) error{

	if err := or.db.Preload("OrderItems.Product").Where("id=?", id).First(order).Error; err!=nil{
		return err
	}
	return nil
}
