package controller

import (
	"net/http"
	"strconv"

	"github.com/Sayeda-fatima/restaurant_kot_backend/model"
	"github.com/Sayeda-fatima/restaurant_kot_backend/usecase"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type (
	RecipeProductController interface {
		GetProductInRecipe(c echo.Context) error
		CreateRecipeProduct(c echo.Context) error
		UpdateRecipeProduct(c echo.Context) error
		DeleteRecipeProduct(c echo.Context) error
	}

	recipeProductController struct{
		ru usecase.RecipeProductUsecase
	}
)

func NewRecipeProductController(ru usecase.RecipeProductUsecase) RecipeProductController{
	return &recipeProductController{ru}
}

func (rc *recipeProductController) GetProductInRecipe(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	recipe := c.Param("recipe")
	recipeID, _ := strconv.Atoi(recipe)

	recipeProductRes, err := rc.ru.GetProductInRecipe(uint(recipeID), uint(organizationID.(float64)), uint(restaurantID.(float64)))

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, recipeProductRes)
}

func (rc *recipeProductController) CreateRecipeProduct(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	recipe := c.Param("recipe")
	recipeID, _ := strconv.Atoi(recipe)

	recipeProduct := model.RecipeProduct{}
	if err := c.Bind(&recipeProduct); err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	recipeProduct.RecipeID = uint(recipeID)
	recipeProduct.OrganizationID = uint(organizationID.(float64))
	recipeProduct.RestaurantID = uint(restaurantID.(float64))
	recipeProductRes, err := rc.ru.CreateRecipeProduct(recipeProduct)

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, recipeProductRes)
}

func (rc *recipeProductController) UpdateRecipeProduct(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	recipe := c.Param("recipe")
	recipeID, _ := strconv.Atoi(recipe)

	id := c.Param("id")
	recipeProductID, _ := strconv.Atoi(id)

	recipeProduct := model.RecipeProduct{}
	if err := c.Bind(&recipeProduct); err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	recipeProduct.RecipeID = uint(recipeID)
	recipeProduct.OrganizationID = uint(organizationID.(float64))
	recipeProduct.RestaurantID = uint(restaurantID.(float64))
	recipeProductRes, err := rc.ru.UpdateRecipeProduct(recipeProduct, uint(recipeProductID), uint(organizationID.(float64)), uint(restaurantID.(float64)))

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, recipeProductRes)
}

func (rc *recipeProductController) DeleteRecipeProduct(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	id := c.Param("id")
	recipeProductID, _ := strconv.Atoi(id)

	recipeProduct := model.RecipeProduct{}
	if err := c.Bind(&recipeProduct); err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := rc.ru.DeleteRecipeProduct(recipeProduct, uint(recipeProductID), uint(organizationID.(float64)), uint(restaurantID.(float64))); err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}