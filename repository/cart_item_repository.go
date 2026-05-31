package repository

import (
	"fmt"

	"github.com/Sayeda-fatima/restaurant_kot_backend/model"
	"gorm.io/gorm"
)

type (
	CartItemRepository interface {
		GetCartItemList(cartItem *[]model.CartItem, cartID uint, restaurantID uint, organizationID uint) error
		GetPendingCartItemList(cartItem *[]model.CartItem, cartID uint, restaurantID uint, organizationID uint) error
		CreateCartitem(cartItem *model.CartItem) error
		UpdateCartitem(cartItem *model.CartItem, id uint, cartID uint, restaurantID uint, organizationID uint) error
		UpdateCartItemStatus(cartItem *model.CartItem, id uint, cartID uint, restaurantID uint, organizationID uint, status string) error
		DeleteCartItem(cartItem *model.CartItem, id uint, cartID uint, restaurantID uint, organizationID uint) error
	}

	cartItemRepository struct{
		db *gorm.DB
	}
)

func NewCartItemRepository(db *gorm.DB) CartItemRepository{
	return &cartItemRepository{db}
}

func (cr *cartItemRepository) GetCartItemList(cartItem *[]model.CartItem, cartID uint, restaurantID uint, organizationID uint) error{

	if err := cr.db.Preload("MenuItem").Where("organization_id=? and restaurant_id=? and cart_id=?", organizationID, restaurantID, cartID).Find(cartItem).Error; err != nil{
		return err
	}

	return nil
}

func (cr *cartItemRepository) GetPendingCartItemList(cartItem *[]model.CartItem, cartID uint, restaurantID uint, organizationID uint) error{

	if err := cr.db.Preload("MenuItem.Recipe.RecipeProducts", "is_deleted=?", "0").Preload("MenuItem.Recipe.RecipeProducts.Product").Where("cart_id=? and item_status='pending' and restaurant_id=? and organization_id=?", cartID, restaurantID, organizationID).Find(cartItem).Error; err != nil{
		return err
	}

	return nil
}

func (cr *cartItemRepository) CreateCartitem(cartItem *model.CartItem) error{

	if err := cr.db.Create(cartItem).Error; err != nil{
		return err
	}
	return nil
}

func (cr *cartItemRepository) UpdateCartitem(cartItem *model.CartItem, id uint, cartID uint, restaurantID uint, organizationID uint) error{

	result := cr.db.Model(cartItem).Where("id=? and cart_id=? and restaurant_id=? and organization_id=?", id, cartID, restaurantID, organizationID).Updates(cartItem)

	if err := result.Error; err != nil{
		return err
	}

	if result.RowsAffected < 1{
		return fmt.Errorf("record not found")
	}

	return nil
}

func (cr *cartItemRepository) UpdateCartItemStatus(cartItem *model.CartItem, id uint, cartID uint, restaurantID uint, organizationID uint, status string) error{

	result := cr.db.Model(cartItem).Where("id=? and cart_id=? and restaurant_id=? and organization_id=?", id, cartID, restaurantID, organizationID).Update("item_status", status)

	if err := result.Error; err != nil{
		return err
	}

	if result.RowsAffected < 1{
		return fmt.Errorf("record not found")
	}

	return nil
}

func (cr *cartItemRepository) DeleteCartItem(cartItem *model.CartItem, id uint, cartID uint, restaurantID uint, organizationID uint) error{

	result := cr.db.Model(cartItem).Where("id=? and cart_id=? and restaurant_id=? and organization_id=?", id, cartID, restaurantID, organizationID).Delete(cartItem)

	if err := result.Error; err != nil{
		return err
	}

	if result.RowsAffected < 1{
		return fmt.Errorf("record not found")
	}
	
	return nil
}