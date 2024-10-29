package repository

import (
	"github.com/NazishAhsan/easy_busy_book_go/model"
	"gorm.io/gorm"
)

type (
	OrderItemRepository interface {
		GetOrderItemList(orderItems *[]model.OrderItem, organizationID uint, orderID uint) error
		CreateOrderItem(orderItem *model.OrderItem) error
		UpdateOrderItem(orderItem *model.OrderItem, id uint) error
		DeleteOrderItem(orderItem *model.OrderItem, id uint) error
	}

	orderItemRepository struct{
		db *gorm.DB
	}
)

func NewOrderItemRepository(db *gorm.DB) OrderItemRepository {
	return &orderItemRepository{db}
}

func (or *orderItemRepository) GetOrderItemList(orderItem *[]model.OrderItem, organizationID uint, orderID uint) error {

	if err := or.db.Where("organization_id=? and order_id=? and is_deleted=0", organizationID, orderID).Find(orderItem).Error; err != nil {
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
	return nil
}

func (or *orderItemRepository) DeleteOrderItem(order *model.OrderItem, id uint) error {

	result := or.db.Model(order).Where("id=?", id).Update("is_deleted", 1)

	if err := result.Error; err != nil {
		return err
	}
	return nil
}
