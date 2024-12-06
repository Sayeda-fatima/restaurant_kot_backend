package controller

import (
	"net/http"
	"strconv"

	"github.com/NazishAhsan/easy_busy_book_go/model"
	"github.com/NazishAhsan/easy_busy_book_go/usecase"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type (
	OrderItemController interface {
		GetOrderItemList(c echo.Context) error
		CreateOrderItem(c echo.Context) error
		UpdateOrderItem(c echo.Context) error
		DeleteOrderItem(c echo.Context) error
		InvoiceCustomer(c echo.Context) error
	}

	orderItemController struct{
		ou usecase.OrderItemUsecase
	}
)

func NewOrderItemController (ou usecase.OrderItemUsecase) OrderItemController{
	return &orderItemController{ou}
}

func (oc *orderItemController) GetOrderItemList(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	id := c.Param("id")
	orderID, _ := strconv.Atoi(id)
	
	orderItems := model.OrderItem{}
	if err := c.Bind(&orderItems); err!=nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	orderItemRes, err := oc.ou.GetOrderItemList(uint(organizationID.(float64)), uint(orderID))
	if err!=nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, orderItemRes)
}

func (oc *orderItemController) CreateOrderItem(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	id := c.Param("id")
	orderID, _ := strconv.Atoi(id)

	orderItem := model.OrderItem{}
	if err := c.Bind(&orderItem); err!=nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	orderItem.OrganizationID = uint(organizationID.(float64))
	orderItem.OrderID = uint(orderID)
	orderRes, err := oc.ou.CreateOrderItem(orderItem)
	if err!=nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, orderRes)
}

func (oc *orderItemController) UpdateOrderItem(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	id := c.Param("id")
	orderItemID, _ := strconv.Atoi(id)

	orderItem := model.OrderItem{}
	if err := c.Bind(&orderItem); err!=nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	orderItem.OrganizationID = uint(organizationID.(float64))
	orderItem.ID = uint(orderItemID)
	orderItemRes, err := oc.ou.UpdateOrderItem(orderItem, uint(orderItemID), uint(organizationID.(float64)))

	if err!=nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, orderItemRes)
}

func (oc *orderItemController) DeleteOrderItem(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	id := c.Param("id")
	orderItemID, _ := strconv.Atoi(id)

	orderItem := model.OrderItem{}
	if err := c.Bind(&orderItem); err!=nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err := oc.ou.DeleteOrderItem(orderItem, uint(orderItemID), uint(organizationID.(float64)))

	if err!=nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

func (oc *orderItemController) InvoiceCustomer(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	id := c.Param("id")
	orderID, _ := strconv.Atoi(id)

	dateFrom := c.FormValue("date_from")
	dateTo := c.FormValue("date_to")

	orderRes, err := oc.ou.InvoiceCustomer(uint(organizationID.(float64)), uint(orderID), dateFrom, dateTo)
	if err!=nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, orderRes)
}