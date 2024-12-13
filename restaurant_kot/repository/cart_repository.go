package repository

import (
	"fmt"

	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/model"
	"gorm.io/gorm"
)

type (
	CartRepository interface {
		GetCartList(carts *[]model.Cart, organizationID uint, restaurantID uint) error
		GetCart(cart *model.Cart, id uint, organizationID uint, restaurantID uint) error
		CreateCart(cart *model.Cart) error
		UpdateCart(cart *model.Cart, id uint, organizationID uint, restaurantID uint) error
		UpdateCartStatus(cart *model.Cart, id uint, organizationID uint, restaurantID uint, status string) error
		DeleteCart(id uint, organizationID uint, restaurantID uint) error
		CheckCartActive(cart *model.Cart, organizationID uint, restaurantID uint, tableID uint) error
	}

	cartRepository struct{
		db *gorm.DB
	}
)

func NewCartRepository(db *gorm.DB) CartRepository{
	return &cartRepository{db}
}

func (cr *cartRepository) GetCartList(carts *[]model.Cart, organizationID uint, restaurantID uint) error{

	if err := cr.db.Where("organization_id=? and restaurant_id=?", organizationID, restaurantID).Find(carts).Error; err != nil{
		return err
	}
	return nil
}

func (cr *cartRepository) GetCart(cart *model.Cart, id uint, organizationID uint, restaurantID uint) error{

	if err := cr.db.Preload("CartItems.MenuItem.Recipe.RecipeProducts.Product").Where("id=? and restaurant_id=? and organization_id=?", id, restaurantID, organizationID).First(cart).Error; err != nil{
		return err
	}
	return nil
}

func (cr *cartRepository) CreateCart(cart *model.Cart) error{

	if err := cr.db.Create(cart).Error; err != nil{
		return err
	}
	return nil
}

func (cr *cartRepository) UpdateCart(cart *model.Cart, id uint, organizationID uint, restaurantID uint) error{

	result := cr.db.Model(cart).Where("id=? and restaurant_id=? and organization_id=?", id, restaurantID, organizationID).Updates(cart)

	if err := result.Error; err != nil{
		return err
	}

	if result.RowsAffected < 1{
		return fmt.Errorf("record not found")
	}
	return nil
}

func (cr *cartRepository) UpdateCartStatus(cart *model.Cart, id uint, organizationID uint, restaurantID uint, status string) error{

	result := cr.db.Model(cart).Where("id=? and restaurant_id=? and organization_id=?", id, restaurantID, organizationID).Update("cart_status", status)

	if err := result.Error; err != nil{
		return err
	}

	if result.RowsAffected < 1{
		return fmt.Errorf("record not found")
	}
	return nil
}

func (cr *cartRepository) DeleteCart(id uint, organizationID uint, restaurantID uint) error{

	result := cr.db.Where("id=? and restaurant_id=? and organization_id=?", id, restaurantID, organizationID).Delete(&model.Cart{})

	if err := result.Error; err != nil{
		return err
	}

	if result.RowsAffected < 1{
		return fmt.Errorf("record not found")
	}
	return nil
}

func (cr *cartRepository) CheckCartActive(cart *model.Cart, organizationID uint, restaurantID uint, tableID uint) error{

	if err := cr.db.Preload("CartItems.MenuItem").Where("organization_id=? and restaurant_id=? and table_id=? and cart_status='active'", organizationID, restaurantID, tableID).First(cart).Error; err != nil{
		return err
	}
	return nil
}