package repository

import (
	"github.com/NazishAhsan/easy_busy_book_go/common"
	"github.com/NazishAhsan/easy_busy_book_go/model"
	"gorm.io/gorm"
)

type ProductRepository interface {
	GetProductList(products *[]model.Product, organizationID uint) error
	CreateProduct(product *model.Product) error
	GetProduct(product *model.Product, id uint) error
	UpdateProduct(product *model.Product, id uint) error
	DeleteProduct(product *model.Product, id uint) error
	SearchProduct(products *[]model.Product, organizationID uint, term string) error
	UpdateStockOfProduct (product *model.Product, id uint, quantity int) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

func (pr *productRepository) GetProductList(products *[]model.Product, organizationID uint) error {

	if err := pr.db.Select("organization_id", "id", "name", "image", "sell_price", "category_id", "quantity").Where("organization_id=? and is_deleted=0", organizationID).Order("name").Find(products).Error; err != nil {
		return err
	}
	return nil
}

func (pr *productRepository) CreateProduct(product *model.Product) error {

	if err := pr.db.Create(product).Error; err != nil {
		return err
	}
	return nil
}

func (pr *productRepository) GetProduct(product *model.Product, id uint) error{

	if err := pr.db.Where("id=?", id).First(product).Error; err != nil{
		return err
	}
	return nil
}

func (pr *productRepository) UpdateProduct(product *model.Product, id uint) error {

	common.Logger.LogInfo().Fields(map[string]interface{}{"data": product}).Msg("test")
	result := pr.db.Model(product).Where("id=?", id).Updates(product)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func (pr *productRepository) DeleteProduct(product *model.Product, id uint) error {

	common.Logger.LogInfo().Fields(map[string]interface{}{"data": product}).Msg("test")
	result := pr.db.Model(product).Where("id=?", id).Update("is_deleted", 1)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func (pr *productRepository) SearchProduct(products *[]model.Product, organizationID uint, term string) error {

	if err := pr.db.Where("id LIKE ? or name LIKE ? or description LIKE ?", "%"+term+"%", "%"+term+"%", "%"+term+"%").Having("organization_id=? and is_deleted=0", organizationID).Find(products).Error; err != nil {
		return err
	}
	return nil
}

func (pr *productRepository) UpdateStockOfProduct(product *model.Product, id uint, quantity int) error{

	if err := pr.db.Model(product).Where("id=?", id).Update("quantity", quantity).Error; err != nil{
		return err
	}
	return nil
}
