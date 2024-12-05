package usecase

import (
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/model"
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/repository"
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/validator"
)

type (
	MenuAllergenUsecase interface {
		GetMenuAllergenList(menuItemID uint) ([]model.MenuAllergenResponse, error)
		CreateMenuAllergen(menuAllergen model.MenuAllergen) (model.MenuAllergenResponse, error)
		UpdateMenuAllergen(menuAllergen model.MenuAllergen, id uint) (model.MenuAllergenResponse, error)
		DeleteMenuAllergen(menuAllergen model.MenuAllergen, id uint) error
	}

	menuAllergenUsecase struct{
		mr repository.MenuAllergenRepository
		mv validator.MenuAllergenValidator
	}
)

func NewMenuAllergenUsecase(mr repository.MenuAllergenRepository, mv validator.MenuAllergenValidator) MenuAllergenUsecase{
	return &menuAllergenUsecase{mr, mv}
}

func (mu *menuAllergenUsecase) GetMenuAllergenList(menuItemID uint) ([]model.MenuAllergenResponse, error){

	menuAllergen := []model.MenuAllergen{}
	if err := mu.mr.GetMenuAllergenList(&menuAllergen, menuItemID); err != nil{
		return nil, err
	}

	resMenuAllergen := []model.MenuAllergenResponse{}
	for _, v := range(menuAllergen){
		res := model.MenuAllergenResponse{
			ID: v.ID,
			MenuItemID: v.MenuItemID,
			AllergenID: v.AllergenID,
			Allergen: v.Allergen,
		}
		resMenuAllergen = append(resMenuAllergen, res)
	}
	return resMenuAllergen, nil
}

func (mu *menuAllergenUsecase) CreateMenuAllergen(menuAllergen model.MenuAllergen) (model.MenuAllergenResponse, error){

	if err := mu.mv.MenuAllergenValidate(&menuAllergen); err != nil{
		return model.MenuAllergenResponse{}, err
	}

	if err := mu.mr.CreateMenuAllergen(&menuAllergen); err != nil{
		return model.MenuAllergenResponse{}, err
	}
	
	resMenuAllergen := model.MenuAllergenResponse{
		ID: menuAllergen.ID,
		MenuItemID: menuAllergen.MenuItemID,
		AllergenID: menuAllergen.AllergenID,
	}
	return resMenuAllergen, nil
}

func (mu *menuAllergenUsecase) UpdateMenuAllergen(menuAllergen model.MenuAllergen, id uint) (model.MenuAllergenResponse, error){

	if err := mu.mv.MenuAllergenValidate(&menuAllergen); err != nil{
		return model.MenuAllergenResponse{}, err
	}

	if err := mu.mr.UpdateMenuAllergen(&menuAllergen, id); err != nil{
		return model.MenuAllergenResponse{}, err
	}
	
	resMenuAllergen := model.MenuAllergenResponse{
		ID: menuAllergen.ID,
		MenuItemID: menuAllergen.MenuItemID,
		AllergenID: menuAllergen.AllergenID,
	}
	return resMenuAllergen, nil
}

func (mu *menuAllergenUsecase) DeleteMenuAllergen(menuAllergen model.MenuAllergen, id uint) error{

	if err := mu.mr.DeleteMenuAllergen(&menuAllergen, id); err != nil{
		return err
	}
	return nil
}