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
	ProductImageController interface {
		GetProductImageList(c echo.Context) error
		AddProductImage(c echo.Context) error
		DeleteProductImage(c echo.Context) error
	}

	productImageController struct{
		pu usecase.ProductImageUsecase
	}
)

func NewProductImageController (pu usecase.ProductImageUsecase) ProductImageController{
	return &productImageController{pu}
}

func (pc *productImageController) GetProductImageList(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]

	id := c.Param("id")
	productID, _ := strconv.Atoi(id)

	productImageRes, err := pc.pu.GetProductImageList(uint(organizationID.(float64)), uint(productID))

	if err!=nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, productImageRes)
}

func (pc *productImageController) AddProductImage(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	id := c.Param("id")
	productID, _ := strconv.Atoi(id)

	file , err := c.FormFile("image")
	if err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	//productImage := model.ProductImage{}
	//productImage.OrganizationID = uint(organizationID.(float64))
	//productImage.ProductID = lkk,,uint(productID)
	productImageRes, err := pc.pu.AddProductImage(file, uint(organizationID.(float64)), uint(productID))

	if err!=nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, productImageRes)
}

func (pc *productImageController) DeleteProductImage(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	id := c.Param("id")
	productImageID,_ := strconv.Atoi(id)
	productImage := model.ProductImage{}
	if err := c.Bind(&productImage); err!=nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := pc.pu.DeleteProductImage(productImage, uint(productImageID), uint(organizationID.(float64))); err!=nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}