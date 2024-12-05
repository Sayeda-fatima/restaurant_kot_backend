package usecase

import (
	"strconv"

	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/model"
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/repository"
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/validator"
)

type (
	MenuItemUsecase interface {
		GetMenuItemList(organizationID uint, restaurantID uint, menuID uint) ([]model.MenuItemResponse, error)
		CreateMenuItem(menuItem model.MenuItem) (model.MenuItemResponse, error)
		UpdateMenuItem(menuItem model.MenuItem, id uint) (model.MenuItemResponse, error)
		DeleteMenuItem(menuItem model.MenuItem, id uint) error
	}

	menuItemUsecase struct{
		mr repository.MenuItemRepository
		mv validator.MenuItemValidator
	}
)

func NewMenuItemUsecase (mr repository.MenuItemRepository, mv validator.MenuItemValidator) MenuItemUsecase{
	return &menuItemUsecase{mr,mv}
}

func (mu *menuItemUsecase) GetMenuItemList(organizationID uint, restaurantID uint, menuID uint) ([]model.MenuItemResponse, error){

	menuItems := []model.MenuItem{}
	if err := mu.mr.GetMenuItemList(&menuItems, organizationID, restaurantID, menuID); err != nil{
		return nil, err
	}

	resMenuItems := []model.MenuItemResponse{}
	for _, v := range(menuItems){
		res := model.MenuItemResponse{
			ID: v.ID,
			MenuID: v.MenuID,
			OrganizationID: v.OrganizationID,
			RestaurantID: v.RestaurantID,
			ItemName: v.ItemName,
			Description: v.Description,
			Currency: v.Currency,
			Price: strconv.FormatFloat(float64(v.Price)/100, 'f', -1, 64),
			RecipeID: v.RecipeID,
			MenuAllergens: v.MenuAllergens,
		}
		resMenuItems = append(resMenuItems, res)
	}
	return resMenuItems, nil
}

func (mu *menuItemUsecase) CreateMenuItem(menuItem model.MenuItem) (model.MenuItemResponse, error){

	if err := mu.mv.MenuItemValidate(&menuItem); err != nil{
		return model.MenuItemResponse{}, err
	}

	if err := mu.mr.CreateMenuItem(&menuItem); err != nil{
		return model.MenuItemResponse{}, err
	}

	resMenuItem := model.MenuItemResponse{
		ID: menuItem.ID,
		MenuID: menuItem.MenuID,
		OrganizationID: menuItem.OrganizationID,
		RestaurantID: menuItem.RestaurantID,
		ItemName: menuItem.ItemName,
		Description: menuItem.Description,
		Currency: menuItem.Currency,
		Price: strconv.FormatFloat(float64(menuItem.Price)/100, 'f', -1, 64),
		RecipeID: menuItem.RecipeID,
		MenuAllergens: menuItem.MenuAllergens,
	}

	return resMenuItem, nil
}

func (mu *menuItemUsecase) UpdateMenuItem(menuItem model.MenuItem, id uint) (model.MenuItemResponse, error){

	if err := mu.mv.MenuItemValidate(&menuItem); err != nil{
		return model.MenuItemResponse{}, err
	}

	if err := mu.mr.UpdateMenuItem(&menuItem, id); err != nil{
		return model.MenuItemResponse{}, err
	}

	resMenuItem := model.MenuItemResponse{
		ID: menuItem.ID,
		MenuID: menuItem.MenuID,
		OrganizationID: menuItem.OrganizationID,
		RestaurantID: menuItem.RestaurantID,
		ItemName: menuItem.ItemName,
		Description: menuItem.Description,
		Currency: menuItem.Currency,
		Price: strconv.FormatFloat(float64(menuItem.Price)/100, 'f', -1, 64),
		RecipeID: menuItem.RecipeID,
		MenuAllergens: menuItem.MenuAllergens,
	}

	return resMenuItem, nil
}

func (mu *menuItemUsecase) DeleteMenuItem(menuItem model.MenuItem, id uint) error{

	if err := mu.mr.DeleteMenuItem(&menuItem, id); err != nil{
		return err
	}
	return nil
}