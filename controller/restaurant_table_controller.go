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
	RestaurantTableController interface {
		GetRestaurantTableList(c echo.Context) error
		CreateRestaurantTable(c echo.Context) error
		UpdateRestaurantTable(c echo.Context) error
		DeleteRestaurantTable(c echo.Context) error
	}

	restaurantTableController struct {
		ru usecase.RestaurantTableUsecase
	}
)

func NewRestaurantTableController(ru usecase.RestaurantTableUsecase) RestaurantTableController {
	return &restaurantTableController{ru}
}

func (rc *restaurantTableController) GetRestaurantTableList(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	restaurantTableRes, err := rc.ru.GetRestaurantTableList(uint(organizationID.(float64)), uint(restaurantID.(float64)))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, restaurantTableRes)
}

func (rc *restaurantTableController) CreateRestaurantTable(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	restaurantTable := model.RestaurantTable{}
	if err := c.Bind(&restaurantTable); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	restaurantTable.OrganizationID = uint(organizationID.(float64))
	restaurantTable.RestaurantID = uint(restaurantID.(float64))

	restaurantTableRes, err := rc.ru.CreateRestaurantTable(restaurantTable)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, restaurantTableRes)
}

func (rc *restaurantTableController) UpdateRestaurantTable(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	id := c.Param("id")
	restaurantTableID, _ := strconv.Atoi(id)

	restaurantTable := model.RestaurantTable{}
	if err := c.Bind(&restaurantTable); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	restaurantTable.ID = uint(restaurantTableID)
	restaurantTable.OrganizationID = uint(organizationID.(float64))
	restaurantTable.RestaurantID = uint(restaurantID.(float64))

	restaurantTableRes, err := rc.ru.UpdateRestaurantTable(restaurantTable, uint(restaurantTableID), uint(organizationID.(float64)), uint(restaurantID.(float64)))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, restaurantTableRes)
}

func (rc *restaurantTableController) DeleteRestaurantTable(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]
	id := c.Param("id")
	restaurantTableID, _ := strconv.Atoi(id)

	restaurantTable := model.RestaurantTable{}
	if err := c.Bind(&restaurantTable); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := rc.ru.DeleteRestaurantTable(restaurantTable, uint(restaurantTableID), uint(organizationID.(float64)), uint(restaurantID.(float64))); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}
