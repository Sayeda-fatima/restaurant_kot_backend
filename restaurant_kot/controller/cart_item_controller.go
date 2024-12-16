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
	CartItemController interface {
		GetCartItemList(c echo.Context) error
		CreateCartItem(c echo.Context) error
		UpdateCartItem(c echo.Context) error
		UpdateCartItemStatus(c echo.Context) error
		DeleteCartItem(c echo.Context) error
		SendCartItemToKitchen(c echo.Context) error
	}

	cartItemController struct{
		cu usecase.CartItemUsecase
	}
)

func NewCartItemController(cu usecase.CartItemUsecase) CartItemController{
	return &cartItemController{cu}
}

func (cc *cartItemController) GetCartItemList(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	cart := c.Param("cart")
	cartID, _ := strconv.Atoi(cart)

	cartItemRes, err := cc.cu.GetCartItemList(uint(cartID), uint(restaurantID.(float64)), uint(organizationID.(float64)))

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, cartItemRes)
}

func(cc *cartItemController) CreateCartItem(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	cart := c.Param("cart")
	cartID, _ := strconv.Atoi(cart)

	cartItem := model.CartItem{}
	if err := c.Bind(&cartItem); err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	cartItem.CartID = uint(cartID)
	cartItem.OrganizationID = uint(organizationID.(float64))
	cartItem.RestaurantID = uint(restaurantID.(float64))

	cartItemRes, err := cc.cu.CreateCartItem(cartItem)

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, cartItemRes)
}

func (cc *cartItemController) UpdateCartItem(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	cart := c.Param("cart")
	cartID, _ := strconv.Atoi(cart)
	id := c.Param("id")
	cartItemID, _ := strconv.Atoi(id)

	cartItem := model.CartItem{}
	if err := c.Bind(&cartItem); err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	cartItem.ID = uint(cartItemID)
	cartItem.CartID = uint(cartID)
	cartItem.OrganizationID = uint(organizationID.(float64))
	cartItem.RestaurantID = uint(restaurantID.(float64))

	cartItemRes, err := cc.cu.UpdateCartItem(cartItem, uint(cartItemID), uint(cartID), uint(organizationID.(float64)), uint(restaurantID.(float64)))

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, cartItemRes)
}

func (cc *cartItemController) UpdateCartItemStatus(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	cart := c.Param("cart")
	cartID, _ := strconv.Atoi(cart)
	id := c.Param("id")
	cartItemID, _ := strconv.Atoi(id)

	status := c.FormValue("status")

	cartItem := model.CartItem{}
	if err := c.Bind(&cartItem); err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	cartItemRes, err := cc.cu.UpdateCartItemStatus(cartItem, uint(cartItemID), uint(cartID), uint(restaurantID.(float64)), uint(organizationID.(float64)), status)

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, cartItemRes)
}

func (cc *cartItemController) DeleteCartItem(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	cart := c.Param("cart")
	cartID, _ := strconv.Atoi(cart)
	id := c.Param("id")
	cartItemID, _ := strconv.Atoi(id)

	cartItem := model.CartItem{}
	if err := c.Bind(&cartItem); err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := cc.cu.DeleteCartItem(cartItem, uint(cartItemID), uint(cartID), uint(organizationID.(float64)), uint(restaurantID.(float64))); err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

func (cc *cartItemController) SendCartItemToKitchen(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	cart := c.Param("cart")
	cartID, _ := strconv.Atoi(cart)

	cartRes, err := cc.cu.SendCartItemToKitchen(uint(cartID), uint(organizationID.(float64)), uint(restaurantID.(float64)))

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, cartRes)
}