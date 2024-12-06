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
	ProductCategoryController interface {
		GetProductCategoryList(c echo.Context) error
		CreateProductCategory(c echo.Context) error
		UpdateProductCategory(c echo.Context) error
		DeleteProductCategory(c echo.Context) error
		SearchProductCategory(c echo.Context) error
	}

	productCategoryController struct {
		pu usecase.ProductCategoryUsecase
	}
)

func NewProductCategoryController(pu usecase.ProductCategoryUsecase) ProductCategoryController {
	return &productCategoryController{pu}
}

func (pc *productCategoryController) GetProductCategoryList(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]

	productList, err := pc.pu.GetProductCategoryList(uint(organizationID.(float64)))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, productList)
}

func (pc *productCategoryController) CreateProductCategory(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]

	productCategory := model.ProductCategory{}
	if err := c.Bind(&productCategory); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	productCategory.OrganizationID = uint(organizationID.(float64))
	productCategoryRes, err := pc.pu.CreateProductCategory(productCategory)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, productCategoryRes)
}

func (pc *productCategoryController) UpdateProductCategory(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	id := c.Param("id")
	productCategoryID, _ := strconv.Atoi(id)

	productCategory := model.ProductCategory{}
	if err := c.Bind(&productCategory); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	productCategory.OrganizationID = uint(organizationID.(float64))
	productCategory.ID = uint(productCategoryID)

	productCategoryRes, err := pc.pu.UpdateProductCategory(productCategory, uint(productCategoryID), uint(organizationID.(float64)))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, productCategoryRes)
}

func (pc *productCategoryController) DeleteProductCategory(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	id := c.Param("id")
	productCategoryID, _ := strconv.Atoi(id)

	productCategory := model.ProductCategory{}
	if err := c.Bind(&productCategory); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := pc.pu.DeleteProductCategory(productCategory, uint(productCategoryID), uint(organizationID.(float64))); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func (pc *productCategoryController) SearchProductCategory(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]

	term := c.QueryParam("term")

	productCategoryRes, err := pc.pu.SearchProductCategory(uint(organizationID.(float64)), term)

	if err!=nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, productCategoryRes)
}
