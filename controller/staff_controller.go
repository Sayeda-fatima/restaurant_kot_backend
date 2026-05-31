package controller

import (
	"net/http"
	"strconv"

	"github.com/Sayeda-fatima/restaurant_kot_backend/common"
	"github.com/Sayeda-fatima/restaurant_kot_backend/model"
	"github.com/Sayeda-fatima/restaurant_kot_backend/usecase"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type (
	StaffController interface {
		GetStaffListByOrganization(c echo.Context) error
		GetStaffListByRestaurant(c echo.Context) error
		CreateStaff(c echo.Context) error
		UpdateStaff(c echo.Context) error
		DeleteStaff(c echo.Context) error
	}

	staffController struct{
		su usecase.StaffUsecase
	}
)

func NewStaffController(su usecase.StaffUsecase) StaffController{
	return &staffController{su}
}

func (sc *staffController) GetStaffListByOrganization(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	
	staffRes, err := sc.su.GetStaffListByOrganization(uint(organizationID.(float64)))

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, staffRes)
}

func (sc *staffController) GetStaffListByRestaurant(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	staffRes, err := sc.su.GetStaffListByRestaurant(uint(organizationID.(float64)), uint(restaurantID.(float64)))

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, staffRes)
}

func (sc *staffController) CreateStaff(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	staff := model.Staff{}
	if err := c.Bind(&staff); err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	staff.OrganizationID = uint(organizationID.(float64))
	staff.RestaurantID = uint(restaurantID.(float64))

	common.Logger.Log().Msgf("organizationID: %d\n restaurantID: %d",staff.OrganizationID, staff.RestaurantID)
	
	staffRes, err := sc.su.CreateStaff(staff); 
	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, staffRes)
}

func (sc *staffController) UpdateStaff(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	id := c.Param("id")
	staffID, _ := strconv.Atoi(id)

	staff := model.Staff{}
	if err := c.Bind(&staff); err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	staff.ID = uint(staffID)
	staff.OrganizationID = uint(organizationID.(float64))
	staff.RestaurantID = uint(restaurantID.(float64))
	staffRes, err := sc.su.UpdateStaff(staff, uint(staffID), uint(organizationID.(float64)), uint(restaurantID.(float64)))

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, staffRes)
}

func (sc *staffController) DeleteStaff(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]
	id := c.Param("id")
	staffID, _ := strconv.Atoi(id)

	staff := model.Staff{}
	if err := c.Bind(&staff); err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := sc.su.DeleteStaff(staff, uint(staffID), uint(organizationID.(float64)), uint(restaurantID.(float64))); err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}