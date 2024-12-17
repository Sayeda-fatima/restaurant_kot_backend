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
	MenuItemController interface {
		GetMenuItemList(c echo.Context) error
		CreateMenuItem(c echo.Context) error
		UpdateMenuItem(c echo.Context) error
		DeleteMenuItem(c echo.Context) error
		UpdateMenuItemIsActivated(c echo.Context) error
	}

	menuItemController struct{
		mu usecase.MenuItemUsecase
	}
)

func NewMenuItemController(mu usecase.MenuItemUsecase) MenuItemController{
	return &menuItemController{mu}
}

func (mc *menuItemController) GetMenuItemList(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	menu := c.Param("menuID")
	menuID, _ := strconv.Atoi(menu)

	menuItemRes, err := mc.mu.GetMenuItemList(uint(organizationID.(float64)), uint(restaurantID.(float64)), uint(menuID))

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, menuItemRes)
}

func (mc *menuItemController) CreateMenuItem(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	menu := c.Param("menuID")
	menuID, _ := strconv.Atoi(menu) 

	menuItem := model.MenuItem{}
	if err := c.Bind(&menuItem); err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	menuItem.OrganizationID = uint(organizationID.(float64))
	menuItem.RestaurantID = uint(restaurantID.(float64))
	menuItem.MenuID = uint(menuID)
	menuItemRes, err := mc.mu.CreateMenuItem(menuItem)

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, menuItemRes)
}

func (mc *menuItemController) UpdateMenuItem(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	menu := c.Param("menuID")
	menuID, _ := strconv.Atoi(menu)
	id := c.Param("id")
	menuItemID, _ := strconv.Atoi(id)

	menuItem := model.MenuItem{}
	if err := c.Bind(&menuItem); err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	menuItem.ID = uint(menuItemID)
	menuItem.MenuID = uint(menuID)
	menuItem.OrganizationID = uint(organizationID.(float64))
	menuItem.RestaurantID = uint(restaurantID.(float64))

	menuItemRes, err := mc.mu.UpdateMenuItem(menuItem, uint(menuItemID), uint(organizationID.(float64)), uint(restaurantID.(float64)))

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, menuItemRes)
}

func (mc *menuItemController) DeleteMenuItem(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]
	id := c.Param("id")
	menuItemID, _ := strconv.Atoi(id)

	menuItem := model.MenuItem{}
	if err := c.Bind(&menuItem); err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := mc.mu.DeleteMenuItem(menuItem, uint(menuItemID), uint(organizationID.(float64)), uint(restaurantID.(float64))); err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func (mc *menuItemController) UpdateMenuItemIsActivated(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	id := c.Param("id")
	menuItemID, _ := strconv.Atoi(id)
	status := c.FormValue("status")
	menuStatus, _ := strconv.ParseBool(status)

	menuItem := model.MenuItem{}
	if err := c.Bind(&menuItem); err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := mc.mu.UpdateMenuItemIsActivated(menuItem, uint(menuItemID), uint(organizationID.(float64)), uint(restaurantID.(float64)), menuStatus); err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "success"})
}