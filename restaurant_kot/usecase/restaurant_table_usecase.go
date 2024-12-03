package usecase

import (
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/model"
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/repository"
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/validator"
)

type (
	RestaurantTableUsecase interface {
		GetRestaurantTableList(organizationID uint, restaurantID uint) ([]model.RestaurantTableResponse, error)
		CreateRestaurantTable(restaurantTable model.RestaurantTable) (model.RestaurantTableResponse, error)
		UpdateRestaurantTable(restaurantTable model.RestaurantTable, id uint) (model.RestaurantTableResponse, error)
		DeleteRestaurantTable(restaurantTable model.RestaurantTable, id uint) error
	}

	restaurantTableUsecase struct{
		rr repository.RestaurantTableRepository
		rv validator.RestaurantTableValidator
	}
)

func NewRestaurantTableUsecase (rr repository.RestaurantTableRepository, rv validator.RestaurantTableValidator) RestaurantTableUsecase{
	return &restaurantTableUsecase{rr,rv}
}

func (ru *restaurantTableUsecase) GetRestaurantTableList(organizationID uint, restaurantID uint) ([]model.RestaurantTableResponse, error){

	restaurantTables := []model.RestaurantTable{}

	if err := ru.rr.GetRestaurantTableList(&restaurantTables, organizationID, restaurantID); err != nil{
		return nil, err
	}

	resRestaurantTable := []model.RestaurantTableResponse{}
	for _, v := range(restaurantTables){
		res := model.RestaurantTableResponse{
			ID : v.ID,
			OrganizationID: v.OrganizationID,
			RestaurantID: v.RestaurantID,
			Capacity: v.Capacity,
			Status: v.Status,
		}
		resRestaurantTable = append(resRestaurantTable, res)
	}
	return resRestaurantTable, nil
}

func (ru *restaurantTableUsecase) CreateRestaurantTable(restaurantTable model.RestaurantTable) (model.RestaurantTableResponse, error){

	if err := ru.rv.RestaurantTableValidate(&restaurantTable); err != nil{
		return model.RestaurantTableResponse{}, err
	}

	if err := ru.rr.CreateRestaurantTable(&restaurantTable); err != nil{
		return model.RestaurantTableResponse{}, err
	}

	resRestaurantTable := model.RestaurantTableResponse{
		ID: restaurantTable.ID,
		OrganizationID: restaurantTable.OrganizationID,
		RestaurantID: restaurantTable.RestaurantID,
		Capacity: restaurantTable.Capacity,
		Status: restaurantTable.Status,
	}

	return resRestaurantTable, nil
}

func (ru *restaurantTableUsecase) UpdateRestaurantTable(restaurantTable model.RestaurantTable, id uint) (model.RestaurantTableResponse, error){

	if err := ru.rv.RestaurantTableValidate(&restaurantTable); err != nil{
		return model.RestaurantTableResponse{}, err
	}

	if err := ru.rr.UpdateRestaurantTable(&restaurantTable, id); err != nil{
		return model.RestaurantTableResponse{}, err
	}

	resRestaurantTable := model.RestaurantTableResponse{
		ID: restaurantTable.ID,
		OrganizationID: restaurantTable.OrganizationID,
		RestaurantID: restaurantTable.RestaurantID,
		Capacity: restaurantTable.Capacity,
		Status: restaurantTable.Status,
	}
	return resRestaurantTable, nil
}

func (ru *restaurantTableUsecase) DeleteRestaurantTable(restaurantTable model.RestaurantTable, id uint) error{

	if err := ru.rr.DeleteRestaurantTable(&restaurantTable, id); err != nil{
		return err
	}

	return nil
}
