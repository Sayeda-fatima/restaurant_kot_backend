package repository

import (
	"github.com/NazishAhsan/easy_busy_book_go/model"
	"gorm.io/gorm"
)

type (
	CartRepository interface {
		GetCartList(carts *[]model.Cart, organizationID uint) error
		CreateCart(cart *model.Cart) error
		UpdateCart(cart *model.Cart, id uint) error
		DeleteCart(id uint) error
	}

	cartRepository struct {
		db *gorm.DB
	}
)

func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{db}
}

func (cr *cartRepository) GetCartList(carts *[]model.Cart, organizationID uint) error {

	if err := cr.db.Where("organization_id=? and is_deleted=0", organizationID).Find(carts).Error; err != nil {
		return err
	}
	return nil
}

func (cr *cartRepository) CreateCart(cart *model.Cart) error {

	if err := cr.db.Create(cart).Error; err != nil {
		return err
	}
	return nil
}

func (cr *cartRepository) UpdateCart(cart *model.Cart, id uint) error {

	result := cr.db.Model(cart).Where("id=?", id).Updates(cart)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func (cr *cartRepository) DeleteCart(id uint) error {

	result := cr.db.Where("id=?", id).Delete(&model.Cart{})

	if err := result.Error; err != nil {
		return err
	}
	return nil
}
