package repository

import (
	"fmt"

	"github.com/NazishAhsan/easy_busy_book_go/model"
	"gorm.io/gorm"
)

type (
	ProductStockRepository interface {
		GetProductStockList(productStock *[]model.ProductStock, organizationID uint) error
		CreateProductStock(productStock *model.ProductStock) error
		UpdateProductStock(productStock *model.ProductStock, id uint, organizationID uint) error
		DeleteProductStock(productStock *model.ProductStock, id uint, organizationID uint) error
		GetProductStockListByUpdateType(productStock *[]model.ProductStock, organizationID uint, term string) error
	}

	productStockRepository struct {
		db *gorm.DB
	}
)

func NewProductStockRepository(db *gorm.DB) ProductStockRepository {
	return &productStockRepository{db}
}

func (pr *productStockRepository) GetProductStockList(productStock *[]model.ProductStock, organizationID uint) error {

	if err := pr.db.Where("organization_id=? and is_deleted=0", organizationID).Find(productStock).Error; err != nil {
		return err
	}
	return nil
}

func (pr *productStockRepository) CreateProductStock(productStock *model.ProductStock) error {

	if err := pr.db.Create(productStock).Error; err != nil {
		return err
	}
	return nil
}

func (pr *productStockRepository) UpdateProductStock(productStock *model.ProductStock, id uint, organizationID uint) error {

	result := pr.db.Model(productStock).Where("id=? and organization_id=?", id, organizationID).Updates(productStock)

	if err := result.Error; err != nil {
		return err
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("record does not exist")
	}
	
	return nil
}

func (pr *productStockRepository) DeleteProductStock(productStock *model.ProductStock, id uint, organizationID uint) error {

	result := pr.db.Model(productStock).Where("id=? and organization_id=?", id, organizationID).Update("is_deleted", 1)

	if err := result.Error; err != nil {
		return err
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("record does not exist")
	}
	
	return nil
}

func (pr *productStockRepository) GetProductStockListByUpdateType(productStock *[]model.ProductStock, organizationID uint, term string) error{

	if err := pr.db.Where("organization_id=? and product_update_type=?", organizationID, term).Find(productStock).Error; err != nil{
		return err
	}
	return nil
}
