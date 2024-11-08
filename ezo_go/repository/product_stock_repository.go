package repository

import (
	"github.com/NazishAhsan/easy_busy_book_go/model"
	"gorm.io/gorm"
)

type (
	ProductStockRepository interface {
		GetProductStockList(productStock *[]model.ProductStock, organizationID uint) error
		CreateProductStock(productStock *model.ProductStock) error
		UpdateProductStock(productStock *model.ProductStock, id uint) error
		DeleteProductStock(productStock *model.ProductStock, id uint) error
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

func (pr *productStockRepository) UpdateProductStock(productStock *model.ProductStock, id uint) error {

	result := pr.db.Model(productStock).Where("id=?", id).Updates(productStock)

	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func (pr *productStockRepository) DeleteProductStock(productStock *model.ProductStock, id uint) error {

	result := pr.db.Model(productStock).Where("id=?", id).Update("is_deleted", 1)

	if err := result.Error; err != nil {
		return err
	}
	return nil
}
