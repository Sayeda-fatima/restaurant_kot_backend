package repository

import (
	"fmt"

	"github.com/NazishAhsan/easy_busy_book_go/common"
	"github.com/NazishAhsan/easy_busy_book_go/model"
	"gorm.io/gorm"
)

type (
	OrderRepository interface {
		GetOrderList(orders *[]model.Order, organizationID uint) error
		CreateOrder(order *model.Order) error
		UpdateOrder(order *model.Order, id uint, organizationID uint) error
		DeleteOrder(order *model.Order, id uint, organizationID uint) error
		InvoiceReportCustomer(order *[]model.Order, organizationID uint, customerID uint, dateFrom string, dateTo string) error
		GetInvoice(order *model.Order, id uint) error
		SaleReport(result *[]map[string]interface{}, organizationID uint, dateFrom string, dateTo string, page int) error
		TotalSales(result *map[string]interface{}, organizationID uint, dateFrom string, dateTo string) error
		TotalSalesQuantity(result *map[string]interface{}, organizationID uint, dateFrom string, dateTo string) error
		TotalSalesAmount(result *map[string]interface{}, organizationID uint, dateFrom string, dateTo string) error
		ProfitReport(result *[]map[string]interface{}, organizationID uint, dateFrom string, dateTo string, page int) error
		TotalProfit(result *map[string]interface{}, organizationID uint, dateFrom string, dateTo string) error
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

func (or *orderRepository) UpdateOrder(order *model.Order, id uint, organizationID uint) error {

	result := or.db.Model(order).Where("id=? and organization_id=?", id, organizationID).Updates(order)

	if err := result.Error; err != nil {
		return err
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("record does not exist")
	}
	
	return nil
}

func (or *orderRepository) DeleteOrder(order *model.Order, id uint, organizationID uint) error {

	result := or.db.Model(order).Where("id=? and organization_id=?", id, organizationID).Update("is_deleted", 1)

	if err := result.Error; err != nil {
		return err
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("record does not exist")
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

func (or *orderRepository) SaleReport(result *[]map[string]interface{}, organizationID uint, dateFrom string, dateTo string, page int) error{

	limit, offset := common.ApplyPagination(page)
	err := or.db.Raw(`SELECT o.created_at as date, 
            o.id as invoice_no, 
            customers.name, 
            customers.phone_no, 
            o.total_price, 
            (select sum(order_items.product_quantity) from order_items where order_items.order_id=o.id group by order_items.order_id) as total_quantity
            from orders as o
            left join customers on o.customer_id = customers.id 
            where o.organization_id=? and date(o.created_at) between ? and ? limit ? offset ?`, organizationID, dateFrom, dateTo, limit, offset).Find(result).Error
	if err != nil{
		return err
	}
	return nil
}

func (or *orderRepository) TotalSales(result *map[string]interface{}, organizationID uint, dateFrom string, dateTo string) error{

	err := or.db.Raw(`SELECT count(id) as total_sales from orders where organization_id=? and date(created_at) between ? and ?`, organizationID, dateFrom, dateTo).Find(result).Error

	if err != nil{
		return err
	}
	return nil
}

func (or *orderRepository) TotalSalesQuantity(result *map[string]interface{}, organizationID uint, dateFrom string, dateTo string) error{

	err := or.db.Raw(`SELECT sum(product_quantity) as total_sales_quantity from order_items where organization_id=? and date(created_at) between ? and ?`, organizationID, dateFrom, dateTo).Find(result).Error

	if err != nil{
		return err
	}
	return nil
}

func (or *orderRepository) TotalSalesAmount(result *map[string]interface{}, organizationID uint, dateFrom string, dateTo string) error{

	err := or.db.Raw(`SELECT sum(total_price) as total_sales_amount from orders where organization_id=? and date(created_at) between ? and ?`, organizationID, dateFrom, dateTo).Find(result).Error

	if err != nil{
		return err
	}
	return err
}

func (or *orderRepository) ProfitReport(result *[]map[string]interface{}, organizationID uint, dateFrom string, dateTo string, page int) error{

	limit, offset := common.ApplyPagination(page)

	err := or.db.Raw(`SELECT date(o.created_at) as date, 
                   o.id as invoice_no, 
                   customers.name as customer_name, 
                   o.total_price,
                   SUM(p.purchase_price * od.product_quantity) AS purchase_price, 
                   (o.total_price - SUM(p.purchase_price * od.product_quantity)) AS profit
            FROM orders AS o
            LEFT JOIN order_items AS od ON od.order_id = o.id
			LEFT JOIN customers on customers.id = o.customer_id
            LEFT JOIN products AS p ON p.id = od.product_id
            WHERE o.organization_id =? and date(o.created_at) BETWEEN ? AND ?
            GROUP BY o.id limit ? offset ?`, organizationID, dateFrom, dateTo, limit, offset).Find(result).Error

	if err != nil{
		return err
	}
	return nil
}

func (or *orderRepository) TotalProfit(result *map[string]interface{}, organizationID uint, dateFrom string, dateTo string) error{

	err := or.db.Raw(`SELECT SUM(o.total_price) as total_sales_amount,
            SUM(p.purchase_price * od.product_quantity) AS total_purchase_price, 
            (SUM(o.total_price) - SUM(p.purchase_price * od.product_quantity)) AS total_profit
            FROM orders AS o
            LEFT JOIN order_items AS od ON od.order_id = o.id
            LEFT JOIN products AS p ON p.id = od.product_id
            WHERE o.organization_id =? and date(o.created_at) BETWEEN ? AND ?`, organizationID, dateFrom, dateTo).Find(result).Error

	if err != nil{
		return err
	}
	return nil
}