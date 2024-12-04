package usecase

import (
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/model"
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/repository"
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/validator"
)

type (
	MenuUsecase interface {
		GetMenuList(organizationID uint, restaurantID uint) ([]model.MenuResponse, error)
		CreateMenu(menu model.Menu) (model.MenuResponse, error)
		UpdateMenu(menu model.Menu, id uint) (model.MenuResponse, error)
		DeleteMenu(menu model.Menu, id uint) error
	}

	menuUsecase struct{
		mr repository.MenuRepository
		mv validator.MenuValidator
	}
)

func NewMenuUsecase(mr repository.MenuRepository, mv validator.MenuValidator) MenuUsecase{
	return &menuUsecase{mr,mv}
}

func (mu *menuUsecase) GetMenuList(organizationID uint, restaurantID uint) ([]model.MenuResponse, error){

	menu := []model.Menu{}

	if err := mu.mr.GetMenuList(&menu, organizationID, restaurantID); err != nil{
		return nil, err
	}
	
	resMenu := []model.MenuResponse{}
	for _, v := range(menu){
		res := model.MenuResponse{
			ID: v.ID,
			OrganizationID: v.OrganizationID,
			RestaurantID: v.RestaurantID,
			Name: v.Name,
			MenuItems: v.MenuItems,
		}
		resMenu = append(resMenu, res)
	}
	return resMenu, nil
}

func (mu *menuUsecase) CreateMenu(menu model.Menu) (model.MenuResponse, error){

	if err := mu.mv.MenuValidate(&menu); err != nil{
		return model.MenuResponse{}, err
	}

	// assigning organizationID and restaurantID in menu items
	for i := range menu.MenuItems{
		menu.MenuItems[i].OrganizationID = menu.OrganizationID
		menu.MenuItems[i].RestaurantID = menu.RestaurantID
	}

	if err := mu.mr.CreateMenu(&menu); err != nil{
		return model.MenuResponse{}, err
	}

	resMenu := model.MenuResponse{
		ID: menu.ID,
		OrganizationID: menu.OrganizationID,
		RestaurantID: menu.RestaurantID,
		Name: menu.Name,
		MenuItems: menu.MenuItems,
	}
	return resMenu, nil
}

func (mu *menuUsecase) UpdateMenu(menu model.Menu, id uint) (model.MenuResponse, error){

	if err := mu.mv.MenuValidate(&menu); err != nil{
		return model.MenuResponse{}, err
	}

	if err := mu.mr.UpdateMenu(&menu, id); err != nil{
		return model.MenuResponse{}, err
	}

	resMenu := model.MenuResponse{
		ID: menu.ID,
		OrganizationID: menu.OrganizationID,
		RestaurantID: menu.RestaurantID,
		Name: menu.Name,
		MenuItems: menu.MenuItems,
	}
	return resMenu, nil
}

func (mu *menuUsecase) DeleteMenu(menu model.Menu, id uint) error{

	if err := mu.mr.DeleteMenu(&menu, id); err != nil{
		return err
	}
	return nil
}