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
	CartController interface {
		GetCartList(c echo.Context) error
		CreateCart(c echo.Context) error
		UpdateCart(c echo.Context) error
		UpdateCartStatus(c echo.Context) error
		DeleteCart(c echo.Context) error
	}

	cartController struct{
		cu usecase.CartUsecase
	}
)

func NewCartController(cu usecase.CartUsecase) CartController{
	return &cartController{cu}
}

func (cc *cartController) GetCartList(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	cartRes, err := cc.cu.GetCartList(uint(organizationID.(float64)), uint(restaurantID.(float64)))

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, cartRes)
}

func (cc *cartController) CreateCart(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	cart := model.Cart{}
	if err := c.Bind(&cart); err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	cart.OrganizationID = uint(organizationID.(float64))
	cart.RestaurantID = uint(restaurantID.(float64))

	cartRes, err := cc.cu.CreateCart(cart)

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, cartRes)
}

func (cc *cartController) UpdateCart(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	id := c.Param("id")
	cartID, _ := strconv.Atoi(id)

	cart := model.Cart{}
	if err := c.Bind(&cart); err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	cart.ID = uint(cartID)
	cart.OrganizationID = uint(organizationID.(float64))
	cart.RestaurantID = uint(restaurantID.(float64))

	cartRes, err := cc.cu.UpdateCart(cart, uint(cartID), uint(organizationID.(float64)), uint(restaurantID.(float64)))

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, cartRes)
}

func (cc *cartController) UpdateCartStatus(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	id := c.Param("id")
	cartID, _ := strconv.Atoi(id)

	status := c.FormValue("status")

	cart := model.Cart{}
	if err := c.Bind(&cart); err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	cartRes, err := cc.cu.UpdateCartStatus(cart, uint(cartID), uint(organizationID.(float64)), uint(restaurantID.(float64)), status)

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, cartRes)
}

func (cc *cartController) DeleteCart(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	id := c.Param("id")
	cartID, _ := strconv.Atoi(id)

	if err := cc.cu.DeleteCart(uint(cartID), uint(organizationID.(float64)), uint(restaurantID.(float64))); err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}