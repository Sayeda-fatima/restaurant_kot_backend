package repository

import (
	"fmt"

	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/common"
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/model"
	"gorm.io/gorm"
)

type (
	RestaurantRepository interface {
		GetRestaurantList(restaurants *[]model.Restaurant, organizationID uint, page int) error
		CreateRestaurant(restaurant *model.Restaurant) error
		UpdateRestaurant(restaurant *model.Restaurant, id uint) error
		DeleteRestaurant(restaurant *model.Restaurant, id uint) error
	}

	restaurantRepository struct{
		db *gorm.DB
	}
)

func NewRestaurantRepository(db *gorm.DB) RestaurantRepository{
	return &restaurantRepository{db}
}

func (rr *restaurantRepository) GetRestaurantList(restaurants *[]model.Restaurant, organizationID uint, page int) error{

	if err := rr.db.Scopes(common.Paginate(page)).Where("organization_id=? and is_deleted=0", organizationID).Find(restaurants).Error; err != nil{
		return err
	}
	return nil
}

func (rr *restaurantRepository) CreateRestaurant(restaurant *model.Restaurant) error{

	if err := rr.db.Create(restaurant).Error; err != nil{
		return err
	}
	return nil
}

func (rr *restaurantRepository) UpdateRestaurant(restaurant *model.Restaurant, id uint) error{

	result := rr.db.Model(restaurant).Where("id=?", id).Updates(restaurant)

	if err := result.Error; err != nil{
		return err
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	
	return nil
}

func (rr *restaurantRepository) DeleteRestaurant(restaurant *model.Restaurant, id uint) error{

	result := rr.db.Model(restaurant).Where("id=?", id).Update("is_deleted", 1)

	if err := result.Error; err != nil{
		return err
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	
	return nil
}