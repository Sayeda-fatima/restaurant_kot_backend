package repository

import (
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/model"
	"gorm.io/gorm"
)

type (
	ProductRepository interface {
		GetAllProduct(products *[]model.Product, organizationID uint, restaurantID uint) error
		CreateProduct(product *model.Product) error
		UpdateProduct(product *model.Product, id uint) error
		DeleteProduct(product *model.Product, id uint) error
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

func (pr *productRepository) CreateProduct(product *model.Product) error{

	if err := pr.db.Create(product).Error; err != nil{
		return err
	}
	return nil
}

func (pr *productRepository) UpdateProduct(product *model.Product, id uint) error{

	result := pr.db.Model(product).Where("id=?", id).Updates(product)

	if err := result.Error; err != nil{
		return err
	}
	return nil
}

func (pr *productRepository) DeleteProduct(product *model.Product, id uint) error{

	result := pr.db.Model(product).Where("id=?", id).Update("is_deleted", 1)

	if err := result.Error; err != nil{
		return err
	}
	return nil
}