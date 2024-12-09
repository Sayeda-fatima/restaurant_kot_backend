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
	RecipeController interface {
		GetRecipeList(c echo.Context) error
		GetRecipe(c echo.Context) error
		CreateRecipe(c echo.Context) error
		UpdateRecipe(c echo.Context) error
		DeleteRecipe(c echo.Context) error
		GetRecipeCost(c echo.Context) error
	}

	recipeController struct{
		ru usecase.RecipeUsecase
	}
)

func NewRecipeController(ru usecase.RecipeUsecase) RecipeController{
	return &recipeController{ru}
}

func (rc *recipeController) GetRecipeList(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	recipeRes, err := rc.ru.GetRecipeList(uint(organizationID.(float64)), uint(restaurantID.(float64)))

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, recipeRes)
}

func (rc *recipeController) GetRecipe(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	id := c.Param("id")
	recipeID, _ := strconv.Atoi(id)
	
	recipeRes, err := rc.ru.GetRecipe(uint(organizationID.(float64)), uint(restaurantID.(float64)), uint(recipeID))

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, recipeRes)
}

func (rc *recipeController) CreateRecipe(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	recipe := model.Recipe{}
	if err := c.Bind(&recipe); err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	recipe.OrganizationID = uint(organizationID.(float64))
	recipe.RestaurantID = uint(restaurantID.(float64))
	recipeRes, err := rc.ru.CreateRecipe(recipe)

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, recipeRes)
}

func (rc *recipeController) UpdateRecipe(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	id := c.Param("id")
	recipeID, _ := strconv.Atoi(id)

	recipe := model.Recipe{}
	if err := c.Bind(&recipe); err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	recipe.ID = uint(recipeID)
	recipe.OrganizationID = uint(organizationID.(float64))
	recipe.RestaurantID = uint(restaurantID.(float64))
	recipeRes, err := rc.ru.UpdateRecipe(recipe, uint(recipeID), uint(organizationID.(float64)), uint(restaurantID.(float64)))

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, recipeRes)
}

func (rc *recipeController) DeleteRecipe(c echo.Context) error{
	
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]
	id := c.Param("id")
	recipeID, _ := strconv.Atoi(id)

	recipe := model.Recipe{}
	if err := c.Bind(&recipe); err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := rc.ru.DeleteRecipe(recipe, uint(recipeID), uint(organizationID.(float64)), uint(restaurantID.(float64))); err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

func (rc *recipeController) GetRecipeCost(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	id := c.Param("id")
	recipeID, _ := strconv.Atoi(id)

	recipeRes, err := rc.ru.GetRecipeCost(uint(recipeID), uint(organizationID.(float64)), uint(restaurantID.(float64)))

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, recipeRes)
}