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
	ProductStockController interface {
		GetProductStockList(c echo.Context) error
		CreateProductStock(c echo.Context) error
		UpdateProductStock(c echo.Context) error
		DeleteProductStock(c echo.Context) error
		GetProductStockListByUpdateType(c echo.Context) error
	}
	productStockController struct{
		pu usecase.ProductStockUsecase
	}
)

func NewProductStockController (pu usecase.ProductStockUsecase) ProductStockController {
	return &productStockController{pu}
}

func (pc *productStockController) GetProductStockList(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]

	productStockRes, err := pc.pu.GetProductStockList(uint(organizationID.(float64)))

	if err!=nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, productStockRes)
}

func (pc *productStockController) CreateProductStock(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	id := c.Param("id")
	productID, _ := strconv.Atoi(id)
	quantity := c.FormValue("quantity")
	productQuantity, _ := strconv.Atoi(quantity)
	
	productStockRes, err := pc.pu.CreateProductStock(uint(organizationID.(float64)), uint(productID), productQuantity)

	if err!=nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, productStockRes)
}

func (pc *productStockController) UpdateProductStock(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	id := c.Param("id")
	productStockID, _ := strconv.Atoi(id)

	productStock := model.ProductStock{}
	if err := c.Bind(&productStock); err!=nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	productStock.ID = uint(productStockID)
	productStock.OrganizationID = uint(organizationID.(float64))
	productStockRes, err := pc.pu.UpdateProductStock(productStock, uint(productStockID))

	if err!=nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, productStockRes)
}

func (pc *productStockController) DeleteProductStock(c echo.Context) error{

	id := c.Param("id")
	productStockID, _ := strconv.Atoi(id)

	productStock := model.ProductStock{}
	if err := c.Bind(&productStock); err!=nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := pc.pu.DeleteProductStock(productStock, uint(productStockID)); err!=nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

func (pc *productStockController) GetProductStockListByUpdateType(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]

	term := c.QueryParam("term")

	productStockRes, err := pc.pu.GetProductStockListByUpdateType(uint(organizationID.(float64)), term)

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, productStockRes)
}