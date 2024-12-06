package controller

import (
	"net/http"
	"strconv"

	"github.com/NazishAhsan/easy_busy_book_go/common"
	"github.com/NazishAhsan/easy_busy_book_go/model"
	"github.com/NazishAhsan/easy_busy_book_go/usecase"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type (
	SupplierController interface {
		GetSupplierList(c echo.Context) error
		CreateSupplier(c echo.Context) error
		UpdateSupplier(c echo.Context) error
		DeleteSupplier(c echo.Context) error
		SearchSupplier(c echo.Context) error
	}

	supplierController struct {
		su usecase.SupplierUsecase
	}
)

func NewSupplierController(su usecase.SupplierUsecase) SupplierController {
	return &supplierController{su}
}

func (sc *supplierController) GetSupplierList(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]

	supplierList, err := sc.su.GetSupplierList(uint(organizationID.(float64)))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, supplierList)
}

func (sc *supplierController) CreateSupplier(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]

	supplier := model.Supplier{}
	if err := c.Bind(&supplier); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	supplier.OrganizationID = uint(organizationID.(float64))
	supplierRes, err := sc.su.CreateSupplier(supplier)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, supplierRes)
}

func (sc *supplierController) UpdateSupplier(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	id := c.Param("id")
	supplierID, _ := strconv.Atoi(id)

	supplier := model.Supplier{}
	if err := c.Bind(&supplier); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	supplier.OrganizationID = uint(organizationID.(float64))
	resSupplier, err := sc.su.UpdateSupplier(supplier, uint(supplierID), uint(organizationID.(float64)))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, resSupplier)
}

func (sc *supplierController) DeleteSupplier(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	id := c.Param("id")
	supplierID, _ := strconv.Atoi(id)
	common.Logger.LogInfo().Msg(id)

	supplier := model.Supplier{}
	if err := c.Bind(&supplier); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err := sc.su.DeleteSupplier(supplier, uint(supplierID), uint(organizationID.(float64)))
	
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

func (sc *supplierController) SearchSupplier(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]

	term := c.QueryParam("term")

	supplierRes, err := sc.su.SearchSupplier(uint(organizationID.(float64)), term)

	if err!=nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, supplierRes)
}