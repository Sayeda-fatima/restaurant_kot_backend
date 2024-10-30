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
	CartItemController interface {
		GetCartItemList(c echo.Context) error
		CreateCartItem(c echo.Context) error
		UpdateCartItem(c echo.Context) error
		DeleteCartItem(c echo.Context) error
	}

	cartItemController struct {
		cu usecase.CartItemUsecase
	}
)

func NewCartItemController(cu usecase.CartItemUsecase) CartItemController {
	return &cartItemController{cu}
}

func (cc *cartItemController) GetCartItemList(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	id := c.Param("id")
	cartID, _ := strconv.Atoi(id)

	cartItemList, err := cc.cu.GetCartItemList(uint(organizationID.(float64)), uint(cartID))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, cartItemList)
}

func (cc *cartItemController) CreateCartItem(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	id := c.Param("id")
	cartID, _ := strconv.Atoi(id)

	cartItem := model.CartItem{}
	if err := c.Bind(&cartItem); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	cartItem.OrganizationID = uint(organizationID.(float64))
	cartItem.CartID = uint(cartID)
	cartItemRes, err := cc.cu.CreateCartItem(cartItem)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, cartItemRes)
}

func (cc *cartItemController) UpdateCartItem(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]

	id := c.Param("id")
	cartItemID, _ := strconv.Atoi(id)
	cart :=c.Param("cart")
	cartID, _ := strconv.Atoi(cart)

	cartItem := model.CartItem{}
	if err := c.Bind(&cartItem); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	cartItem.ID = uint(cartItemID)
	cartItem.CartID = uint(cartID)
	cartItem.OrganizationID = uint(organizationID.(float64))
	cartItemRes, err := cc.cu.UpdateCartItem(cartItem, uint(cartItemID))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, cartItemRes)
}

func (cc *cartItemController) DeleteCartItem(c echo.Context) error {

	id := c.Param("id")
	cartItemID, _ := strconv.Atoi(id)

	cartItem := model.CartItem{}
	if err := c.Bind(&cartItem); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err := cc.cu.DeleteCartItem(cartItem, uint(cartItemID))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}
