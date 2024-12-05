package repository

import (
	"fmt"

	"github.com/NazishAhsan/easy_busy_book_go/model"
	"gorm.io/gorm"
)

type (
	ProductImageRepository interface {
		GetProductImageList(productImage *[]model.ProductImage, organizationID uint, productID uint) error
		AddProductImage(productImage *model.ProductImage) error
		DeleteProductImage(productImage *model.ProductImage, id uint) error
	}

	productImageRepository struct {
		db *gorm.DB
	}
)

func NewProductImageRepository(db *gorm.DB) ProductImageRepository {
	return &productImageRepository{db}
}

func (pr *productImageRepository) GetProductImageList(productImage *[]model.ProductImage, organizationID uint, productID uint) error {

	if err := pr.db.Where("organization_id=? and product_id=? and is_deleted=0", organizationID, productID).Find(productImage).Error; err != nil {
		return err
	}
	return nil
}

func (pr *productImageRepository) AddProductImage(productImage *model.ProductImage) error {

	if err := pr.db.Create(productImage).Error; err != nil {
		return err
	}
	return nil
}

func (pr *productImageRepository) DeleteProductImage(productImage *model.ProductImage, id uint) error {

	result := pr.db.Model(productImage).Where("id=?", id).Update("is_deleted", 1)

	if err := result.Error; err != nil {
		return err
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("record does not exist")
	}
	
	return nil
}
