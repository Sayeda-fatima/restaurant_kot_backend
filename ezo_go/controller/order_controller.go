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
	OrderController interface {
		GetOrderList(c echo.Context) error
		CreateOrder(c echo.Context) error
		UpdateOrder(c echo.Context) error
		DeleteOrder(c echo.Context) error
	}

	orderController struct {
		ou usecase.OrderUsecase
	}
)

func NewOrderController(ou usecase.OrderUsecase) OrderController {
	return &orderController{ou}
}

func (oc *orderController) GetOrderList(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]

	orderList := model.Order{}
	if err := c.Bind(&orderList); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	orderListRes, err := oc.ou.GetOrderList(uint(organizationID.(float64)))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, orderListRes)
}

func (oc *orderController) CreateOrder(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]

	order := model.Order{}
	if err := c.Bind(&order); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	order.OrganizationID = uint(organizationID.(float64))
	orderRes, err := oc.ou.CreateOrder(order)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, orderRes)
}

func (oc *orderController) UpdateOrder(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	id := c.Param("id")
	orderID, _ := strconv.Atoi(id)

	order := model.Order{}
	if err := c.Bind(&order); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	order.ID = uint(orderID)
	order.OrganizationID = uint(organizationID.(float64))
	orderRes, err := oc.ou.UpdateOrder(order, uint(orderID))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, orderRes)
}

func (oc *orderController) DeleteOrder(c echo.Context) error {

	id := c.Param("id")
	orderID, _ := strconv.Atoi(id)

	order := model.Order{}
	if err := c.Bind(&order); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err := oc.ou.DeleteOrder(order, uint(orderID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}
