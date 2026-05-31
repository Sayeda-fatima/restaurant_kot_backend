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
	OrderItemController interface {
		GetOrderItemList(c echo.Context) error
		CreateOrderItem(c echo.Context) error
		UpdateOrderItem(c echo.Context) error
		DeleteOrderItem(c echo.Context) error
		MostOrderedItems(c echo.Context) error
		DailySaleByItem(c echo.Context) error
	}

	orderItemController struct{
		ou usecase.OrderItemUsecase
	}
)

func NewOrderItemController(ou usecase.OrderItemUsecase) OrderItemController{
	return &orderItemController{ou}
}

func (oc *orderItemController) GetOrderItemList(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	order := c.Param("order")
	orderID, _ := strconv.Atoi(order)

	orderItemRes, err := oc.ou.GetOrderItemList(uint(organizationID.(float64)), uint(restaurantID.(float64)), uint(orderID))

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, orderItemRes)
}

func (oc *orderItemController) CreateOrderItem(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	order := c.Param("order")
	orderID, _ := strconv.Atoi(order)

	orderItem := model.OrderItem{}
	if err := c.Bind(&orderItem); err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	orderItem.OrderID = uint(orderID)
	orderItem.OrganizationID = uint(organizationID.(float64))
	orderItem.RestaurantID = uint(restaurantID.(float64))
	orderItemRes, err := oc.ou.CreateOrderItem(orderItem)

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, orderItemRes)
}

func (oc *orderItemController) UpdateOrderItem(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	order := c.Param("order")
	orderID, _ := strconv.Atoi(order)
	id := c.Param("id")
	orderItemID, _ := strconv.Atoi(id)

	orderItem := model.OrderItem{}
	if err := c.Bind(&orderItem); err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	orderItem.ID = uint(orderItemID)
	orderItem.OrderID = uint(orderID)
	orderItem.OrganizationID = uint(organizationID.(float64))
	orderItem.RestaurantID = uint(restaurantID.(float64))

	orderItemRes, err := oc.ou.UpdateOrderItem(orderItem, uint(orderItemID), uint(orderID), uint(organizationID.(float64)), uint(restaurantID.(float64)))

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, orderItemRes)
}

func (oc *orderItemController) DeleteOrderItem(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	order := c.Param("order")
	orderID, _ := strconv.Atoi(order)
	id := c.Param("id")
	orderItemID, _ := strconv.Atoi(id)

	orderItem := model.OrderItem{}
	if err := c.Bind(&orderItem); err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := oc.ou.DeleteOrderItem(orderItem, uint(orderItemID), uint(orderID), uint(organizationID.(float64)), uint(restaurantID.(float64))); err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func (oc *orderItemController) MostOrderedItems(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	dateFrom := c.FormValue("date_from")
	dateTo := c.FormValue("date_to")

	result, err := oc.ou.MostOrderedItems(uint(organizationID.(float64)), uint(restaurantID.(float64)), dateFrom, dateTo)

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}


func (oc *orderItemController) DailySaleByItem(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	dateFrom := c.FormValue("date_from")
	dateTo := c.FormValue("date_to")
	pageNo := c.QueryParam("page")
	page, _ := strconv.Atoi(pageNo)

	result, err := oc.ou.DailySaleByItem(uint(organizationID.(float64)), uint(restaurantID.(float64)), dateFrom, dateTo, page)

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}