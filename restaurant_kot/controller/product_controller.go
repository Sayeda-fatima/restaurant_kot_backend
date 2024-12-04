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
	ProductController interface {
		GetAllProduct(c echo.Context) error
		CreateProduct(c echo.Context) error
		UpdateProduct(c echo.Context) error
		DeleteProduct(c echo.Context) error
	}

	productController struct{
		pu usecase.ProductUsecase
	}
)

func NewProductController(pu usecase.ProductUsecase) ProductController{
	return &productController{pu}
}

func (pc *productController) GetAllProduct(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	productRes, err := pc.pu.GetAllProduct(uint(organizationID.(float64)), uint(restaurantID.(float64)))

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, productRes)
}

func (pc *productController) CreateProduct(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	product := model.Product{}
	if err := c.Bind(&product); err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	product.OrganizationID = uint(organizationID.(float64))
	product.RestaurantID = uint(restaurantID.(float64))
	productRes, err := pc.pu.CreateProduct(product)

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, productRes)
}

func (pc *productController) UpdateProduct(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	id := c.Param("id")
	productID, _ := strconv.Atoi(id)

	product := model.Product{}
	if err := c.Bind(&product); err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	product.ID = uint(productID)
	product.OrganizationID = uint(organizationID.(float64))
	product.RestaurantID = uint(restaurantID.(float64))

	productRes, err := pc.pu.UpdateProduct(product, uint(productID))

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, productRes)
}

func (pc *productController) DeleteProduct(c echo.Context) error{

	id := c.Param("id")
	productID, _ := strconv.Atoi(id)

	product := model.Product{}
	if err := c.Bind(&product); err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := pc.pu.DeleteProduct(product, uint(productID)); err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}