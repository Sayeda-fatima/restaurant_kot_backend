package repository

import (
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/model"
	"gorm.io/gorm"
)

type (
	InventoryTransactionRepository interface {
		GetInventoryTransactionList(inventoryTransaction *[]model.InventoryTransaction, organizationID uint, restaurantID uint) error
		CreateInventoryTransaction(inventoryTransaction *model.InventoryTransaction) error
		UpdateInventoryTransaction(inventoryTransaction *model.InventoryTransaction, id uint, organizationID uint, restaurantID uint) error
		GetCostOfGoodsSold(result *map[string]interface{}, organizationID uint, restaurantID uint, dateFrom string, dateTo string) error
	}

	inventoryTransactionRepository struct{
		db *gorm.DB
	}
)

func NewInventoryTransactionRepository(db *gorm.DB) InventoryTransactionRepository{
	return &inventoryTransactionRepository{db}
}

func (ir *inventoryTransactionRepository) GetInventoryTransactionList(inventoryTransaction *[]model.InventoryTransaction, organizationID uint, restaurantID uint) error{

	if err := ir.db.Where("organization_id=? and restaurant_id=?", organizationID, restaurantID).Find(inventoryTransaction).Error; err != nil{
		return err
	}

	return nil
}

func (ir *inventoryTransactionRepository) CreateInventoryTransaction(inventoryTransaction *model.InventoryTransaction) error{

	if err := ir.db.Create(inventoryTransaction).Error; err != nil{
		return err
	}

	return nil
}

func (ir *inventoryTransactionRepository) UpdateInventoryTransaction(inventoryTransaction *model.InventoryTransaction, id uint, organizationID uint, restaurantID uint) error{

	result := ir.db.Model(inventoryTransaction).Where("id=? and organization_id=? and restaurant_id=?", id, organizationID, restaurantID).Updates(inventoryTransaction)

	if err := result.Error; err != nil{
		return err
	}

	return nil
}

func (ir *inventoryTransactionRepository) GetCostOfGoodsSold(result *map[string]interface{}, organizationID uint, restaurantID uint, dateFrom string, dateTo string) error{

	
	err := ir.db.Raw(`SELECT sum(total_cost) as cost_of_goods_sold from inventory_transactions
						where transaction_type='sale' and organization_id=? and restaurant_id=? and recorded_at between ? and ?
					`, organizationID, restaurantID, dateFrom, dateTo).Find(result).Error

	// err := ir.db.Raw(`SELECT sum(quantity*unit_cost) as starting_inventory from inventory_transactions
	// 					where organization_id=? and restaurant_id=? and recorded_at < ?
	// 				`, organizationID, restaurantID, dateFrom).Find(result).Error
					
	if err != nil{
		return err
	}

	// err2 := ir.db.Raw(`SELECT sum(total_cost) as purchases from inventory_transactions
	// 					where transaction_type='purchase' and organization_id=? and restaurant_id=? and recorded_at between ? and ?
	//  				`, organizationID, restaurantID, dateFrom, dateTo).Find(result).Error
	
	// if err2 != nil{
	// 	return err2
	// }

	// err3 := ir.db.Raw(`SELECT sum(total_cost) as sold_goods from inventory_transactions
	// 					where transaction_type='sale' and organization_id=? and restaurant_id=? and recorded_at between ? and ?
	// 				`, organizationID, restaurantID, dateFrom, dateTo).Find(result).Error	
					
	// if err3 != nil{
	// 	return err3
	// }

	// err4 := ir.db.Raw(`SELECT sum(quantity*unit_cost) as ending_inventory from inventory_transactions
	// 					where organization_id=? and restaurant_id=? and recorded_at <= ?
	// 				`, organizationID, restaurantID, dateTo).Find(result).Error

	// if err4 != nil{
	// 	return err4
	// }

	// err5 := ir.db.Raw(`SELECT sum(total_cost) as adjustment from inventory_transactions
	// 					where transaction_type in ('adjustment','waste') and organization_id=? and restaurant_id=? and recorded_at between ? and ?
	// 				`, organizationID, restaurantID, dateFrom, dateTo).Find(result).Error
	
	// if err5 != nil{
	// 	return err5
	// }

	// err6 := ir.db.Raw(`SELECT 
	// 					SUM(CASE 
	// 						WHEN transaction_type = 'sale' THEN total_cost 
	// 						WHEN transaction_type = 'waste' THEN total_cost 
	// 						WHEN transaction_type = 'adjustment' AND total_cost < 0 THEN -total_cost
	// 						WHEN transaction_type = 'adjustment' AND total_cost > 0 THEN total_cost
	// 					END) AS cogs
	// 				FROM inventory_transactions
	// 				WHERE organization_id=? and restaurant_id=? and recorded_at BETWEEN ? AND ?;
	// 				`, organizationID, restaurantID, dateFrom, dateTo).Find(result).Error
					
	// if err6 != nil{
	// 	return err6
	// }
	return nil
}