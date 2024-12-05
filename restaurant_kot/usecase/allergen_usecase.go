package usecase

import (
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/model"
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/repository"
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/validator"
)

type (
	AllergenUsecase interface {
		GetAllergenList() ([]model.AllergenResponse, error)
		CreateAllergen(allergen model.Allergen) (model.AllergenResponse, error)
		UpdateAllergen(allergen model.Allergen, id uint) (model.AllergenResponse, error)
		DeleteAllergen(allergen model.Allergen, id uint) error
	}

	allergenUsecase struct{
		ar repository.AllergenRepository
		av validator.AllergenValidator
	}
)

func NewAllergenUsecase(ar repository.AllergenRepository, av validator.AllergenValidator) AllergenUsecase{
	return &allergenUsecase{ar,av}
}

func (au *allergenUsecase) GetAllergenList() ([]model.AllergenResponse, error){

	allergens := []model.Allergen{}

	if err := au.ar.GetAllergenList(&allergens); err != nil{
		return nil, err
	}

	resAllergen := []model.AllergenResponse{}
	for _, v := range(allergens){
		res := model.AllergenResponse{
			ID: v.ID,
			AllergenName: v.AllergenName,
		}
		resAllergen = append(resAllergen, res)
	}
	return resAllergen, nil
}

func (au *allergenUsecase) CreateAllergen(allergen model.Allergen) (model.AllergenResponse, error){

	if err := au.av.AllergenValidate(&allergen); err != nil{
		return model.AllergenResponse{}, err
	}

	if err := au.ar.CreateAllergen(&allergen); err != nil{
		return model.AllergenResponse{}, err
	}

	resAllergen := model.AllergenResponse{
		ID: allergen.ID,
		AllergenName: allergen.AllergenName,
	}
	return resAllergen, nil
}

func (au *allergenUsecase) UpdateAllergen(allergen model.Allergen, id uint) (model.AllergenResponse, error){

	if err := au.av.AllergenValidate(&allergen); err != nil{
		return model.AllergenResponse{}, err
	}

	if err := au.ar.UpdateAllergen(&allergen, id); err != nil{
		return model.AllergenResponse{}, err
	}

	resAllergen := model.AllergenResponse{
		ID: allergen.ID,
		AllergenName: allergen.AllergenName,
	}
	return resAllergen, nil
}

func (au *allergenUsecase) DeleteAllergen(allergen model.Allergen, id uint) error{

	if err := au.ar.DeleteAllergen(&allergen, id); err != nil{
		return err
	}
	return nil
}