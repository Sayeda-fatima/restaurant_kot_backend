package repository

import (
	"fmt"

	"github.com/Sayeda-fatima/restaurant_kot_backend/model"
	"gorm.io/gorm"
)

type (
	OrderRepository interface {
		GetOrderList(orders *[]model.Order, organizationID uint, restaurantID uint) error
		CreateOrder(order *model.Order) error
		UpdateOrder(order *model.Order, id uint, organizationID uint, restaurantID uint) error
		DeleteOrder(order *model.Order, id uint, organizationID uint, restaurantID uint) error
		TotalSales(result *map[string]interface{}, organizationID uint, restaurantID uint, dateFrom string, dateTo string) error
		TotalSalesByOrderType(result *[]map[string]interface{}, organizationID uint, restaurantID uint, dateFrom string, dateTo string) error
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

func (or *orderRepository) TotalSales(result *map[string]interface{}, organizationID uint, restaurantID uint, dateFrom string, dateTo string) error{

	err := or.db.Raw(`SELECT count(orders.id) as total_sales, sum(total_price)/100 as total_sales_amount, sum(item_quantity) as total_items_sold
					from orders
					left join order_items on orders.id = order_items.order_id 
					where orders.restaurant_id=? and orders.organization_id=? and orders.created_at between ? and ?
					order by orders.created_at desc
					`, organizationID, restaurantID, dateFrom, dateTo).Find(result).Error

	if err != nil{
		return err
	}
	return nil
}

func (or *orderRepository) TotalSalesByOrderType(result *[]map[string]interface{}, organizationID uint, restaurantID uint, dateFrom string, dateTo string) error{

	err := or.db.Raw(`SELECT count(orders.id) as total_sales, sum(total_price)/100 as total_sales_amount, sum(item_quantity) as total_items_sold, orders.order_type
					from orders
					left join order_items on orders.id = order_items.order_id
					where orders.restaurant_id=? and orders.organization_id=? and orders.created_at between ? and ?
					group by orders.order_type
					order by orders.created_at desc
					`, organizationID, restaurantID, dateFrom, dateTo).Find(result).Error
					
	if err != nil{
		return err
	}

	return nil
}

func (or *orderRepository) GrossProfit(result *map[string]interface{}, organizationID uint, restaurantID uint, dateFrom string, dateTo string) error{

	err := or.db.Raw(`SELECT sum(total_price) as total_revenue, sum(total_cost) as cost_of_goods_sold,
						from orders 
						where organization_id=? and restaurant_id=? and created_at between ? and ? 
					`, organizationID, restaurantID, dateFrom, dateTo).Find(result).Error
					
	if err != nil{
		return err
	}

	return nil
}