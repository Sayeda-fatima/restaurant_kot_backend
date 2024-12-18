package repository

import (
	"fmt"

	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/model"
	"gorm.io/gorm"
)

type (
	ProductRepository interface {
		GetAllProduct(products *[]model.Product, organizationID uint, restaurantID uint) error
		GetProduct(product *model.Product, id uint, organizationID uint, restaurantID uint) error
		CreateProduct(product *model.Product) error
		UpdateProduct(product *model.Product, id uint, organizationID uint, restaurantID uint) error
		UpdateProductQuantity(product *model.Product, id uint, organizationID uint, restaurantID uint, quantity int) error
		DeleteProduct(product *model.Product, id uint, organizationID uint, restaurantID uint) error
		InventoryValue(result *map[string]interface{}, organizationID uint, restaurantID uint) error
	}

	productRepository struct{
		db *gorm.DB
	}
)

func NewProductRepository(db *gorm.DB) ProductRepository{
	return &productRepository{db}
}

func (pr *productRepository) GetAllProduct(products *[]model.Product, organizationID uint, restaurantID uint) error{

	if err := pr.db.Where("organization_id=? and restaurant_id=? and is_deleted=0", organizationID, restaurantID).Find(products).Error; err != nil{
		return err
	}
	return nil
}

func (pr *productRepository) GetProduct(product *model.Product, id uint, organizationID uint, restaurantID uint) error{

	if err := pr.db.Where("id=? and restaurant_id=? and organization_id=?", id, restaurantID, organizationID).First(product).Error; err != nil{
		return err
	}
	return nil
}

func (pr *productRepository) CreateProduct(product *model.Product) error{

	if err := pr.db.Create(product).Error; err != nil{
		return err
	}
	return nil
}

func (pr *productRepository) UpdateProduct(product *model.Product, id uint, organizationID uint, restaurantID uint) error{

	result := pr.db.Model(product).Where("id=? and organization_id=? and restaurant_id=?", id, organizationID, restaurantID).Updates(product)

	if err := result.Error; err != nil{
		return err
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	
	return nil
}

func (pr *productRepository) UpdateProductQuantity(product *model.Product, id uint, organizationID uint, restaurantID uint, quantity int) error{

	result := pr.db.Model(product).Where("id=? and organization_id=? and restaurant_id=?", id, organizationID, restaurantID).Update("quantity", quantity)

	if err := result.Error; err != nil{
		return err
	}
	return nil
}

func (pr *productRepository) DeleteProduct(product *model.Product, id uint, organizationID uint, restaurantID uint) error{

	result := pr.db.Model(product).Where("id=? and organization_id=? and restaurant_id=?", id, organizationID, restaurantID).Update("is_deleted", 1)

	if err := result.Error; err != nil{
		return err
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	
	return nil
}

func (pr *productRepository) InventoryValue(result *map[string]interface{}, organizationID uint, restaurantID uint) error{

	err := pr.db.Raw(`SELECT sum(inventory_value)/100 as total_inventory_value from products
							where organization_id=? and restaurant_id=?
						`, organizationID, restaurantID).Find(result).Error
						
	if err != nil{
		return err
	}
	return nil
}