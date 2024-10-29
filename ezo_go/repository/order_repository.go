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
	}

	orderRepository struct {
		db *gorm.DB
	}
)

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db}
}

func (ir *orderRepository) GetOrderList(orders *[]model.Order, organizationID uint) error {

	if err := ir.db.Where("organization_id=? and is_deleted=0", organizationID).Find(orders).Error; err != nil {
		return err
	}
	return nil
}

func (ir *orderRepository) CreateOrder(order *model.Order) error {

	if err := ir.db.Create(order).Error; err != nil {
		return err
	}
	return nil
}

func (ir *orderRepository) UpdateOrder(order *model.Order, id uint) error {

	result := ir.db.Model(order).Where("id=?", id).Updates(order)

	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func (ir *orderRepository) DeleteOrder(order *model.Order, id uint) error {

	result := ir.db.Model(order).Where("id=?", id).Update("is_deleted", 1)

	if err := result.Error; err != nil {
		return err
	}
	return nil
}
