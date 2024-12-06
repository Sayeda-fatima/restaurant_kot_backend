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
	CustomerController interface {
		GetCustomerList(c echo.Context) error
		CreateCustomer(c echo.Context) error
		UpdateCustomer(c echo.Context) error
		DeleteCustomer(c echo.Context) error
	}

	customerController struct{
		cu usecase.CustomerUsecase
	}
)

func NewCustomerController(cu usecase.CustomerUsecase) CustomerController{
	return &customerController{cu}
}

func (cc *customerController) GetCustomerList(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	customerRes, err := cc.cu.GetCustomerList(uint(organizationID.(float64)), uint(restaurantID.(float64)))
	
	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, customerRes)
}

func (cc *customerController) CreateCustomer(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	customer := model.Customer{}
	if err := c.Bind(&customer); err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	customer.OrganizationID = uint(organizationID.(float64))
	customer.RestaurantID = uint(restaurantID.(float64))
	
	customerRes, err := cc.cu.CreateCustomer(customer)

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, customerRes)
}

func (cc *customerController) UpdateCustomer(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	id := c.Param("id")
	customerID, _ := strconv.Atoi(id)

	customer := model.Customer{}
	if err := c.Bind(&customer); err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	
	customer.ID	= uint(customerID)
	customer.OrganizationID = uint(organizationID.(float64))
	customer.RestaurantID = uint(restaurantID.(float64))

	customerRes, err := cc.cu.UpdateCustomer(customer, uint(customerID), uint(organizationID.(float64)), uint(restaurantID.(float64)))

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, customerRes)
}

func (cc *customerController) DeleteCustomer(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	id := c.Param("id")
	customerID, _ := strconv.Atoi(id)

	customer := model.Customer{}
	if err := c.Bind(&customer); err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := cc.cu.DeleteCustomer(customer, uint(customerID), uint(organizationID.(float64)), uint(restaurantID.(float64))); err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}