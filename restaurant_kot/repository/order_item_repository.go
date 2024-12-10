package repository

import (
	"fmt"

	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/model"
	"gorm.io/gorm"
)

type (
	OrderItemRepository interface {
		GetOrderItemList(orderItem *[]model.OrderItem, organizationID uint, restaurantID uint, orderID uint) error
		CreateOrderItem(orderItem *model.OrderItem) error
		UpdateOrderItem(orderItem *model.OrderItem, id uint, orderID uint, organizationID uint, restaurantID uint) error
		DeleteOrderItem(orderItem *model.OrderItem, id uint, orderID uint, organizationID uint, restaurantID uint) error
	}

	orderItemRepository struct{
		db *gorm.DB
	}
)

func NewOrderItemRepository(db *gorm.DB) OrderItemRepository{
	return &orderItemRepository{db}
}

func (or *orderItemRepository) GetOrderItemList(orderItem *[]model.OrderItem, organizationID uint, restaurantID uint, orderID uint) error{

	if err := or.db.Where("organization_id=? and restaurant_id=? and order_id=? and is_deleted=0",organizationID, restaurantID, orderID).Find(orderItem).Error; err != nil{
		return err
	}
	return nil
}

func (or *orderItemRepository) CreateOrderItem(orderItem *model.OrderItem) error{

	if err := or.db.Create(orderItem).Error; err != nil{
		return err
	}
	return nil
}

func (or *orderItemRepository) UpdateOrderItem(orderItem *model.OrderItem, id uint, orderID uint, organizationID uint, restaurantID uint) error{

	result := or.db.Model(orderItem).Where("id=? and order_id=? and organization_id=? and restaurant_id=?", id, orderID, organizationID, restaurantID).Updates(orderItem)

	if err := result.Error; err != nil{
		return err
	}

	if result.RowsAffected < 1{
		return fmt.Errorf("record not found")
	}
	
	return nil
}

func (or *orderItemRepository) DeleteOrderItem(orderItem *model.OrderItem, id uint, orderID uint, organizationID uint, restaurantID uint) error{

	result := or.db.Model(orderItem).Where("id=? and order_id=? and organization_id=? and restaurant_id=?", id, orderID, organizationID, restaurantID).Update("is_deleted", 1)

	if err := result.Error; err != nil{
		return err
	}

	if result.RowsAffected < 1{
		return fmt.Errorf("record not found")
	}
	
	return nil
}