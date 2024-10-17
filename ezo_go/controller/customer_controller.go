package controller

import (
	"net/http"
	"strconv"

	"github.com/NazishAhsan/easy_busy_book_go/model"
	"github.com/NazishAhsan/easy_busy_book_go/usecase"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type CustomerController interface {
	CreateCustomer (c echo.Context) error
	GetCustomerList (c echo.Context) error
	UpdateCustomer (c echo.Context) error
	DeleteCustomer ( c echo.Context) error
}

type customerController struct {
	cu usecase.CustomerUsecase
}

func NewCustomerController (cu usecase.CustomerUsecase) CustomerController {
	return &customerController{cu}
}

func (cc *customerController) CreateCustomer (c echo.Context) error{

	// get organization_id from jwt token
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]

	customer := model.Customer{}
	if err := c.Bind(&customer); err!=nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	customer.OrganizationID = uint(organizationID.(float64))
	customerRes, err := cc.cu.CreateCustomer(customer)

	if err!=nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, customerRes)
}

func (cc *customerController) GetCustomerList (c echo.Context) error{

	user :=c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]

	customerList, err := cc.cu.GetCustomerList(uint(organizationID.(float64)))

	if err!=nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, customerList)
}

func (cc *customerController) UpdateCustomer (c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	id := c.Param("id")
	customerID, _ := strconv.Atoi(id)

	customer := model.Customer{}
	resCustomer := model.CustomerResponse{}

	if err := c.Bind(&customer); err!=nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	customer.OrganizationID = uint(organizationID.(float64))
	customerRes, err := cc.cu.UpdateCustomer(customer, uint(customerID), resCustomer) 

	if err!=nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, customerRes)
}

func (cc *customerController) DeleteCustomer (c echo.Context) error{

	// user := c.Get("user").(*jwt.Token)
	// claims := user.Claims.(jwt.MapClaims)
	// organizationID := claims["organization_id"]
	id := c.Param("id")
	customerID, _ := strconv.Atoi(id) 

	customer := model.Customer{}

	if err := c.Bind(&customer); err!=nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err := cc.cu.DeleteCustomer(customer, uint(customerID))

	if err!=nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}