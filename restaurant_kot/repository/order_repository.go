package repository

import (
	"fmt"

	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/model"
	"gorm.io/gorm"
)

type (
	OrderRepository interface {
		GetOrderList(orders *[]model.Order, organizationID uint, restaurantID uint) error
		CreateOrder(order *model.Order) error
		UpdateOrder(order *model.Order, id uint, organizationID uint, restaurantID uint) error
		DeleteOrder(order *model.Order, id uint, organizationID uint, restaurantID uint) error
	}

	orderRepository struct{
		db *gorm.DB
	}
)

func NewOrderRepository(db *gorm.DB) OrderRepository{
	return &orderRepository{db}
}

func (or *orderRepository) GetOrderList(orders *[]model.Order, organizationID uint, restaurantID uint) error{

	if err := or.db.Where("organization_id=? and restaurant_id=? and is_deleted=0", organizationID, restaurantID).Find(orders).Error; err != nil{
		return err
	}
	return nil
}

func (or *orderRepository) CreateOrder(order *model.Order) error{

	if err := or.db.Create(order).Error; err != nil{
		return err
	}
	return nil
}

func (or *orderRepository) UpdateOrder(order *model.Order, id uint, organizationID uint, restaurantID uint) error{

	result := or.db.Model(order).Where("id=? and organization_id=? and restaurant_id=?", id, organizationID, restaurantID).Updates(order)

	if err := result.Error; err != nil{
		return err
	}

	if result.RowsAffected < 1{
		return fmt.Errorf("record not found")
	}
	
	return nil
}

func (or *orderRepository) DeleteOrder(order *model.Order, id uint, organizationID uint, restaurantID uint) error{

	result := or.db.Model(order).Where("id=? and organization_id=? and restaurant_id=?", id, organizationID, restaurantID).Update("is_deleted", 1)

	if err := result.Error; err != nil{
		return err
	}

	if result.RowsAffected < 1{
		return fmt.Errorf("record not found")
	}

	return nil
}