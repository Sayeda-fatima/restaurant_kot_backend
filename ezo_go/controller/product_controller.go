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
	ProductController interface {
		GetProductList(c echo.Context) error
		CreateProduct(c echo.Context) error
		UpdateProduct(c echo.Context) error
		DeleteProduct(c echo.Context) error
		SearchProduct(c echo.Context) error
		GetProduct(c echo.Context) error
		UpdateStockOfProduct(c echo.Context) error
	}

	productController struct {
		pu usecase.ProductUsecase
	}
)

func NewProductController(pu usecase.ProductUsecase) ProductController {
	return &productController{pu}
}

func (pc *productController) GetProductList(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]

	productList, err := pc.pu.GetProductList(uint(organizationID.(float64)))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, productList)
}

func (pc *productController) CreateProduct(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]

	product := model.Product{}

	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	product.OrganizationID = uint(organizationID.(float64))
	productRes, err := pc.pu.CreateProduct(product)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, productRes)
}

func (pc *productController) UpdateProduct(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	id := c.Param("id")
	productID, _ := strconv.Atoi(id)
	common.Logger.LogInfo().Msg(id)

	product := model.Product{}

	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	product.OrganizationID = uint(organizationID.(float64))
	product.ID = uint(productID)
	productRes, err := pc.pu.UpdateProduct(product, uint(productID))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, productRes)
}

func (pc *productController) DeleteProduct(c echo.Context) error {

	id := c.Param("id")
	productID, _ := strconv.Atoi(id)

	product := model.Product{}

	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err := pc.pu.DeleteProduct(product, uint(productID))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func (pc *productController) SearchProduct(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]

	term := c.QueryParam("term")

	productRes, err := pc.pu.SearchProduct(uint(organizationID.(float64)), term)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, productRes)
}

func (pc *productController) GetProduct(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	id := c.Param("id")
	productID, _ := strconv.Atoi(id)

	product := model.Product{}
	if err := c.Bind(&product); err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	productRes, err := pc.pu.GetProduct(uint(organizationID.(float64)), uint(productID))
	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, productRes)
}

func (pc *productController) UpdateStockOfProduct(c echo.Context) error {

	id := c.Param("id")
	productID, _ := strconv.Atoi(id)
	quantity := c.FormValue("quantity")
	productQuantity, _ := strconv.Atoi(quantity)

	product := model.Product{}
	productRes, err := pc.pu.UpdateStockOfProduct(product, uint(productID), productQuantity)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, productRes)
}
