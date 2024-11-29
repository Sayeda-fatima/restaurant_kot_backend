package controller

import (
	"net/http"
	"strconv"

	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/model"
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/usecase"
	"github.com/labstack/echo/v4"
)

type (
	RestaurantController interface {
		GetRestaurantList(c echo.Context) error
		CreateRestaurant(c echo.Context) error
		UpdateRestaurant(c echo.Context) error
		DeleteRestaurant(c echo.Context) error
	}

	restaurantController struct{
		ru usecase.RestaurantUsecase
	}
)

func NewRestaurantController (ru usecase.RestaurantUsecase) RestaurantController{
	return &restaurantController{ru}
}

func (rc *restaurantController) GetRestaurantList(c echo.Context) error{

	id := c.Param("id")
	organizationID, _ := strconv.Atoi(id)
	pageNo := c.QueryParam("page")
	page, _ := strconv.Atoi(pageNo)

	restaurantRes, err := rc.ru.GetRestaurantList(uint(organizationID), page)

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, restaurantRes)
}

func (rc *restaurantController) CreateRestaurant(c echo.Context) error{

	restaurant := model.Restaurant{}
	if err := c.Bind(&restaurant); err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	restaurantRes, err := rc.ru.CreateRestaurant(restaurant)

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, restaurantRes)
}

func (rc *restaurantController) UpdateRestaurant(c echo.Context) error{

	id := c.Param("id")
	restaurantID, _ := strconv.Atoi(id)
	restaurant := model.Restaurant{}

	if err := c.Bind(&restaurant); err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	restaurant.ID = uint(restaurantID)
	restaurantRes, err := rc.ru.UpdateRestaurant(restaurant, uint(restaurantID))

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, restaurantRes)
}

func (rc * restaurantController) DeleteRestaurant(c echo.Context) error{

	id := c.Param("id")
	restaurantID, _ := strconv.Atoi(id)

	restaurant := model.Restaurant{}
	if err := c.Bind(&restaurant); err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := rc.ru.DeleteRestaurant(restaurant, uint(restaurantID)); err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}