package repository

import (
	"github.com/NazishAhsan/easy_busy_book_go/model"
	"gorm.io/gorm"
)

type (
	CartItemRepository interface {
		GetCartItemList(cartItems *[]model.CartItem, organizationID uint, cartID uint) error
		CreateCartItem(cartItem *model.CartItem) error
		UpdateCartItem(cartItem *model.CartItem, id uint) error
		DeleteCartItem(cartItem *model.CartItem, id uint) error
	}

	cartItemRepository struct {
		db *gorm.DB
	}
)

func NewCartItemRepository(db *gorm.DB) CartItemRepository {
	return &cartItemRepository{db}
}

func (cr *cartItemRepository) GetCartItemList(cartItems *[]model.CartItem, organizationID uint, cartID uint) error {

	if err := cr.db.Where("organization_id=? and cart_id=?", organizationID, cartID).Find(cartItems).Error; err != nil {
		return err
	}
	return nil
}

func (cr *cartItemRepository) CreateCartItem(cartItem *model.CartItem) error {

	if err := cr.db.Create(cartItem).Error; err != nil {
		return err
	}
	return nil
}

func (cr *cartItemRepository) UpdateCartItem(cartItem *model.CartItem, id uint) error {

	result := cr.db.Model(cartItem).Where("id=?", id).Updates(cartItem)

	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func (cr *cartItemRepository) DeleteCartItem(cartItem *model.CartItem, id uint) error {

	result := cr.db.Model(cartItem).Where("id=?", id).Delete(cartItem)

	if err := result.Error; err != nil {
		return err
	}
	return nil
}
