package repository

import (
	"fmt"

	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/common"
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/model"
	"gorm.io/gorm"
)

type (
	OrderItemRepository interface {
		GetOrderItemList(orderItem *[]model.OrderItem, organizationID uint, restaurantID uint, orderID uint) error
		CreateOrderItem(orderItem *model.OrderItem) error
		UpdateOrderItem(orderItem *model.OrderItem, id uint, orderID uint, organizationID uint, restaurantID uint) error
		DeleteOrderItem(orderItem *model.OrderItem, id uint, orderID uint, organizationID uint, restaurantID uint) error
		MostOrderedItems(result *[]map[string]interface{}, organizationID uint, restaurantID uint, dateFrom string, dateTo string) error
		DailySaleByItem(result *[]map[string]interface{}, organizationID uint, restaurantID uint, dateFrom string, dateTo string, page int) error
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

func (or *orderItemRepository) MostOrderedItems(result *[]map[string]interface{}, organizationID uint, restaurantID uint, dateFrom string, dateTo string) error{

	err := or.db.Raw(`SELECT order_items.organization_id, order_items.restaurant_id, menu_item_id, 
						sum(item_quantity) as total_quantity_sold, 
						sum(total_item_price)/100 as total_sale, order_items.created_at,
						menu_items.item_name, menu_items.price/100 as price
						from order_items left join menu_items 
						on order_items.menu_item_id=menu_items.id 
						where order_items.restaurant_id=? and order_items.organization_id=? and order_items.created_at between ? and ?
						group by menu_item_id 
						order by sum(item_quantity) desc limit 10`, restaurantID, organizationID, dateFrom, dateTo).Find(result).Error
	if err != nil{
		return err
	}
	
	return nil
}

func (or *orderItemRepository) DailySaleByItem(result *[]map[string]interface{}, organizationID uint, restaurantID uint, dateFrom string, dateTo string, page int) error{

	limit, offset := common.ApplyPagination(page)
	err := or.db.Raw(`SELECT menu_item_id, sum(item_quantity) as total_quantity_sold,
						sum(total_item_price)/100 as total_sale, order_items.created_at,
						menu_items.item_name, menu_items.price/100 as price
						from order_items left join menu_items
						on order_items.menu_item_id=menu_items.id
						where order_items.restaurant_id=? and order_items.organization_id=? and order_items.created_at between ? and ?
						group by menu_item_id
						order by sum(item_quantity) desc limit ? offset ?
					`, restaurantID, organizationID, dateFrom, dateTo, limit, offset).Find(result).Error
					
	if err != nil{
		return err
	}

	return nil
}