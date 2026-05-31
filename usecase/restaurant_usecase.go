package usecase

import (
	"github.com/Sayeda-fatima/restaurant_kot_backend/model"
	"github.com/Sayeda-fatima/restaurant_kot_backend/repository"
	"github.com/Sayeda-fatima/restaurant_kot_backend/validator"
)

type (
	RestaurantUsecase interface {
		GetRestaurantList(organizationID uint, page int) ([]model.RestaurantResponse, error)
		CreateRestaurant(restaurant model.Restaurant) (model.RestaurantResponse, error)
		UpdateRestaurant(restaurant model.Restaurant, id uint) (model.RestaurantResponse, error)
		DeleteRestaurant(restaurant model.Restaurant, id uint) error
	}

	restaurantUsecase struct{
		rr repository.RestaurantRepository
		rv validator.RestaurantValidator
	}
)

func NewRestaurantUsecase (rr repository.RestaurantRepository, rv validator.RestaurantValidator) RestaurantUsecase{
	return &restaurantUsecase{rr, rv}
}

func (ru *restaurantUsecase) GetRestaurantList(organizationID uint, page int) ([]model.RestaurantResponse, error){

	restaurants := []model.Restaurant{}
	if err := ru.rr.GetRestaurantList(&restaurants, organizationID, page); err != nil{
		return nil, err
	}

	resRestaurants := []model.RestaurantResponse{}
	for _, v := range(restaurants){
		res := model.RestaurantResponse{
			ID: v.ID,
			OrganizationID: v.OrganizationID,
			Name: v.Name,
			PhoneNo: v.PhoneNo,
			Email: v.Email,
			Address: v.Address,
		}
		resRestaurants = append(resRestaurants, res)
	}

	return resRestaurants, nil
}

func (ru *restaurantUsecase) CreateRestaurant(restaurant model.Restaurant) (model.RestaurantResponse, error){

	if err := ru.rv.RestaurantValidate(restaurant); err != nil{
		return model.RestaurantResponse{}, err
	}

	if err := ru.rr.CreateRestaurant(&restaurant); err != nil{
		return model.RestaurantResponse{}, err
	}

	resRestaurant := model.RestaurantResponse{
		ID: restaurant.ID,
		OrganizationID: restaurant.OrganizationID,
		Name: restaurant.Name,
		PhoneNo: restaurant.PhoneNo,
		Email: restaurant.Email,
		Address: restaurant.Address,
	}

	return resRestaurant, nil
}

func (ru *restaurantUsecase) UpdateRestaurant(restaurant model.Restaurant, id uint) (model.RestaurantResponse, error){

	if err := ru.rv.RestaurantValidate(restaurant); err != nil{
		return model.RestaurantResponse{}, err
	}

	if err := ru.rr.UpdateRestaurant(&restaurant, id); err != nil{
		return model.RestaurantResponse{}, err
	}

	resRestaurant := model.RestaurantResponse{
		ID: restaurant.ID,
		OrganizationID: restaurant.OrganizationID,
		Name: restaurant.Name,
		PhoneNo: restaurant.PhoneNo,
		Email: restaurant.Email,
		Address: restaurant.Address,
	}

	return resRestaurant, nil
}

func (ru *restaurantUsecase) DeleteRestaurant(restaurant model.Restaurant, id uint) error{

	if err := ru.rr.DeleteRestaurant(&restaurant, id); err != nil{
		return err
	}
	return nil
}
