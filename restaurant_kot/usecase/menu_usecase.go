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
		UpdateMenu(menu model.Menu, id uint, organizationID uint, restaurantID uint) (model.MenuResponse, error)
		DeleteMenu(menu model.Menu, id uint, organizationID uint, restaurantID uint) error
		FoodCost(id uint, organizationID uint, restaurantID uint) ([]map[string]interface{}, error)
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

func (mu *menuUsecase) UpdateMenu(menu model.Menu, id uint, organizationID uint, restaurantID uint) (model.MenuResponse, error){

	if err := mu.mv.MenuValidate(&menu); err != nil{
		return model.MenuResponse{}, err
	}

	if err := mu.mr.UpdateMenu(&menu, id, organizationID, restaurantID); err != nil{
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

func (mu *menuUsecase) DeleteMenu(menu model.Menu, id uint, organizationID uint, restaurantID uint) error{

	if err := mu.mr.DeleteMenu(&menu, id, organizationID, restaurantID); err != nil{
		return err
	}
	return nil
}

func (mu *menuUsecase) FoodCost(id uint, organizationID uint, restaurantID uint) ([]map[string]interface{}, error){

	menu := model.Menu{}
	if err := mu.mr.GetMenu(&menu, id, organizationID, restaurantID); err != nil{
		return nil, err
	}

	var result []map[string]interface{}
	for i, menuItem := range(menu.MenuItems){
		recipeCost := 0.00
		// calculate recipe price
		for _, v := range(menu.MenuItems[i].Recipe.RecipeProducts){
			recipeCost += float64(v.Product.UnitCost * v.Quantity)
		}
		// calculate per plate cost
		perPlateCost := float64(recipeCost/float64(menuItem.Recipe.Serving))*float64(menuItem.Serving)
		// calculate food cost 
		foodCost := (perPlateCost/float64(menuItem.Price))*100
		// calculate contribution margin - revenue generated from each dish
		contributionMargin := float64(menuItem.Price) - perPlateCost
		res := map[string]interface{}{
			"menu_item_id": menuItem.ID,
			"recipe_cost": recipeCost,
			"per_plate_cost": perPlateCost,
			"food_cost_percentage": foodCost,
			"contribution_margin": contributionMargin/100,
		}
		result = append(result, res)
	}

	return result, nil
}