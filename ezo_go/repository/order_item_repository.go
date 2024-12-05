package repository

import (
	"fmt"

	"github.com/NazishAhsan/easy_busy_book_go/model"
	"gorm.io/gorm"
)

type (
	OrderItemRepository interface {
		GetOrderItemList(orderItems *[]model.OrderItem, organizationID uint, orderID uint) error
		CreateOrderItem(orderItem *model.OrderItem) error
		UpdateOrderItem(orderItem *model.OrderItem, id uint) error
		DeleteOrderItem(orderItem *model.OrderItem, id uint) error
		InvoiceCustomer(orderItems *[]model.OrderItem, organizationID uint, orderID uint, dateFrom string, dateTo string) error
	}

	orderItemRepository struct{
		db *gorm.DB
	}
)

func NewOrderItemRepository(db *gorm.DB) OrderItemRepository {
	return &orderItemRepository{db}
}

func (or *orderItemRepository) GetOrderItemList(orderItem *[]model.OrderItem, organizationID uint, orderID uint) error {

	if err := or.db.Preload("Product", func(db *gorm.DB) *gorm.DB { return db.Select("id", "name", "image", "mrp") }).Where("organization_id=? and order_id=? and is_deleted=0", organizationID, orderID).Find(orderItem).Error; err != nil {
		return err
	}
	return nil
}

func (or *orderItemRepository) CreateOrderItem(order *model.OrderItem) error {

	if err := or.db.Create(order).Error; err != nil {
		return err
	}
	return nil
}

func (or *orderItemRepository) UpdateOrderItem(order *model.OrderItem, id uint) error {

	result := or.db.Model(order).Where("id=?", id).Updates(order)

	if err := result.Error; err != nil {
		return err
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("record does not exist")
	}
	
	return nil
}

func (or *orderItemRepository) DeleteOrderItem(order *model.OrderItem, id uint) error {

	result := or.db.Model(order).Where("id=?", id).Update("is_deleted", 1)

	if err := result.Error; err != nil {
		return err
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("record does not exist")
	}
	
	return nil
}

func (or *orderItemRepository) InvoiceCustomer(orderItems *[]model.OrderItem, organizationID uint, orderID uint, dateFrom string, dateTo string) error{

	if err := or.db.Where("organization_id=? and id=? and is_deleted=0 and date(created_at) between ? and ?", organizationID, orderID, dateFrom, dateTo).Find(orderItems).Error; err!=nil{
		return err
	}
	return nil
}