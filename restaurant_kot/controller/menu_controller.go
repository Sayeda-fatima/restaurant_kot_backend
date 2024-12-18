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
	MenuController interface {
		GetMenuList(c echo.Context) error
		CreateMenu(c echo.Context) error
		UpdateMenu(c echo.Context) error
		DeleteMenu(c echo.Context) error
		FoodCost(c echo.Context) error
	}

	menuController struct{
		mu usecase.MenuUsecase
	}
)

func NewMenuController(mu usecase.MenuUsecase) MenuController{
	return &menuController{mu}
}

func (mc *menuController) GetMenuList(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	menuRes, err := mc.mu.GetMenuList(uint(organizationID.(float64)),uint(restaurantID.(float64)))

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, menuRes)
}

func (mc *menuController) CreateMenu(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	menu := model.Menu{}
	if err := c.Bind(&menu); err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	menu.OrganizationID = uint(organizationID.(float64))
	menu.RestaurantID = uint(restaurantID.(float64))
	menuRes, err := mc.mu.CreateMenu(menu)

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, menuRes)
}

func (mc *menuController) UpdateMenu(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	id := c.Param("id")
	menuID, _ := strconv.Atoi(id)

	menu := model.Menu{}
	if err := c.Bind(&menu); err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	menu.ID = uint(menuID)
	menu.OrganizationID = uint(organizationID.(float64))
	menu.RestaurantID = uint(restaurantID.(float64))

	menuRes, err := mc.mu.UpdateMenu(menu, uint(menuID), uint(organizationID.(float64)), uint(restaurantID.(float64)))
	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, menuRes)
}

func (mc *menuController) DeleteMenu(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]
	id := c.Param("id")
	menuID, _ := strconv.Atoi(id)

	menu := model.Menu{}
	if err := c.Bind(&menu); err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := mc.mu.DeleteMenu(menu, uint(menuID), uint(organizationID.(float64)), uint(restaurantID.(float64))); err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

func (mc *menuController) FoodCost(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	id := c.Param("id")
	menuID, _ := strconv.Atoi(id)


	foodCost, err := mc.mu.FoodCost(uint(menuID), uint(organizationID.(float64)), uint(restaurantID.(float64)))

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, foodCost)
}