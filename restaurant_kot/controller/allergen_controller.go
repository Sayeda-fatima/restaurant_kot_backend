package controller

import (
	"net/http"
	"strconv"

	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/model"
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/usecase"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type (
	AllergenController interface {
		GetAllergenList(c echo.Context) error
		CreateAllergen(c echo.Context) error
		UpdateAllergen(c echo.Context) error
		DeleteAllergen(c echo.Context) error
	}

	allergenController struct{
		au usecase.AllergenUsecase
	}
)

func NewAllergenController(au usecase.AllergenUsecase) AllergenController{
	return &allergenController{au}
}

func (ac *allergenController) GetAllergenList(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	allergenRes, err := ac.au.GetAllergenList(uint(organizationID.(float64)), uint(restaurantID.(float64)))

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, allergenRes)
}

func (ac *allergenController) CreateAllergen(c echo.Context) error{

	// user := c.Get("user").(*jwt.Token)
	// claims := user.Claims.(jwt.MapClaims)
	// organizationID := claims["organization_id"]
	// restaurantID := claims["restaraunt_id"]

	allergen := model.Allergen{}
	if err := c.Bind(&allergen); err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	allergenRes, err := ac.au.CreateAllergen(allergen)

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, allergenRes)
}

func (ac *allergenController) UpdateAllergen(c echo.Context) error{

	id := c.Param("id")
	allergenID, _ := strconv.Atoi(id)

	allergen := model.Allergen{}
	if err := c.Bind(&allergen); err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	allergen.ID = uint(allergenID)
	allergenRes, err := ac.au.UpdateAllergen(allergen, uint(allergenID))

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, allergenRes)
}

func (ac *allergenController) DeleteAllergen(c echo.Context) error{

	id := c.Param("id")
	allergenID, _ := strconv.Atoi(id)
	
	allergen := model.Allergen{}
	if err := c.Bind(&allergen); err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := ac.au.DeleteAllergen(allergen, uint(allergenID)); err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}