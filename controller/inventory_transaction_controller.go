package controller

import (
	"net/http"
	"strconv"

	"github.com/Sayeda-fatima/restaurant_kot_backend/usecase"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type (
	InventoryTransactionController interface {
		GetInventoryTransactionList(c echo.Context) error
		AddStock(c echo.Context) error
		AdjustStock(c echo.Context) error
		RecordWaste(c echo.Context) error
		GetCostOfGoodsSold(c echo.Context) error
		CreateCurrentInventoryValue(c echo.Context) error
		GetWasteDuringTimePeriod(c echo.Context) error
		GetDailyConsumption(c echo.Context) error
	}

	inventoryTransactionController struct{
		iu usecase.InventoryTransactionUsecase
	}
)

func NewInventoryTransactionController(iu usecase.InventoryTransactionUsecase) InventoryTransactionController{
	return &inventoryTransactionController{iu}
}

func (ic *inventoryTransactionController) GetInventoryTransactionList(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	inventoryTransactions, err := ic.iu.GetInventoryTransactionList(uint(organizationID.(float64)), uint(restaurantID.(float64)))

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, inventoryTransactions)
}

func (ic *inventoryTransactionController) AddStock(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	product := c.Param("product")
	productID , _ := strconv.Atoi(product)

	quantity := c.FormValue("quantity")
	productQuantity, _ := strconv.Atoi(quantity)

	unitCost := c.FormValue("unit_cost")
	cost, _ := strconv.Atoi(unitCost)

	resInventory, err := ic.iu.AddStock(uint(organizationID.(float64)), uint(restaurantID.(float64)), uint(productID), productQuantity, cost)

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, resInventory)
}

func (ic *inventoryTransactionController) AdjustStock(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	product := c.Param("product")
	productID, _ := strconv.Atoi(product)
	quantity := c.FormValue("quantity")
	adjustmentQuantity, _ := strconv.Atoi(quantity)
	reason := c.FormValue("reason")

	resInventory, err := ic.iu.AdjustStock(uint(organizationID.(float64)), uint(restaurantID.(float64)), uint(productID), adjustmentQuantity, reason)

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, resInventory)
}

func (ic *inventoryTransactionController) RecordWaste(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	product := c.Param("product")
	productID, _ := strconv.Atoi(product)
	quantity := c.FormValue("quantity")
	wasteQuantity, _ := strconv.Atoi(quantity)
	reason := c.FormValue("reason")

	resInventory, err := ic.iu.RecordWaste(uint(organizationID.(float64)), uint(restaurantID.(float64)), uint(productID), wasteQuantity, reason)

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, resInventory)
}

func (ic *inventoryTransactionController) GetCostOfGoodsSold(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	dateFrom := c.FormValue("date_from")
	dateTo := c.FormValue("date_to")

	inventory, err := ic.iu.GetCostOfGoodsSold(uint(organizationID.(float64)), uint(restaurantID.(float64)), dateFrom, dateTo)

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, inventory)
}

func (ic *inventoryTransactionController) CreateCurrentInventoryValue(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	inventoryValue, err := ic.iu.CreateCurrentInventoryValue(uint(organizationID.(float64)), uint(restaurantID.(float64)))

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, inventoryValue)
}

func (ic *inventoryTransactionController) GetWasteDuringTimePeriod(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	dateFrom := c.FormValue("date_from")
	dateTo := c.FormValue("date_to")

	inventoryWaste, err := ic.iu.GetWasteDuringTimePeriod(uint(organizationID.(float64)), uint(restaurantID.(float64)), dateFrom, dateTo)

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, inventoryWaste)
}

func (ic *inventoryTransactionController) GetDailyConsumption(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	inventoryConsumption, err := ic.iu.GetDailyConsumption(uint(organizationID.(float64)), uint(restaurantID.(float64)))

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, inventoryConsumption)
}