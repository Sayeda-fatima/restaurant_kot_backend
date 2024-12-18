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
	OrderController interface {
		GetOrderList(c echo.Context) error
		CreateOrder(c echo.Context) error
		UpdateOrder(c echo.Context) error
		DeleteOrder(c echo.Context) error
		Checkout(c echo.Context) error
		TotalSales(c echo.Context) error
		TotalSalesByOrderType(c echo.Context) error
	}

	orderController struct{
		ou usecase.OrderUsecase
	}
)

func NewOrderController(ou usecase.OrderUsecase) OrderController{
	return &orderController{ou}
}

func (oc *orderController) GetOrderList(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	orderRes, err := oc.ou.GetOrderList(uint(organizationID.(float64)), uint(restaurantID.(float64)))

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, orderRes)
}

func (oc *orderController) CreateOrder(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	order := model.Order{}
	if err := c.Bind(&order); err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	order.OrganizationID = uint(organizationID.(float64))
	order.RestaurantID = uint(restaurantID.(float64))
	orderRes, err := oc.ou.CreateOrder(order)

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, orderRes)
}

func (oc *orderController) UpdateOrder(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	id := c.Param("id")
	orderID, _ := strconv.Atoi(id)


	order := model.Order{}
	if err := c.Bind(&order); err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	order.ID = uint(orderID)
	order.OrganizationID = uint(organizationID.(float64))
	order.RestaurantID = uint(restaurantID.(float64))
	orderRes, err := oc.ou.UpdateOrder(order, uint(orderID), uint(organizationID.(float64)), uint(restaurantID.(float64)))

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, orderRes)
}

func (oc *orderController) DeleteOrder(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	id := c.Param("id")
	orderID, _ := strconv.Atoi(id)

	order := model.Order{}
	if err := c.Bind(&order); err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := oc.ou.DeleteOrder(order, uint(orderID), uint(organizationID.(float64)), uint(restaurantID.(float64))); err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

func (oc *orderController) Checkout(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	cart := c.Param("cart")
	cartID, _ := strconv.Atoi(cart)

	order := model.Order{}
	if err := c.Bind(&order); err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	orderRes, err := oc.ou.Checkout(order, uint(organizationID.(float64)), uint(restaurantID.(float64)), uint(cartID))

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, orderRes)
}

func (oc *orderController) TotalSales(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	dateFrom := c.FormValue("date_from")
	dateTo := c.FormValue("date_to")

	result, err := oc.ou.TotalSales(uint(organizationID.(float64)), uint(restaurantID.(float64)), dateFrom, dateTo)

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func (oc *orderController) TotalSalesByOrderType(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	dateFrom := c.FormValue("date_from")
	dateTo := c.FormValue("date_to")

	result, err := oc.ou.TotalSalesByOrderType(uint(organizationID.(float64)), uint(restaurantID.(float64)), dateFrom, dateTo)

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}