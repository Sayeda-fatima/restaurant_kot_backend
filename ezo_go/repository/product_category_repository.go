package repository

import (
	"fmt"

	"github.com/NazishAhsan/easy_busy_book_go/model"
	"gorm.io/gorm"
)

type (
	ProductCategoryRepository interface {
		GetProductCategoryList(productCategory *[]model.ProductCategory, organizationID uint) error
		CreateProductCategory(productCategory *model.ProductCategory) error
		UpdateProductCategory(productCategory *model.ProductCategory, id uint) error
		DeleteProductCategory(productCategory *model.ProductCategory, id uint) error
		SearchProductCategory(productCategory *[]model.ProductCategory, organizationID uint, term string) error
	}

	productCategoryRepository struct {
		db *gorm.DB
	}
)

func NewProductCategoryRepository(db *gorm.DB) ProductCategoryRepository {
	return &productCategoryRepository{db}
}

func (pr *productCategoryRepository) GetProductCategoryList(productCategory *[]model.ProductCategory, organizationID uint) error {

	if err := pr.db.Where("organization_id=? and is_deleted=0", organizationID).Find(productCategory).Error; err != nil {
		return err
	}
	return nil
}

func (pr *productCategoryRepository) CreateProductCategory(productCategory *model.ProductCategory) error {

	if err := pr.db.Create(productCategory).Error; err != nil {
		return err
	}
	return nil
}

func (pr *productCategoryRepository) UpdateProductCategory(productCategory *model.ProductCategory, id uint) error {

	result := pr.db.Model(productCategory).Where("id=?", id).Updates(productCategory)

	if err := result.Error; err != nil {
		return err
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("record does not exist")
	}
	
	return nil
}

func (pr *productCategoryRepository) DeleteProductCategory(productCategory *model.ProductCategory, id uint) error {

	result := pr.db.Model(productCategory).Where("id=?", id).Update("is_deleted", 1)

	if err := result.Error; err != nil {
		return err
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("record does not exist")
	}
	
	return nil
}

func (pr *productCategoryRepository) SearchProductCategory(productCategory *[]model.ProductCategory, organizationID uint, term string) error{

	if err := pr.db.Where("id LIKE ? or category LIKE ?", "%"+term+"%", "%"+term+"%").Having("organization_id=? and is_deleted=0", organizationID).Find(productCategory).Error; err!=nil{
		return err
	}
	return nil
}
